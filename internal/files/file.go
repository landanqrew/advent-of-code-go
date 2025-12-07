package files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)


func ReadFile(path string) ([]byte, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return bytes, nil
}

func WriteFile(path string, bytes []byte) error {
	os.MkdirAll(filepath.Dir(path), 0755)
	err := os.WriteFile(path, bytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

func GetInputFromFile(year int, day int) string {
	path := fmt.Sprintf("input/%d/day_%02d.txt", year, day)
	bytes, err := ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read input file for year %d, day %d: %v", year, day, err)
	}
	return string(bytes)
}