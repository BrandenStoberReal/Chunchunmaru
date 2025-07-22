package macros

import (
	"chunchunmaru/internal/utilities"
	"encoding/json"
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

// internalGenerateRandomJson creates a random JSON structure with a specified depth.
// The returned interface{} can be a map[string]interface{}, a slice of interface{},
// a string, an integer, or a boolean.
func internalGenerateRandomJson(depth, maxElements, maxStringLength int) interface{} {
	if depth <= 0 {
		// Base case: return a random primitive type
		switch rand.Intn(3) {
		case 0:
			return utilities.RandomStringFromCharset(rand.Intn(maxStringLength)+1, utilities.MixedDigitChars)
		case 1:
			return rand.Intn(100)
		default:
			return rand.Intn(2) == 1
		}
	}

	// Recursive step: create a map or a slice
	if rand.Intn(2) == 0 {
		// Create a JSON object
		obj := make(map[string]interface{})
		numElements := rand.Intn(maxElements) + 1
		for i := 0; i < numElements; i++ {
			key := utilities.RandomStringFromCharset(5, utilities.MixedDigitChars)
			obj[key] = internalGenerateRandomJson(depth-1, maxElements, maxStringLength)
		}
		return obj
	}

	// Create a JSON array
	arr := make([]interface{}, 0)
	numElements := rand.Intn(maxElements) + 1
	for i := 0; i < numElements; i++ {
		arr = append(arr, internalGenerateRandomJson(depth-1, maxElements, maxStringLength))
	}
	return arr

}

// randomLink Provies a randomly generated URL
func randomLink() string {
	config := utilities.AppConfig.GetConfig()
	word := randomWord()
	for ; slices.Contains(config.PathWhitelist, "/"+word); word = randomWord() {
		// Logic is in the loop header lmao
	}

	var builder strings.Builder
	builder.WriteString(config.HostName)
	builder.WriteString("/")
	builder.WriteString(word)
	builder.WriteString("/")
	pathNum := rand.Intn(config.MaxSubpaths-config.MinSubpaths) + config.MinSubpaths - 1
	for _ = range pathNum {
		builder.WriteString(randomWord())
		builder.WriteString("/")
	}
	return builder.String()
}

// randomQueryLink Provides a randomly generated query URL
func randomQueryLink(keyCount int) string {
	if keyCount == 1 {
		return randomLink() + "?" + randomWord() + "=" + utilities.RandomStringFromCharset(rand.Intn(16)+5, utilities.AlphabetChars)
	} else {
		var builder strings.Builder
		builder.WriteString(randomLink())
		builder.WriteString("?")
		for i := 0; i < keyCount; i++ {
			if i != 0 {
				builder.WriteString("&")
			}
			builder.WriteString(randomWord())
			builder.WriteByte('=')
			builder.WriteString(utilities.RandomStringFromCharset(rand.Intn(16)+5, utilities.AlphabetChars))
		}

		return builder.String()
	}
}

// randomJSON Generates a random nested JSON object
func randomJSON(depth, maxElements, maxStringLength int) (string, error) {
	if depth < 0 {
		return "", fmt.Errorf("depth cannot be negative")
	}

	randomData := internalGenerateRandomJson(depth, maxElements, maxStringLength)

	jsonData, err := json.MarshalIndent(randomData, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return string(jsonData), nil
}
