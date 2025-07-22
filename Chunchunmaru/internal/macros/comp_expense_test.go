package macros

import (
	"testing"
)

func BenchmarkNestDivs(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = nestDivs(10)
	}

	printTestResults(b.Name(), result)
}
func BenchmarkRandomComplexTable(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomComplexTable(10, 5)
	}

	printTestResults(b.Name(), result)
}
func BenchmarkRandomStyleBlock(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomStyleBlock("utility", 8)
	}

	printTestResults(b.Name(), result)
}
func BenchmarkRandomStyleBlock2(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomStyleBlock("nested", 6)
	}

	printTestResults(b.Name(), result)
}
func BenchmarkRandomStyleBlock3(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomStyleBlock("complex", 7)
	}

	printTestResults(b.Name(), result)
}
func BenchmarkRandomSVG(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomSVG("fractal")
	}

	printTestResults(b.Name(), result)
}
func BenchmarkRandomSVG2(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomSVG("filters")
	}

	printTestResults(b.Name(), result)
}
func BenchmarkRandomCSSVars(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomCSSVars(11)
	}

	printTestResults(b.Name(), result)
}
func BenchmarkJsInteractiveContent(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = jsInteractiveContent("div", "Hello World!")
	}

	printTestResults(b.Name(), result)
}
