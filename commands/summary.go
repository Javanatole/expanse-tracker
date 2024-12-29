package commands

import (
	"expanse-tracker/expanse"
	"flag"
	"fmt"
	"os"
	"time"
)

func Summary() {
	// create a new flag based on summary command
	monthFlags := flag.NewFlagSet("summary", flag.ExitOnError)
	// configure month option
	month := monthFlags.Int("month", -1, "month to summary")
	// parse flags from summary command
	err := monthFlags.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("Parse didn't work with summary command")
	}
	service := expanse.GetExpansesService()
	expanses, err := service.ListExpanses()
	if err != nil {
		panic(err)
	}
	// detect if id is missing
	if *month == -1 {
		totalExpanses := 0.0
		for _, element := range expanses {
			totalExpanses += element.Amount
		}
		fmt.Println(fmt.Sprintf("Total Expanses: $%v", totalExpanses))
	} else {
		totalExpanses := 0.0
		for _, element := range expanses {
			formattedLayout, err := time.Parse(time.DateTime, element.Date)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Failed to parse date")
				return
			}
			if int(formattedLayout.Month()) == *month {
				totalExpanses += element.Amount
			}
		}
		fmt.Println(fmt.Sprintf("Total Expanses for month %d: $%v", *month, totalExpanses))

	}
}
