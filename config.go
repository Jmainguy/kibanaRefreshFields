package main

import (
	"crypto/x509"
	"io/ioutil"
	"os"
)

// Config : configuration settings
type Config struct {
	Username        string
	Password        string
	KibanaURL       string
	KibanaIndex     string
	KibanaFilter    string
	Insecure        bool
	Certificate     string
	CertificatePool *x509.CertPool
}

func getConfig() (config Config, err error) {
	config.Username = os.Getenv("KIBANA_USERNAME")
	config.Password = os.Getenv("KIBANA_PASSWORD")
	config.KibanaURL = os.Getenv("KIBANA_URL")
	config.KibanaIndex = os.Getenv("KIBANA_INDEX")
	config.Certificate = os.Getenv("KIBANA_CERT")
	config.KibanaFilter = os.Getenv("KIBANA_FILTER")
	insecureString := os.Getenv("KIBANA_INSECURE")

	if insecureString == "true" {
		config.Insecure = true
	} else {
		config.Insecure = false
	}

	if config.Certificate != "" {

		certificate, err := ioutil.ReadFile(config.Certificate)
		if err != nil {
			return config, err
		}
		config.CertificatePool = x509.NewCertPool()
		config.CertificatePool.AppendCertsFromPEM(certificate)
	}

	return config, err
}
