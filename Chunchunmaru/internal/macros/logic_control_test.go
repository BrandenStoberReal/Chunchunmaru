package macros

import (
	"testing"
)

func BenchmarkRandomInt(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomInt(0, 100)
	}

	printTestResults(b.Name(), result)
}

func BenchmarkRepeat(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = repeat(10)
	}

	printTestResults(b.Name(), result)
}

func BenchmarkList(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = list([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}

	printTestResults(b.Name(), result)
}

func BenchmarkRandomChoice(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomChoice([]interface{}{"penis", 1, 5, 71.0, "wow"})
	}

	printTestResults(b.Name(), result)
}
