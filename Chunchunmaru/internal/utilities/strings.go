package utilities

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func CleanString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const UpperHexChars = "0123456789ABCDEF"
const LowerHexChars = "0123456789abcdef"
const UpperAlphabetChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LowerAlphabetChars = "abcdefghijklmnopqrstuvwxyz"
const AlphabetChars = UpperAlphabetChars + LowerAlphabetChars
const AlphabetNumericChars = "0123456789"
const SpeciaLChars = "!@#$%^&*()_+-=<>?"
const MixedDigitChars = AlphabetChars + AlphabetNumericChars
const AllChars = MixedDigitChars + SpeciaLChars

func RandomStringFromCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// randomKeyword selects a random item from a slice of strings.
func RandomKeyword(keywords []string) string {
	return keywords[rand.Intn(len(keywords))]
}

// randomHexColor generates a random hexadecimal color.
func RandomHexColor() string {
	return fmt.Sprintf("#%06x", rand.Intn(0xFFFFFF))
}

// Helper: reverse a string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Helper: hex encode a string
func HexEncode(s string) string {
	return fmt.Sprintf("%x", s)
}

// Helper: xor encode a string with a key
func XorEncode(s string, key byte) string {
	b := []byte(s)
	for i := range b {
		b[i] ^= key
	}
	return base64.StdEncoding.EncodeToString(b)
}
