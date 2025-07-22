package macros

import (
	"testing"
)

func BenchmarkRandomLink(b *testing.B) {
	var result interface{}
	for i := 0; i < b.N; i++ {
		result = randomLink()
	}
	printTestResults(b.Name(), result)
}

func BenchmarkRandomQueryLink(b *testing.B) {
	var result interface{}
	for i := 0; i < b.N; i++ {
		result = randomQueryLink(10)
	}
	printTestResults(b.Name(), result)
}

func BenchmarkRandomJSON(b *testing.B) {
	var result interface{}
	for i := 0; i < b.N; i++ {
		result, _ = randomJSON(5, 5, 20)
	}
	printTestResults(b.Name(), result)
}
