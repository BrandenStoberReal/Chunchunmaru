package utilities

import (
	"fmt"
	"github.com/mb-14/gomarkov"
	"strings"
	"testing"
)

const dataset = "This may not be a formal definition, but when using Markov models for predictive data compression or speech recognition, the order is generally the number of previously seen symbols that are considered when calculating the probabilities for the next symbol.\n\nAs an example, if I was using an order-5 markov model to predict individual letters in a text stream, and the previous five characters were BARAC, the probability of letter K being next might be 93%, A might be 2%, H might bet 4%, and so on.\n\nIncreasing the order of the model usually improves the quality of the predictions for the next symbol, resulting in better speech recognition or higher compression ratios."

func printTestResults(name, result interface{}) {
	fmt.Printf("--- BENCHMARK FINISHED ---\n")
	fmt.Printf("Function:    %s\n", name)
	fmt.Printf("Last Result: %s\n", result)
	fmt.Printf("--------------------------\n")
}

func TestChainGenerate(t *testing.T) {
	chain := TrainMarkovModel(dataset, 5, 2.0)
	order := chain.Order
	tokens := make([]string, 0)
	for i := 0; i < order; i++ {
		tokens = append(tokens, gomarkov.StartToken)
	}
	for tokens[len(tokens)-1] != gomarkov.EndToken && len(tokens) < 80 {
		next, _ := chain.Generate(tokens[(len(tokens) - order):])
		tokens = append(tokens, next)
	}
	printTestResults("Train & Test Markov Chain", strings.Join(tokens[order:len(tokens)-1], " "))
}
