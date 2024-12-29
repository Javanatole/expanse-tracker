package expanse

import (
	"encoding/json"
	"fmt"
)

// JSONStorage interface for JSON operations
type JSONStorage interface {
	Save(expanses JSONExpanses) error
	Load() (JSONExpanses, error)
}

// JSONExpansesRepository handles JSON-specific operations
type JSONExpansesRepository struct {
	Storage        FileStorage
	DefaultContent string
}

func (jsonExpansesRepository *JSONExpansesRepository) Save(expanses JSONExpanses) error {
	data, err := json.Marshal(expanses)
	if err != nil {
		return fmt.Errorf("failed to marshal expanses: %w", err)
	}
	return jsonExpansesRepository.Storage.Write(string(data))
}

func (jsonExpansesRepository *JSONExpansesRepository) Load() (JSONExpanses, error) {
	content, err := jsonExpansesRepository.Storage.Read()
	if err != nil {
		if writeErr := jsonExpansesRepository.Storage.Write(jsonExpansesRepository.DefaultContent); writeErr != nil {
			return JSONExpanses{}, fmt.Errorf("failed to write default content: %w", writeErr)
		}
		content = jsonExpansesRepository.DefaultContent
	}

	var expanses JSONExpanses
	err = json.Unmarshal([]byte(content), &expanses)
	if err != nil {
		return JSONExpanses{}, fmt.Errorf("failed to parse JSON content: %w", err)
	}
	return expanses, nil
}
