package main

import (
	"fmt"
	"os"
)

func command_parser(command_args []string) {
	index := len(command_args)

	if index < 3 {
		fmt.Println(command_args[1], "command is missing an arguments")
		os.Exit(1)
	}
	switch command_args[1] {
	case "install":
		fmt.Println("Installing", command_args[2])
	case "uninstall":
		fmt.Println("Uninstalling", command_args[2])
	case "search":
		fmt.Println("Searching for", command_args[2])
	default:
		fmt.Println("Unknown command", command_args[1])
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing command")
		os.Exit(1)
	}
	command_parser(os.Args)
}
