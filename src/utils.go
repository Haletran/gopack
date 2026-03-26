package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func verifyChecksum(filePath string, expected string) error {
	expected = strings.TrimPrefix(expected, "sha256:")
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return err
	}

	actual := hex.EncodeToString(h.Sum(nil))
	if actual != expected {
		return fmt.Errorf("invalid checksum")
	}
	return nil
}

func CheckDatabase() {
	_, err := os.Open(DATABASE_PATH)
	if err != nil {
		fmt.Println("Index not found, Downloading the list of all packages...")
		// TODO : maybe consider having multiple sources of database instead of only one ?? or add possibility to change the source
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
