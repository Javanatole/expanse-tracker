package expanse

import (
	"time"
)

// GetExpansesService generate expanses service
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

// AddExpanse add expanse with his description and amount
func (expansesService *ExpansesService) AddExpanse(description string, amount float64) error {
	expanses, err := expansesService.Repo.Load()
	if err != nil {
		return err
	}
	expanse := Expanse{
		Id:          expanses.FindLastId() + 1,
		Description: description,
		Amount:      amount,
		Date:        time.Now().Format(time.DateTime),
	}
	expanses.Expanses = append(expanses.Expanses, expanse)

	return expansesService.Repo.Save(expanses)
}

// DeleteExpanse delete specific expanse
func (expansesService *ExpansesService) DeleteExpanse(id int) error {
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

func (expansesService *ExpansesService) ListExpanses() ([]Expanse, error) {
	expanses, err := expansesService.Repo.Load()
	if err != nil {
		return []Expanse{}, err
	}
	return expanses.Expanses, nil
}
