package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func validatePath(path string) error {
	cleanPath := filepath.Clean(path)
	if filepath.IsAbs(cleanPath) {
		return fmt.Errorf("absolute paths are not allowed")
	}
	parts := strings.Split(cleanPath, string(os.PathSeparator))
	for _, part := range parts {
		if part == ".." {
			return fmt.Errorf("directory traversal is not allowed")
		}
	}
	return nil
}
