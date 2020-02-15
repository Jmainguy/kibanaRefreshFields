package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func formatPayload(fieldList *FieldList) (body io.Reader) {
	kibanaFilter := os.Getenv("KIBANA_FILTER")
	payload := Payload{}
	payload.Attributes.Title = "filebeat-*"
	payload.Attributes.TimeFieldName = "@timestamp"

	payload.Attributes.Fields = "["
	if kibanaFilter != "" {
		fmt.Printf("Filtering out %s\n", kibanaFilter)
    }
	for _, v := range fieldList.Fields {
		// If you want to filter out some fields
		if kibanaFilter != "" {
			if !strings.Contains(v.Name, kibanaFilter) {
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

func getFields(username, password, kibanaUrl, kibanaIndex string) (fieldList *FieldList) {

	url := fmt.Sprintf("https://%s/api/index_patterns/_fields_for_wildcard?pattern=%s&meta_fields=_source&meta_fields=_id&meta_fields=_type&meta_fields=_index&meta_fields=_score", kibanaUrl, kibanaIndex)

	fieldList = &FieldList{}

	req, err := http.NewRequest("GET", url, nil)
	check(err)
	req.SetBasicAuth(username, password)
	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	check(err)
	json.Unmarshal(bodyBytes, fieldList)

	return fieldList
}

func submitPayload(username, password, kibanaUrl, kibanaIndex string, body io.Reader) {

	url := fmt.Sprintf("https://%s/api/saved_objects/index-pattern/%s", kibanaUrl, kibanaIndex)

	req, err := http.NewRequest("PUT", url, body)
	check(err)
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Kbn-Xsrf", "true")

	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	check(err)
	bodyString := string(bodyBytes)

	fmt.Println(resp.StatusCode)

	if resp.StatusCode != 200 {
		fmt.Println(bodyString)
	}

}
