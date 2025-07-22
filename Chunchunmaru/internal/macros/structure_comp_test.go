package macros

import (
	"testing"
)

func BenchmarkRandomForm(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomForm(10, 12)
	}

	printTestResults(b.Name(), result)
}

func BenchmarkRandomDefinitionData(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomDefinitionData(10, 12)
	}

	printTestResults(b.Name(), result)
}
