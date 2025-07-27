package macros

import (
	"fmt"
	"testing"
)

func printTestResults(name, result interface{}) {
	fmt.Printf("--- BENCHMARK FINISHED ---\n")
	fmt.Printf("Function:    %s\n", name)
	fmt.Printf("Last Result: %s\n", result)
	fmt.Printf("--------------------------\n")
}

func BenchmarkMarkovSentence(b *testing.B) {
	var result interface{}
	for i := 0; i < b.N; i++ {
		result = markovSentence(24)
	}
	printTestResults(b.Name(), result)
}

func BenchmarkMarkovParagraphs(b *testing.B) {
	var result interface{}
	for i := 0; i < b.N; i++ {
		result = markovParagraphs(10, 5, 20, 25, 50)
	}
	printTestResults(b.Name(), result)
}
func BenchmarkRandomWord(b *testing.B) {
	var result interface{}
	for i := 0; i < b.N; i++ {
		result = randomWord()
	}
	printTestResults(b.Name(), result)
}
func BenchmarkRandomSentence(b *testing.B) {
	var result interface{}
	for i := 0; i < b.N; i++ {
		result = randomSentence(15)
	}
	printTestResults(b.Name(), result)
}
func BenchmarkRandomParagraphs(b *testing.B) {
	var result interface{}
	for i := 0; i < b.N; i++ {
		result = randomParagraphs(20, 5, 10, 20, 40)
	}
	printTestResults(b.Name(), result)
}
func BenchmarkRandomString(b *testing.B) {
	var result interface{}
	for i := 0; i < b.N; i++ {
		result = randomString("username", 19)
	}
	printTestResults(b.Name(), result)
}
func BenchmarkRandomString2(b *testing.B) {
	var result interface{}
	for i := 0; i < b.N; i++ {
		result = randomString("email", 21)
	}
	printTestResults(b.Name(), result)
}
func BenchmarkRandomString3(b *testing.B) {
	var result interface{}
	for i := 0; i < b.N; i++ {
		result = randomString("uuid", 13)
	}
	printTestResults(b.Name(), result)
}
func BenchmarkRandomString4(b *testing.B) {
	var result interface{}
	for i := 0; i < b.N; i++ {
		result = randomString("hex", 8)
	}
	printTestResults(b.Name(), result)
}
func BenchmarkRandomString5(b *testing.B) {
	var result interface{}
	for i := 0; i < b.N; i++ {
		result = randomString("alphanum", 24)
	}
	printTestResults(b.Name(), result)
}
