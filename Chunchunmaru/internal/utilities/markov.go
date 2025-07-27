package utilities

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mb-14/gomarkov"
	"os"
	"strings"
)

var MarkovModel *gomarkov.Chain

func TrainMarkovModel(data string, maxOrder int, minSamplesPerState float64, existingChain *gomarkov.Chain) *gomarkov.Chain {
	var chain *gomarkov.Chain
	if existingChain != nil {
		chain = existingChain
	} else {
		chain = gomarkov.NewChain(autoOrder(data, maxOrder, minSamplesPerState))
	}
	for _, sentence := range strings.Split(data, ".") {
		sentence = strings.TrimSpace(sentence)
		if sentence == "" {
			continue
		}
		chain.Add(strings.Fields(sentence))
	}
	return chain
}

func SaveMarkovModel(chain *gomarkov.Chain) {
	jsonObj, _ := json.Marshal(chain)
	err := os.WriteFile("model.json", jsonObj, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func LoadMarkovModel() (*gomarkov.Chain, error) {
	var chain gomarkov.Chain
	data, err := os.ReadFile("model.json")
	if err != nil {
		return &chain, err
	}
	err = json.Unmarshal(data, &chain)
	if err != nil {
		return &chain, err
	}
	MarkovModel = &chain
	return &chain, nil
}

func getDataset(inputData string) []string {
	reader := strings.NewReader(inputData)
	scanner := bufio.NewScanner(reader)
	var list []string
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	return list
}

func split(str string) []string {
	return strings.Fields(str) // Split into words
}

func countNGrams(tokens []string, n int) map[string]int {
	ngrams := make(map[string]int)
	for i := 0; i <= len(tokens)-n; i++ {
		ngram := strings.Join(tokens[i:i+n], " ")
		ngrams[ngram]++
	}
	return ngrams
}

func autoOrder(input string, maxOrder int, minSamplesPerState float64) int {
	tokens := strings.Fields(input)
	bestOrder := 1
	for order := 1; order <= maxOrder; order++ {
		ngrams := countNGrams(tokens, order)
		if len(ngrams) == 0 {
			break
		}
		avgSamples := float64(len(tokens)-order+1) / float64(len(ngrams))
		if avgSamples < minSamplesPerState {
			break
		}
		bestOrder = order
	}
	return bestOrder
}
