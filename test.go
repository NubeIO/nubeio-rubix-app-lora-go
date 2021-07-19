package main

import (
	"encoding/json"
	"fmt"
	model "github.com/NubeIO/nubeio-rubix-app-lora-go/model/networks"
	"io/ioutil"
	"log"
	"net/http"
)


//type Todo struct {
//	response
//	Status string `json:"status"`
//}

type response struct {
	response model.Network
}


func parseMap(aMap map[string]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println(key)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println(key)
			parseArray(val.([]interface{}))
		default:
			fmt.Println(key, ":", concreteVal)

		}
	}
}

func parseArray(anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println("Index:", i)
			parseArray(val.([]interface{}))
		default:
			fmt.Println("Index", i, ":", concreteVal)

		}
	}
}

func get() {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("http://0.0.0.0:1920/api/network/?uuid=072505efe2b643d0")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	m := map[string]interface{}{}

	// Parsing/Unmarshalling JSON encoding/json

	// Convert response body to string
	bodyString := string(bodyBytes)

	err = json.Unmarshal([]byte(bodyString), &m)

	if err != nil {
		panic(err)
	}


	parseMap(m)
	//fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	//var todoStruct response
	//json.Unmarshal(bodyBytes, &todoStruct)
	//
	//fmt.Println(3333, todoStruct)
	//fmt.Printf("API Response as struct %+v\n", todoStruct)
}

func main() {
	get()
}
//	resp, err := http.Get("http://0.0.0.0:1920/api/networks")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	//We Read the response body on the line below.
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	//Convert the body to type string
//	sb := string(body)
//	log.Printf(sb)
//}