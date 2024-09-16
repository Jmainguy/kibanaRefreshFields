package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
)

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

	bodyBytes, err = io.ReadAll(resp.Body)
	check(err)
	if resp.StatusCode != 200 {
		fmt.Println(string(bodyBytes))
	}
	return bodyBytes

}

func submitPayload(config Config) {

	url := fmt.Sprintf("https://%s/api/saved_objects/index-pattern/%s", config.KibanaURL, config.KibanaIndex)
	curl(config, url, "PUT")
}
