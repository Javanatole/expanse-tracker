package commands

import (
	"expanse-tracker/expanse"
	"flag"
	"fmt"
	"os"
)

func Delete() {
	// create a new flag based on add command
	deleteFlags := flag.NewFlagSet("delete", flag.ExitOnError)
	// configure id option
	id := deleteFlags.Int("id", 0, "id to delete")
	// parse flags from delete command
	err := deleteFlags.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("Parse didn't work with add command")
	}
	// detect if id is missing
	if *id == 0 {
		fmt.Println("Please provide a valid id")
		return
	}
	service := expanse.GetExpansesService()
	err = service.DeleteExpense(*id)
	if err != nil {
		fmt.Println("Can't add expense to bdd")
	}
}
