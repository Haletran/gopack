package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/cavaliergopher/grab/v3"
	lua "github.com/yuin/gopher-lua"
)

var PACKAGES_PATH string = "../packages"

func findLuaFile(package_name string) string {
	// TODO: do that online not in a file for update this doesnt make any sense
	return fmt.Sprintf("%s/%s/%s.lua", PACKAGES_PATH, package_name, package_name)
}

func Download(fileUrl string, checksum string) {
	resp, err := grab.Get("/tmp", fileUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Download saved to", resp.Filename)

	if checksum != "" {
		fmt.Println("--> Verifying checksum...")
		if err := verifyChecksum(resp.Filename, checksum); err != nil {
			os.Remove(resp.Filename)
			log.Fatal(err)
		}
		fmt.Println("Checksum verified")
	}
}

func Extract(source string, destination string) {
	var cmd *exec.Cmd
	// TODO: should do a native solution instead of using the user shell
	if strings.HasSuffix(source, ".zip") {
		cmd = exec.Command("unzip", "-o", source, "-d", destination)
	} else {
		cmd = exec.Command("tar", "-xf", source, "-C", destination)
	}
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Extracted %s to %s\n", source, destination)
}

func Install(source string, destination string) {
	ChangePermisions(source, 0755)
	// TODO: might need to make a much better solution than this shit
	cmd := exec.Command("mv", source, destination)
	err2 := cmd.Run()
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("Sucessfully installed")
	os.Remove(source)
	// TODO : if successfull then keep track of it to not install it multiple times
}

func luaParser(package_name string) {
	var filePath string = findLuaFile(package_name)
	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("describe", L.NewFunction(func(L *lua.LState) int {
		tbl := L.CheckTable(1)
		urlVal := tbl.RawGetString("url")
		sha256Val := tbl.RawGetString("sha256")
		L.SetGlobal("url", urlVal)
		L.SetGlobal("sha256", sha256Val)
		name := tbl.RawGetString("name")
		version := tbl.RawGetString("version")
		fmt.Printf("--> Package : %s version %s\n", name, version)
		return 0
	}))

	L.SetGlobal("Extract", L.NewFunction(func(L *lua.LState) int {
		source := L.CheckString(1)
		destination := L.CheckString(2)
		fmt.Printf("--> Extracting %s to %s\n", source, destination)
		Extract(source, destination)
		return 0
	}))

	L.SetGlobal("Install", L.NewFunction(func(L *lua.LState) int {
		source := L.CheckString(1)
		destination := L.CheckString(2)
		fmt.Printf("--> Installing %s to %s\n", source, destination)
		Install(source, destination)
		return 0
	}))

	L.SetGlobal("Download", L.NewFunction(func(L *lua.LState) int {
		url := L.CheckString(1)
		checksum := L.OptString(2, "")
		fmt.Println("--> Downloading from", url)
		Download(url, checksum)
		return 0
	}))
	if err := L.DoFile(filePath); err != nil {
		panic(err)
	}
}
