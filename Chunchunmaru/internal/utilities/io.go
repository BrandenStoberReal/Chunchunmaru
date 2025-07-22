package utilities

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// RandomHTMLFromDir selects a random .html file from dir and returns its contents as a string alongside the file name.
func RandomHTMLFromDir(dir string) (string, string, error) {
	files, err := ioutil.ReadDir(dir)
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

	rand.Seed(time.Now().UnixNano())
	randomFile := htmlFiles[rand.Intn(len(htmlFiles))]

	content, err := os.ReadFile(randomFile)
	if err != nil {
		return "", "", err
	}

	return string(content), path.Base(randomFile), nil
}
