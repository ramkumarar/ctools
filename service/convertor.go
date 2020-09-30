package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// MapJSONToPropsByteArray maps json bytearray to properties bytearray
func MapJSONToPropsByteArray(jsonData []byte) []byte {

	var v interface{}
	json.Unmarshal(jsonData, &v)

	data := v.(map[string]interface{})
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

func WriteToFile(fileName string, bytes []byte) error {
	err := ioutil.WriteFile(fileName, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
