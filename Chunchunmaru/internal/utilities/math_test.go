package utilities

import (
	"testing"
)

func BenchmarkRandomNumber(b *testing.B) {
	printTestResults("Random Number", RandomNumber(0, 100))
}
