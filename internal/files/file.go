package files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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

// FindDirectoryBFS searches for a directory matching targetPath starting from the system root (/) using BFS.
// targetPath can be a single directory name or a joined path (e.g., "usr/local/bin").
// It returns the absolute path of the first matching directory found.
func FindDirectoryBFS(targetPath string) (string, error) {
	if targetPath == "" {
		return "", fmt.Errorf("target path cannot be empty")
	}

	// Start at system root
	root := "/"

	// Clean and normalize target path
	targetPath = filepath.Clean(targetPath)
	if !filepath.IsAbs(targetPath) {
		// If targetPath is relative, we'll check if fullPath ends with it
		if !strings.HasPrefix(targetPath, string(filepath.Separator)) {
			targetPath = string(filepath.Separator) + targetPath
		}
	}

	queue := []string{root}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}
		visited[current] = true

		entries, err := os.ReadDir(current)
		if err != nil {
			continue // Skip directories we can't read (Permission Denied, etc.)
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}

			fullPath := filepath.Join(current, entry.Name())

			// Match if the full path is exactly targetPath OR if it ends with the targetPath sequence
			if fullPath == targetPath || strings.HasSuffix(fullPath, targetPath) {
				return fullPath, nil
			}

			queue = append(queue, fullPath)
		}
	}

	return "", fmt.Errorf("directory %q not found starting from system root", targetPath)
}

