package utilities

import (
	"testing"
)

func TestRandomNumber(t *testing.T) {
	printTestResults("Random Number", RandomNumber(0, 100))
}
