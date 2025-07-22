package macros

import (
	"math/rand"
	"reflect"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func repeat(count int) []int {
	if count < 0 {
		return []int{}
	}
	var l = make([]int, count)
	for i := range count {
		l[i] = i
	}
	return l
}

func list(args ...interface{}) []interface{} {
	return args
}

func randomChoice(slice interface{}) interface{} {
	if slice == nil {
		return nil
	}

	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice || val.Len() == 0 {
		return nil
	}

	return val.Index(rand.Intn(val.Len())).Interface()
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func div(a, b int) int {
	if b == 0 {
		return 0 // or panic, depending on your needs
	}
	return a / b
}

func mult(a, b int) int {
	return a * b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
