package main

import (
	"fmt"
	"os"
)

func main() {
	// Get CLI arguments
	args := os.Args[1:]

	if len(args) < 1 || args[0] == "help" {
		fmt.Println("No argument supplied.")
		return
	}

	switch args[0] {
	case "add":
		fmt.Println("add")
	case "restore":
		fmt.Println("restore")
	case "backup":
		fmt.Println("backup")
	case "remove":
		fmt.Println("remove")
	default:
		fmt.Printf("Unknown command: %s", args[0])
	}
}
