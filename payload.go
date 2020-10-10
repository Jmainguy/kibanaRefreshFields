package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func formatPayload(fieldList *FieldList, config Config) (body io.Reader) {
	payload := Payload{}
	payload.Attributes.Title = "filebeat-*"
	payload.Attributes.TimeFieldName = "@timestamp"

	payload.Attributes.Fields = "["
	if config.KibanaFilter != "" {
		fmt.Printf("Filtering out %s\n", config.KibanaFilter)
	}
	for _, v := range fieldList.Fields {
		// If you want to filter out some fields
		if config.KibanaFilter != "" {
			if !strings.Contains(v.Name, config.KibanaFilter) {
				fields, err := json.Marshal(v)
				check(err)
				stringFields := string(fields)
				payload.Attributes.Fields = fmt.Sprintf("%s%s,", payload.Attributes.Fields, stringFields)
			}
		} else {
			fields, err := json.Marshal(v)
			check(err)
			stringFields := string(fields)
			payload.Attributes.Fields = fmt.Sprintf("%s%s,", payload.Attributes.Fields, stringFields)
		}
	}
	payload.Attributes.Fields = strings.TrimRight(payload.Attributes.Fields, ",")
	payload.Attributes.Fields = fmt.Sprintf("%s]", payload.Attributes.Fields)

	jsonNewData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	body = bytes.NewReader(jsonNewData)

	return body
}

func curl(config Config, url, requestMethod string) (bodyBytes []byte) {
	cfg := &tls.Config{
		InsecureSkipVerify: false,
	}
	http.DefaultClient.Transport = &http.Transport{
		TLSClientConfig: cfg,
	}
	if config.Insecure {
		cfg.InsecureSkipVerify = true
	}
	blankCertificatePool := x509.NewCertPool()
	if config.CertificatePool != blankCertificatePool {
		cfg.RootCAs = config.CertificatePool
	}
	req, err := http.NewRequest(requestMethod, url, nil)
	check(err)
	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Kbn-Xsrf", "true")
	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	check(err)
	if resp.StatusCode != 200 {
		fmt.Println(string(bodyBytes))
	}
	return bodyBytes

}

func getFields(config Config) (fieldList *FieldList) {

	url := fmt.Sprintf("https://%s/api/index_patterns/_fields_for_wildcard?pattern=%s&meta_fields=_source&meta_fields=_id&meta_fields=_type&meta_fields=_index&meta_fields=_score", config.KibanaURL, config.KibanaIndex)
	fieldList = &FieldList{}

	bodyBytes := curl(config, url, "GET")

	json.Unmarshal(bodyBytes, fieldList)

	return fieldList
}

func submitPayload(config Config, body io.Reader) {

	url := fmt.Sprintf("https://%s/api/saved_objects/index-pattern/%s", config.KibanaURL, config.KibanaIndex)
	curl(config, url, "PUT")
}
