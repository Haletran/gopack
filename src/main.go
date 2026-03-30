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
		isRoot() // might need to check later if rootless install
		InstallCommand(command_args[2])
	case "uninstall":
		isRoot()
		UninstallCommand(command_args[2])
	case "search":
		SearchCommand(command_args[2])
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
