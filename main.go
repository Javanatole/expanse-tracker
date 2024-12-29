package main

import (
	"expanse-tracker/commands"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		log.Fatal("Please specify a subcommand.")
	} else {
		switch os.Args[1] {
		case "add":
			commands.Add()
		case "delete":
			commands.Delete()
		case "list":
			commands.List()
		case "summary":
			commands.Summary()
		}
	}
}
