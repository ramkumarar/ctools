package main

import (
	"consul/config"
	"consul/service"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	space, err := config.GetSpace("service1")
	if err != nil {
		fmt.Println(err)
	}

	consulKey, err := config.GetConsulKey("space1", "test1")
	if err != nil {
		fmt.Println(err)
	}

	jsonData := []byte(`{"Name": "Alice", "Age": 25 ,"occupation": "SE", "smoker": false}`)

	err = ioutil.WriteFile("sample.properties", service.MapJSONToPropsByteArray(jsonData), 0644)
	if err != nil {
		log.Fatalln("Not able to write to file")
	}

	fmt.Println(space)
	fmt.Println(consulKey)

}
