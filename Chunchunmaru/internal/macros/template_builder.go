package macros

import (
	"html/template"
)

var funcMap = template.FuncMap{
	// Category 1: Content Generation
	"markovSentence":   markovSentence,
	"markovParagraphs": markovParagraphs,
	"randomParagraphs": randomParagraphs,
	"randomSentence":   randomSentence,
	"randomWord":       randomWord,
	"randomString":     randomString,
	"randomDate":       randomDate,

	// Category 2: Structure & Composition
	"randomForm":           randomForm,
	"randomDefinitionData": randomDefinitionData,

	// Category 3: Styling
	"randomColor":       randomColor,
	"randomId":          randomId,
	"randomClasses":     randomClasses,
	"randomInlineStyle": randomInlineStyle,
	"randomCSSStyle":    randomCSSStyle,

	// Category 4: Link & Navigation
	"randomLink":      randomLink,
	"randomQueryLink": randomQueryLink,
	"randomJSON":      randomJSON,

	// Category 5: Logic & Control
	"randomInt":    randomInt,
	"repeat":       repeat,
	"list":         list,
	"randomChoice": randomChoice,
	"add":          add,
	"sub":          sub,
	"div":          div,
	"mult":         mult,
	"max":          max,
	"min":          min,

	// Category 6: Computationally Expensive & Anti-Extraction
	"nestDivs":             nestDivs,
	"randomComplexTable":   randomComplexTable,
	"randomStyleBlock":     randomStyleBlock,
	"randomSVG":            randomSVG,
	"randomCSSVars":        randomCSSVars,
	"jsInteractiveContent": jsInteractiveContent,
}

type TemplateInput struct {
	Aggression int
}

func BuildTemplate(name, content string) (*template.Template, error) {
	tmp, err := template.New(name).Funcs(funcMap).Parse(content)
	return tmp, err
}
