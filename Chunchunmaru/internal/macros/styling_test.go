package macros

import "testing"

func BenchmarkRandomColor(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomColor()
	}

	printTestResults(b.Name(), result)
}

func BenchmarkRandomId(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomId("test", 10)
	}

	printTestResults(b.Name(), result)
}

func BenchmarkRandomClasses(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomClasses("test", 5)
	}

	printTestResults(b.Name(), result)
}

func BenchmarkRandomCSSStyle(b *testing.B) {
	var result interface{}

	for i := 0; i < b.N; i++ {
		result = randomCSSStyle(10)
	}

	printTestResults(b.Name(), result)
}
