package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	numArgs := os.Args
	if len(numArgs) < 2 {
		fmt.Println("expecting filename as a parameter")
		return
	}
	fileName := os.Args[1]
	dataMap, err := readJSONFile(fileName)
	if err != nil {
		fmt.Println("Got error", err)
		return
	}
	groups := dataMap["groups"]
	if groups == nil {
		fmt.Println("Empty groups, hence returing")
		return
	}

	for _, group := range groups.([]interface{}) {
		groupMap := group.(map[string]interface{})
		fmt.Println("Group id is ", groupMap["id"])
	}
}

func readJSONFile(fileName string) (map[string]interface{}, error) {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file", fileName)
		return nil, err
	}
	//fmt.Println("Data is ", string(data))
	var intf interface{}
	err = json.Unmarshal(data, &intf)
	if err != nil {
		fmt.Println("Error during unmarshalling json ", err)
		return nil, err
	}
	return intf.(map[string]interface{}), nil
}
