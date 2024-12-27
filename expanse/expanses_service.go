package expanse

import (
	"time"
)

func GetExpansesService() *ExpansesService {
	// File name to store the tasks
	filename := "expanses.json"
	// Default content if the file is empty or doesn't exist
	defaultContent := "{\"expanses\":[]}"

	// Create storage and repository
	fileStorage := &FileTaskStorage{Filename: filename}
	jsonRepository := &JSONExpansesRepository{
		Storage:        fileStorage,
		DefaultContent: defaultContent,
	}

	// Initialize the TaskService
	expansesService := &ExpansesService{
		Repo: jsonRepository,
	}
	return expansesService
}

// ExpansesService manages expanses operations
type ExpansesService struct {
	Repo JSONStorage
}

func (expansesService *ExpansesService) AddExpense(description string, amount float64) error {
	expanses, err := expansesService.Repo.Load()
	if err != nil {
		return err
	}
	expanse := Expanse{
		Id:          expanses.FindLastId() + 1,
		Description: description,
		Amount:      amount,
		Date:        time.Now().Format(time.RFC3339),
	}
	expanses.Expanses = append(expanses.Expanses, expanse)

	return expansesService.Repo.Save(expanses)
}

func (expansesService *ExpansesService) DeleteExpense(id int) error {
	expanses, err := expansesService.Repo.Load()
	if err != nil {
		return err
	}
	index, err := expanses.FindIndex(id)
	if err != nil {
		return err
	}
	expanses.DeleteElement(index)
	return expansesService.Repo.Save(expanses)
}
