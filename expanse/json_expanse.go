package expanse

import (
	"bytes"
	"errors"
	"fmt"
)

// Expanse represents a single expanse
type Expanse struct {
	Id          int     `json:"id"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

// String show expanse representation
func (expanse Expanse) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("ID: %d\n", expanse.Id))
	buffer.WriteString(fmt.Sprintf("Description: %s\n", expanse.Description))
	buffer.WriteString(fmt.Sprintf("Updated: %s\n", expanse.Date))
	buffer.WriteString(fmt.Sprintf("Status: %s\n", expanse.Amount))
	return buffer.String()
}

// JSONExpanses represents a collection of expanse
type JSONExpanses struct {
	Expanses []Expanse
}

// String representation of all tasks
func (expanses *JSONExpanses) String() string {
	var buffer bytes.Buffer
	for _, expanse := range expanses.Expanses {
		buffer.WriteString(expanse.String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (expanses *JSONExpanses) FindLastId() int {
	maxId := 0
	for _, expanse := range expanses.Expanses {
		if expanse.Id > maxId {
			maxId = expanse.Id
		}
	}
	return maxId
}

func (expanses *JSONExpanses) FindIndex(id int) (int, error) {
	for i, expanse := range expanses.Expanses {
		if expanse.Id == id {
			return i, nil
		}
	}
	return -1, errors.New("task not found")
}

func (expanses *JSONExpanses) DeleteElement(index int) {
	expanses.Expanses = append(expanses.Expanses[:index], expanses.Expanses[index+1:]...)
}
