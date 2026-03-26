package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

var DATABASE_PATH string = "packages/database.json"

func checkdatabase() {
	_, err := os.Open(DATABASE_PATH)
	if err != nil {
		fmt.Println("Index not found, Downloading the list of all packages...")
		// maybe consider having multiple sources of database instead of only one ?? or add possibility to change the source
		resp, err := http.Get("https://raw.githubusercontent.com/Haletran/gopack/refs/heads/main/database.json?token=GHSAT0AAAAAADTS6R6DASPL3CESDOSRM4KI2OBLEWQ")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		f, err := os.Create(DATABASE_PATH)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := io.Copy(f, resp.Body); err != nil {
			log.Fatal(err)
		}
	}
}

func SearchCommand(package_name string) {
	checkdatabase()
	data, err := os.ReadFile(DATABASE_PATH)
	if err != nil {
		log.Fatal(err)
	}

	var index Index
	json.Unmarshal(data, &index) // TODO : need to add error check here in case of a wrong json file

	var iteration int
	for _, pkg := range index.Packages {
		if strings.Contains(pkg.Name, package_name) {
			iteration++
		}
	}

	fmt.Println(iteration, "result found")
	for _, pkg := range index.Packages {
		if strings.Contains(pkg.Name, package_name) {
			fmt.Println("---")
			fmt.Printf("Name: %s\n", pkg.Name)
			fmt.Printf("Description: %s\n", pkg.Description)
			fmt.Printf("Author: %s\n", pkg.Author)
			fmt.Printf("Latest: %s\n", pkg.Latest)
			fmt.Printf("Versions: %v\n", pkg.Versions)
			fmt.Println("---")
		}
	}
}
