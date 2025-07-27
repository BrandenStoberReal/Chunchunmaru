package macros

import (
	"chunchunmaru/internal/utilities"
	"github.com/mb-14/gomarkov"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func markovSentence(length int) string {
	model, err := utilities.LoadMarkovModel()
	if err != nil {
		return ""
	}
	order := model.Order
	tokens := make([]string, 0)
	for i := 0; i < order; i++ {
		tokens = append(tokens, gomarkov.StartToken)
	}
	for tokens[len(tokens)-1] != gomarkov.EndToken && len(tokens) < length {
		next, _ := model.Generate(tokens[(len(tokens) - order):])
		tokens = append(tokens, next)
	}
	return strings.Join(tokens[order:len(tokens)-1], " ") + "."
}

func markovParagraphs(count, minSentences, maxSentences, minSentenceLength, maxSentenceLength int) string {
	output := ""

	for i := 0; i < count; i++ {
		for j := 0; j < randomInt(minSentences, maxSentences); j++ {
			output = output + markovSentence(randomInt(minSentenceLength, maxSentenceLength)) + " "
		}
		output = output + "\n"
	}
	return output
}

func randomWord() string {
	return utilities.CleanString(utilities.RandomWord())
}

func randomSentence(len int) string {
	if len <= 0 {
		return ""
	}

	var builder strings.Builder

	builder.Grow(len * utilities.AvgWordLen)

	for i := 0; i < len; i++ {
		word := utilities.RandomWord()
		if i == 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			builder.WriteString(string(runes))
		} else {
			builder.WriteString(word)
		}
		if i < len-1 {
			builder.WriteByte(' ')
		}
	}
	builder.WriteByte('.')
	return builder.String()
}

func randomParagraphs(count, minSentences, maxSentences, minSentenceLength, maxSentenceLength int) string {
	var builder strings.Builder

	builder.Grow(count * utilities.AvgWordLen * (minSentenceLength + maxSentenceLength) * (minSentences + maxSentences) / 4)

	for i := 0; i < count; i++ {
		sentences := rand.Intn(maxSentences-minSentences) + minSentences
		for j := 0; j < sentences; j++ {
			builder.WriteString(randomSentence(rand.Intn(maxSentenceLength-minSentenceLength) + minSentenceLength))
			if j < sentences-1 {
				builder.WriteByte(' ')
			}
		}

		if i < count-1 {
			builder.WriteString("\n\n")
		}
	}

	return builder.String()
}

func randomString(t string, len int) string {
	switch t {
	case "username":
		return utilities.RandomWord() + utilities.RandomWord()
	case "email":
		return utilities.RandomWord() + utilities.RandomWord() + "@gmail.com"
	case "uuid":
		return utilities.RandomStringFromCharset(8, utilities.LowerHexChars) + "-" + utilities.RandomStringFromCharset(4, utilities.LowerHexChars) + "-" + utilities.RandomStringFromCharset(4, utilities.LowerHexChars) + "-" + utilities.RandomStringFromCharset(4, utilities.LowerHexChars) + "-" + utilities.RandomStringFromCharset(12, utilities.LowerHexChars)
	case "hex":
		return utilities.RandomStringFromCharset(len, utilities.UpperHexChars)
	case "alphanum":
		return utilities.RandomStringFromCharset(len, utilities.MixedDigitChars)
	default:
		return ""
	}
}

func randomDate(format, start, end string) string {
	startTime, err := time.Parse(format, start)
	if err != nil {
		return ""
	}
	endTime, err := time.Parse(format, end)
	if err != nil {
		return ""
	}

	if startTime.After(endTime) {
		return ""
	}

	startUnix := startTime.Unix()
	endUnix := endTime.Unix()

	delta := endUnix - startUnix

	if delta <= 0 {
		return startTime.Format(format)
	}

	randomTime := time.Unix(rand.Int63n(delta)+startUnix, 0)

	return randomTime.Format(format)
}
