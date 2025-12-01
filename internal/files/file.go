package files

import (
	"fmt"
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