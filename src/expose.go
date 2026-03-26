package main

import (
	"fmt"
	"log"

	"github.com/cavaliergopher/grab/v3"
)

func Download(fileUrl string) {
	resp, err := grab.Get("/tmp", fileUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Download saved to", resp.Filename)
}
