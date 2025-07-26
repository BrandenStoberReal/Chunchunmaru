package utilities

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// RandomHTMLFromDir selects a random .html file from dir and returns its contents as a string alongside the file name.
func RandomHTMLFromDir(dir string) (string, string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", "", err
	}

	var htmlFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".html") {
			htmlFiles = append(htmlFiles, filepath.Join(dir, file.Name()))
		}
	}

	if len(htmlFiles) == 0 {
		return "", "", errors.New("no HTML files found in directory")
	}

	randomFile := htmlFiles[rand.Intn(len(htmlFiles))]

	content, err := os.ReadFile(randomFile)
	if err != nil {
		return "", "", err
	}

	return string(content), path.Base(randomFile), nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true // File exists and no error occurred
	}
	if errors.Is(err, os.ErrNotExist) {
		return false // File does not exist
	}
	// Other error occurred (e.g., permissions issue)
	fmt.Printf("Error checking file %s: %v\n", filePath, err)
	return false // Or handle the error as appropriate for your application
}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
