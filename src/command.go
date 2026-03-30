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
	Modes       []string
}

type Index struct {
	Packages []Package
}

var DATABASE_PATH string = "../packages/database.json"
var INSTALL_PATH string = "/usr/local/bin/"

func getdatabase() Index {
	data, err := os.ReadFile(DATABASE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	var database Index
	json.Unmarshal(data, &database) // TODO : need to add error check here in case of a wrong json file
	return database
}

func searchdatabase(package_name string, database Index) int {
	var iteration int
	for _, pkg := range database.Packages {
		if strings.Contains(pkg.Name, package_name) {
			iteration++
		}
	}
	return iteration
}

func printAlldatabase(database Index) {
	for _, pkg := range database.Packages {
		fmt.Println(pkg.Name, "-", pkg.Description, pkg.Versions)
	}
}

func printsearchdb(package_name string, database Index) {
	var nb_result int
	for _, pkg := range database.Packages {
		if strings.Contains(pkg.Name, package_name) {
			fmt.Println(pkg.Name, "-", pkg.Description, pkg.Versions)
			nb_result++
		}
	}
	fmt.Println("--> Found", nb_result, "available packages")
}

func SearchCommand(package_name string) {
	CheckDatabase()
	var index Index = getdatabase()
	if searchdatabase(package_name, index) >= 1 {
		//fmt.Println("[", package_name, "]", "found in the database")
		printsearchdb(package_name, index)
	} else {
		fmt.Println(package_name, "not found")
	}
}

func InstallCommand(package_name string) {
	luaParser(package_name)
}

func searchBin(package_name string) bool {
	files, err := os.ReadDir(INSTALL_PATH)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if strings.Contains(f.Name(), package_name) {
			return true
		}
	}
	return false
}

func UninstallCommand(package_name string) {
	// TODO : need to add check to the installed db (not setup yet)
	if searchBin(package_name) == true {
		files, err := os.ReadDir(INSTALL_PATH)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			if strings.Contains(f.Name(), package_name) {
				fmt.Println(f.Name())
				err := os.Remove(INSTALL_PATH + f.Name())
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Successfully Uninstalled", f.Name())
			}
		}
	} else {
		fmt.Println("Package", package_name, "not found")
	}
}
