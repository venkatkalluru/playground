package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type App struct {
	Name    string
	Parents []string
}

type Group struct {
	Id   string
	Apps []App
}

type Config struct {
	Id     string
	Groups []Group
}

func main() {
	numArgs := os.Args
	if len(numArgs) < 2 {
		fmt.Println("expecting filename as a parameter")
		return
	}
	fileName := os.Args[1]
	c, err := readJSONFile(fileName)
	if err != nil {
		fmt.Println("Got error", err)
		return
	}
	fmt.Println(c.Id)
	fmt.Println(c.Groups)
}

func readJSONFile(fileName string) (Config, error) {

	var c Config
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file", fileName)
		return c, err
	}
	//fmt.Println("Data is ", string(data))
	err = json.Unmarshal(data, &c)
	if err != nil {
		fmt.Println("Error during unmarshalling json ", err)
		return c, err
	}
	return c, nil
}
