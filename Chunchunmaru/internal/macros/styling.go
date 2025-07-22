package macros

import (
	"chunchunmaru/internal/utilities"
	"html/template"
	"math/rand"
	"strings"
)

var classNames = [11]string{"mx-auto", "flex", "max-w-sm", "items-center", "gap-x-4", "rounded-xl", "bg-white", "p-6", "shadow-lg", "outline", "outline-black"}

// randomColor Fetches a random hex color
func randomColor() string {
	return utilities.RandomHexColor()
}

// randomId Fetches a random HTML element ID
func randomId(prefix string, length int) string {
	return utilities.RandomStringFromCharset(length, prefix+"-"+utilities.MixedDigitChars)
}

// randomClasses Fetches random CSS classes and returns them in a string
func randomClasses(count int) template.CSS {
	var builder strings.Builder
	for i := 0; i < count; i++ {
		if i == count-1 {
			builder.WriteString(classNames[rand.Intn(len(classNames))])
		} else {
			builder.WriteString(classNames[rand.Intn(len(classNames))])
			builder.WriteString(" ")
		}
	}
	return template.CSS(builder.String())
}

// randomInlineStyle Provides a full HTML style tag with random style data
func randomInlineStyle(length int) template.CSS {
	var builder strings.Builder
	builder.WriteString("style=\"")
	for i := 0; i < length; i++ {
		builder.WriteString(utilities.GenerateRandomInlineCSS())
	}
	builder.WriteString("\"")

	return template.CSS(builder.String())
}

// randomCSSStyle
func randomCSSStyle(length int) template.CSS {
	if length == 1 {
		return template.CSS(utilities.GenerateRandomInlineCSS())
	}
	var builder strings.Builder

	for i := 0; i < length; i++ {
		builder.WriteString(utilities.GenerateRandomInlineCSS())
	}

	return template.CSS(builder.String())
}
