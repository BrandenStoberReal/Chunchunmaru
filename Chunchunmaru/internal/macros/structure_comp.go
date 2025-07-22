package macros

import (
	"html/template"
	"math/rand"
	"strings"
)

var types = [5]string{"text", "radio", "checkbox", "submit", "button"}

func randomForm(count, styleCount int) template.HTML {
	var builder strings.Builder
	builder.WriteString("<form>\n")

	for i := 0; i < count; i++ {
		id := randomId("input", 5)

		// Input
		builder.WriteString("<input type=\"")
		builder.WriteString(types[rand.Intn(len(types))]) // Input type
		builder.WriteString("\" id=\"")
		builder.WriteString(id) // ID of element
		builder.WriteString("\">\n")

		// Label
		builder.WriteString("<label for=\"")
		builder.WriteString(id) // ID of element label will use
		builder.WriteString("\" ")
		builder.WriteString(string(randomInlineStyle(rand.Intn(styleCount)))) // Random CSS garbage
		builder.WriteString(">")
		builder.WriteString(randomWord()) // Random label word
		builder.WriteString("</label>\n")
	}
	builder.WriteString("</form>")
	return template.HTML(builder.String())
}

type definition struct {
	Term string
	Def  string
}

func randomDefinitionData(count, len int) []definition {
	arr := make([]definition, count)
	for i := 0; i < count; i++ {
		arr[i] = definition{
			Term: randomWord(),
			Def:  randomSentence(len),
		}
	}
	return arr
}
