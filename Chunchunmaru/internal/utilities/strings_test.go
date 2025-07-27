package utilities

import (
	"testing"
)

func BenchmarkRandomStringFromCharset(b *testing.B) {
	printTestResults("Random String From Charset", RandomStringFromCharset(12, AlphabetChars))
}

func BenchmarkRandomKeyword(b *testing.B) {
	items := []string{"foo", "bar", "baz"}
	printTestResults("Random String From Array", RandomKeyword(items))
}

func BenchmarkRandomHexColor(b *testing.B) {
	printTestResults("Random Hex Color (string utility)", RandomHexColor())
}

func BenchmarkReverseString(b *testing.B) {
	printTestResults("Reverse String", ReverseString("foobarbaz!"))
}

func BenchmarkHexEncode(b *testing.B) {
	printTestResults("Hex Encode String", HexEncode("foobarbaz!"))
}

func BenchmarkXorEncode(b *testing.B) {
	printTestResults("Hex Encode String", XorEncode("foobarbaz!", 3))
}
