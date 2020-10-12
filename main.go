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

	config, err := getConfig()
	check(err)

	fieldList := getFields(config)

	body := formatPayload(fieldList, config)

	submitPayload(config, body)
}
