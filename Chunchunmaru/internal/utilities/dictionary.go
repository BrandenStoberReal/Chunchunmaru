package utilities

import (
	"bufio"
	"bytes"
	_ "embed"
	"math/rand"
)

//go:embed words.txt
var embeddedWords []byte

var dict []string

var AvgWordLen int

func init() {
	// Read words from the embedded words.txt
	reader := bytes.NewReader(embeddedWords)
	var words []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	dict = words
	AvgWordLen = int(len(embeddedWords) / WordCount())
}

func RandomWord() string {
	if len(dict) == 0 {
		return ""
	}
	return dict[rand.Intn(len(dict))]
}

func WordCount() int {
	return len(dict)
}
