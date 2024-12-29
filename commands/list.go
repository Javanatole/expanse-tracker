package commands

import (
	"expanse-tracker/expanse"
	"fmt"
	"time"
)

func List() {
	service := expanse.GetExpansesService()
	expanses, err := service.ListExpanses()
	if err != nil {
		panic(err)
	}
	fmt.Println("# ID  Date       Description  Amount")
	for _, element := range expanses {
		formattedLayout, err := time.Parse(time.DateTime, element.Date)
		if err != nil {
			fmt.Println("Failed to parse date")
			return
		}
		fmt.Println(fmt.Sprintf(
			"# %d   %s %s %f",
			element.Id,
			formattedLayout.Format(time.DateOnly),
			element.Description,
			element.Amount))
	}
}
