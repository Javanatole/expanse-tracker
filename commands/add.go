package commands

import (
	"expanse-tracker/expanse"
	"flag"
	"fmt"
	"os"
)

// Add command
func Add() {
	// create a new flag based on add command
	addFlags := flag.NewFlagSet("add", flag.ExitOnError)
	// configure both options
	desc := addFlags.String("description", "", "description")
	amount := addFlags.Float64("amount", 0.0, "amount")
	// parse flags from add command
	err := addFlags.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("Parse didn't work with add command")
	}
	// detect if desc or amount are missing
	if *desc == "" {
		fmt.Println("Please provide a description")
		return
	}
	if *amount == 0.0 {
		fmt.Println("Please provide a valid amount")
		return
	}
	service := expanse.GetExpansesService()
	err = service.AddExpanse(*desc, *amount)
	if err != nil {
		fmt.Println("Can't add expense to bdd")
	}
}
