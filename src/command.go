package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Package struct {
	Name        string
	Description string
	Author      string
	Latest      string
	Versions    []string
}

type Index struct {
	Packages []Package
}

var DATABASE_PATH string = "../packages/database.json"

func getdatabase() Index {
	data, err := os.ReadFile(DATABASE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	var database Index
	json.Unmarshal(data, &database) // TODO : need to add error check here in case of a wrong json file
	return database
}

func searchdatabase(package_name string, database Index) bool {
	for _, pkg := range database.Packages {
		if strings.Contains(pkg.Name, package_name) {
			return true
		}
	}
	return false
}

func printAlldatabase(database Index) {
	for _, pkg := range database.Packages {
		fmt.Println(pkg.Name, pkg.Description, pkg.Versions)
	}
}

func printsearchdb(package_name string, database Index) {
	var nb_result int
	for _, pkg := range database.Packages {
		if strings.Contains(pkg.Name, package_name) {
			fmt.Println(pkg.Name, pkg.Description, pkg.Versions)
			nb_result++
		}
	}
	fmt.Println("--> Found", nb_result, "available packages")
}

func SearchCommand(package_name string) {
	CheckDatabase()
	var index Index = getdatabase()
	if searchdatabase(package_name, index) == true {
		fmt.Println("[", package_name, "]", "found in the database")
		printsearchdb(package_name, index)
	} else {
		fmt.Println(package_name, "not found")
	}
}
