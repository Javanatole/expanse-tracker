package expanse

import (
	"fmt"
	"os"
)

// FileStorage interface for file operations
type FileStorage interface {
	Read() (string, error)
	Write(content string) error
}

// FileTaskStorage handles file-based storage
type FileTaskStorage struct {
	Filename string
}

// Read content from file
func (fileTaskStorage *FileTaskStorage) Read() (string, error) {
	content, err := os.ReadFile(fileTaskStorage.Filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(content), nil
}

// Write content into file
func (fileTaskStorage *FileTaskStorage) Write(content string) error {
	return os.WriteFile(fileTaskStorage.Filename, []byte(content), 0644)
}
