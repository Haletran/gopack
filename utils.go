package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func CheckDatabase() {
	_, err := os.Open(DATABASE_PATH)
	if err != nil {
		fmt.Println("Index not found, Downloading the list of all packages...")
		// maybe consider having multiple sources of database instead of only one ?? or add possibility to change the source
		resp, err := http.Get("https://raw.githubusercontent.com/Haletran/gopack/refs/heads/main/packages/database.json")
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
