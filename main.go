package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	username := os.Getenv("KIBANA_USERNAME")
	password := os.Getenv("KIBANA_PASSWORD")
	kibanaUrl := os.Getenv("KIBANA_URL")
	kibanaIndex := os.Getenv("KIBANA_INDEX")

	fieldList := getFields(username, password, kibanaUrl, kibanaIndex)

	body := formatPayload(fieldList)

	submitPayload(username, password, kibanaUrl, kibanaIndex, body)
}
