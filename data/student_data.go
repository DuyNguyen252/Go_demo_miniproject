package data

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func WriteInfoStudent(dir string, fileName string, content string) error {
	fullPath := filepath.Join(dir, fileName)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	err = os.WriteFile(fullPath, []byte(content), 0644)
	if err != nil {
		log.Fatalf("Unable to write file %s: %v", fullPath, err)
	}

	fmt.Printf("File successfully written to %s\n", fullPath)
	return nil
}

func ReadInfoStudent(dir string, fileName string) (string, error) {
	fullPath := filepath.Join(dir, fileName)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	return string(data), nil
}
