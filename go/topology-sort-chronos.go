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
	//printAppAndParents(&c)

	if len(c.Groups) == 0 {
		fmt.Println("Empty Groups")
		return
	}
	for _, g := range c.Groups {

		sortedAppNames := topoSort(g.Apps)

		fmt.Println("Sorted Apps for group ", g.Id)
		for _, name := range sortedAppNames {
			fmt.Println(name)
		}
		fmt.Println("--------------")
	}
}

func topoSort(apps []App) []string {

	appsMap := make(map[string]App, len(apps))

	for _, app := range apps {
		appsMap[app.Name] = app
	}

	visitedApps := make(map[string]bool, len(apps))
	sortedAppNames := make([]string, 0, len(apps))
	for _, app := range apps {

		sortedAppNames = recurTopoSort(app, appsMap, visitedApps, sortedAppNames)
	}

	return sortedAppNames
}

func recurTopoSort(app App, appsMap map[string]App, visitedApps map[string]bool, sortedAppNames []string) []string {

	for _, p := range app.Parents {

		if !visitedApps[p] {
			sortedAppNames = recurTopoSort(appsMap[p], appsMap, visitedApps, sortedAppNames)
			visitedApps[p] = true
		}
	}
	if !visitedApps[app.Name] {
		visitedApps[app.Name] = true
		sortedAppNames = append(sortedAppNames, app.Name)
	}
	return sortedAppNames
}

func printAppAndParents(c *Config) {
	for _, g := range c.Groups {
		for _, a := range g.Apps {
			fmt.Printf("%s : ", a.Name)
			fmt.Println(a.Parents)
		}
	}
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
