package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {

	jsonData := []byte(`{"Name": "Alice", "Age": 25 ,"occupation": "SE", "smoker": false}`)

	var v interface{}
	json.Unmarshal(jsonData, &v)

	err := writeToFile("sample.properties", mapInterfaceToByteArray(v))

	if err != nil {
		log.Fatalln("Unable to write a property file")
	}

}

func mapInterfaceToByteArray(response interface{}) []byte {

	data := response.(map[string]interface{})
	b := new(bytes.Buffer)

	for key, value := range data {
		fmt.Fprintf(b, "%s=\"%v\"\n", key, value)
	}

	propStr := b.String()
	propStr = strings.ReplaceAll(propStr, `="`, "=")

	re := regexp.MustCompile(`"\r?\n`)
	propStr = re.ReplaceAllString(propStr, "\n")

	bytes := []byte(propStr)

	return bytes

}

func writeToFile(fileName string, bytes []byte) error {
	err := ioutil.WriteFile(fileName, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
