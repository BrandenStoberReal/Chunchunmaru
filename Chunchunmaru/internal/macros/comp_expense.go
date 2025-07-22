package macros

import (
	"chunchunmaru/internal/utilities"
	"encoding/base64"
	"fmt"
	"html/template"
	"math"
	"math/rand"
	"strings"
)

// FilterGenerator defines a function signature for generators of SVG filter primitives.
type FilterGenerator func() string

const (
	svgWidth  = 500
	svgHeight = 500
)

// List of unary and binary math operations as templates
var unaryOps = []string{
	"Math.sin(%s)", "Math.cos(%s)", "Math.tan(%s)", "Math.sqrt(%s)", "Math.log(%s)",
	"Math.abs(%s)", "Math.exp(%s)", "Math.sinh(%s)", "Math.cosh(%s)", "Math.tanh(%s)",
	"Math.asin(%s)", "Math.acos(%s)", "Math.atan(%s)", "Math.floor(%s)", "Math.ceil(%s)",
	"Math.round(%s)", "Math.trunc(%s)", "Math.fround(%s)", "Math.sign(%s)", "Math.cbrt(%s)",
}
var binaryOps = []string{
	"(%s + %s)", "(%s - %s)", "(%s * %s)", "(%s / (%s+1))", "(%s % (%s+1))",
	"Math.pow(%s, %s)", "Math.max(%s, %s)", "Math.min(%s, %s)", "Math.atan2(%s, %s)",
	"Math.imul(%s, %s)", "Math.hypot(%s, %s)",
}

// Possible variables to use as leaves
var vars = []string{"i", "j", "waste", "i+j", "i-j", "i*j", "i/j", "j+1", "i%10", "j%10"}

// randomJSExpr Recursively build a random JS math expression
func randomJSExpr(depth int) string {
	if depth <= 0 || rand.Float64() < 0.3 {
		v := vars[rand.Intn(len(vars))]
		if strings.TrimSpace(v) == "" {
			return "1"
		}
		return v
	}
	if rand.Float64() < 0.5 {
		op := unaryOps[rand.Intn(len(unaryOps))]
		inner := randomJSExpr(depth - 1)
		if strings.TrimSpace(inner) == "" {
			inner = "1"
		}
		return strings.Replace(op, "%s", inner, 1)
	} else {
		op := binaryOps[rand.Intn(len(binaryOps))]
		left := randomJSExpr(depth - 1)
		right := randomJSExpr(depth - 1)
		if strings.TrimSpace(left) == "" {
			left = "1"
		}
		if strings.TrimSpace(right) == "" {
			right = "1"
		}
		// Defensive: if op contains %s more than twice, replace all
		out := op
		out = strings.Replace(out, "%s", left, 1)
		out = strings.Replace(out, "%s", right, 1)
		// Final check: if any remaining %s, replace with "1"
		out = strings.ReplaceAll(out, "%s", "1")
		return out
	}
}

// nestDivs
func nestDivs(count int) template.HTML {
	if count <= 0 {
		return ""
	}
	var classPool []string
	for i := 0; i < 100; i++ {
		classPool = append(classPool, randomWord()+fmt.Sprint(rand.Intn(10000)))
	}
	var build func(depth int) string
	build = func(depth int) string {
		if depth == 0 {
			return "content"
		}
		// Random class list
		classCount := rand.Intn(8) + 3
		classes := make([]string, classCount)
		for i := range classes {
			classes[i] = classPool[rand.Intn(len(classPool))]
		}
		// Random inline styles
		styleCount := rand.Intn(6) + 3
		var style strings.Builder
		for i := 0; i < styleCount; i++ {
			style.WriteString(utilities.GenerateRandomInlineCSS())
		}
		// Random attributes
		attrCount := rand.Intn(4) + 1
		var attrs strings.Builder
		for i := 0; i < attrCount; i++ {
			switch rand.Intn(3) {
			case 0:
				attrs.WriteString(fmt.Sprintf(" data-%s=\"%s\"", randomWord(), randomWord()))
			case 1:
				attrs.WriteString(fmt.Sprintf(" aria-%s=\"%s\"", randomWord(), randomWord()))
			case 2:
				attrs.WriteString(fmt.Sprintf(" id=\"%s\"", utilities.RandomStringFromCharset(8, utilities.LowerAlphabetChars)))
			}
		}
		// Randomly insert other elements
		var inner string
		if rand.Float64() < 0.3 && depth > 1 {
			inner = fmt.Sprintf("<span>%s</span>%s<b>%s</b>", randomWord(), build(depth-1), randomWord())
		} else {
			inner = build(depth - 1)
		}
		return fmt.Sprintf("<div class=\"%s\" style=\"%s\"%s>%s</div>", strings.Join(classes, " "), style.String(), attrs.String(), inner)
	}
	return template.HTML(build(count))
}

// randomComplexTable
func randomComplexTable(rows, cols int) template.HTML {
	if rows < 1 || cols < 1 {
		return ""
	}
	var builder strings.Builder
	builder.WriteString("<table")
	if rand.Float64() < 0.5 {
		builder.WriteString(fmt.Sprintf(" class=\"%s\"", randomWord()))
	}
	builder.WriteString(">\n")
	sections := []string{"thead", "tbody", "tfoot"}
	sectionRows := []int{1, rows - 2, 1}
	rowIdx := 0
	for s, section := range sections {
		if sectionRows[s] <= 0 {
			continue
		}
		builder.WriteString(fmt.Sprintf("<%s>\n", section))
		for r := 0; r < sectionRows[s]; r++ {
			builder.WriteString("<tr>\n")
			c := 0
			for c < cols {
				if rand.Float64() < 0.2 && c < cols-1 {
					// Nested table
					builder.WriteString("<td>")
					builder.WriteString(string(randomComplexTable(1, rand.Intn(2)+1)))
					builder.WriteString("</td>")
					c++
					continue
				}
				colspan := rand.Intn(5) + 1
				rowspan := rand.Intn(5) + 1
				if c+colspan > cols {
					colspan = cols - c
				}
				tag := "td"
				if rand.Float64() < 0.2 {
					tag = "th"
				}
				style := utilities.GenerateRandomInlineCSS()
				attr := ""
				if rand.Float64() < 0.5 {
					attr = fmt.Sprintf(" data-%s=\"%s\"", randomWord(), randomWord())
				}
				content := randomWord() + fmt.Sprintf(" (%d,%d)", rowIdx, c)
				builder.WriteString(fmt.Sprintf("<%s colspan=\"%d\" rowspan=\"%d\" style=\"%s\"%s>%s</%s>\n", tag, colspan, rowspan, style, attr, content, tag))
				c += colspan
			}
			builder.WriteString("</tr>\n")
			rowIdx++
		}
		builder.WriteString(fmt.Sprintf("</%s>\n", section))
	}
	builder.WriteString("</table>")
	return template.HTML(builder.String())
}

type StyleBlock struct {
	Style template.HTML
	Class template.CSS
}

// randomStyleBlock
func randomStyleBlock(styleType string, count int) StyleBlock {
	parentClass := "p" + utilities.RandomStringFromCharset(12, utilities.LowerAlphabetChars)
	var builder strings.Builder
	builder.WriteString("<style>\n")
	switch styleType {
	case "utility":
		for i := 0; i < count; i++ {
			class := "u" + utilities.RandomStringFromCharset(10, utilities.LowerAlphabetChars)
			props := ""
			for j := 0; j < rand.Intn(8)+3; j++ {
				props += utilities.GenerateRandomInlineCSS()
			}
			builder.WriteString(fmt.Sprintf(".%s .%s { %s }\n", parentClass, class, props))
		}
	case "nested":
		for i := 0; i < count; i++ {
			selector := "." + parentClass
			depth := rand.Intn(8) + 3
			for d := 0; d < depth; d++ {
				selector += " > div"
				if rand.Float64() < 0.5 {
					selector += fmt.Sprintf(":nth-child(%d)", rand.Intn(10)+1)
				}
				if rand.Float64() < 0.3 {
					selector += fmt.Sprintf(" + span")
				}
			}
			props := ""
			for j := 0; j < rand.Intn(8)+3; j++ {
				props += utilities.GenerateRandomInlineCSS()
			}
			builder.WriteString(fmt.Sprintf("%s { %s }\n", selector, props))
		}
	case "complex":
		for i := 0; i < count; i++ {
			attr := fmt.Sprintf("[data-%s=\"%s\"]", randomWord(), randomWord())
			pseudo := ":" + utilities.RandomKeyword([]string{
				"hover", "active", "focus", "nth-child(2)", "not(:last-child)", "first-of-type", "last-of-type", "nth-of-type(odd)", "nth-of-type(even)",
			})
			pseudoElem := ""
			if rand.Float64() < 0.5 {
				pseudoElem = "::" + utilities.RandomKeyword([]string{"before", "after", "marker", "selection"})
			}
			combinator := ""
			if rand.Float64() < 0.5 {
				combinator = " > " + utilities.RandomKeyword([]string{"span", "b", "i", "u"})
			}
			props := ""
			for j := 0; j < rand.Intn(8)+3; j++ {
				props += utilities.GenerateRandomInlineCSS()
			}
			builder.WriteString(fmt.Sprintf(".%s %s%s%s%s { %s }\n", parentClass, attr, pseudo, pseudoElem, combinator, props))
		}
	}
	builder.WriteString("</style>\n")
	return StyleBlock{
		Style: template.HTML(builder.String()),
		Class: template.CSS(parentClass),
	}
}

// generateFractalSVG creates an SVG of a fractal tree. The parameters of the tree
// (depth, angles, length, branches) are randomized to make each SVG unique and
// computationally expensive in a different way for the client's renderer.
func generateFractalSVG() string {
	var pathBuilder strings.Builder
	initialLength := float64(rand.Intn(110-70) + 70)
	initialAngle := -90.0                // Start pointing up
	recursionDepth := rand.Intn(5-4) + 4 // Depth between 4 and 5

	// Begin the recursive drawing process from the bottom-center of the SVG canvas.
	drawFractalBranch(&pathBuilder, svgWidth/2, svgHeight, initialAngle, initialLength, recursionDepth)

	// Assemble the final SVG string.
	svg := fmt.Sprintf(
		`<svg width="%d" height="%d" viewBox="0 0 %d %d" xmlns="http://www.w3.org/2000/svg">`+
			`<path d="%s" stroke="%s" stroke-width="2" fill="none" />`+
			`</svg>`,
		svgWidth, svgHeight, svgWidth, svgHeight,
		pathBuilder.String(),
		randomColor(), // Assumes randomColor() is an available macro function
	)
	return svg
}

// drawFractalBranch is a recursive helper function that builds the fractal path data.
func drawFractalBranch(path *strings.Builder, x1, y1, angle, length float64, depth int) {
	if depth <= 0 {
		return
	}

	// Calculate the end point (x2, y2) of the current branch.
	rad := angle * (math.Pi / 180.0)
	x2 := x1 + length*math.Cos(rad)
	y2 := y1 + length*math.Sin(rad)

	// Append the SVG path command for the new line segment.
	path.WriteString(fmt.Sprintf("M%.2f,%.2f L%.2f,%.2f ", x1, y1, x2, y2))

	// For the next recursion level, randomize the number of new branches and their properties.
	numBranches := rand.Intn(3-2) + 2 // 2 or 3 branches
	for i := 0; i < numBranches; i++ {
		// Each new branch has a randomly perturbed angle and reduced length.
		newAngle := angle + (-40 + rand.Float64()*(40-(-40)))     // Angle change between -40 and +40
		newLength := length * (0.70 + rand.Float64()*(0.85-0.70)) // Length scales between 70% and 85%
		drawFractalBranch(path, x2, y2, newAngle, newLength, depth-1)
	}
}

// A list of available filter generator functions. This list is shuffled on each call
// to randomize the filter chain's composition and order.
var filterGenerators = []FilterGenerator{
	generateFeGaussianBlur,
	generateFeMorphology,
	generateFeTurbulence,
	generateFeColorMatrix,
	generateFeConvolveMatrix,
	generateFeDisplacementMap,
}

// generateFiltersSVG creates an SVG with a randomized chain of complex filters applied
// to a simple shape. This forces a heavy computational load on the client's rendering engine.
func generateFiltersSVG() string {
	// Updated to match spec: randomId "prefix" count
	filterID := randomId("filter", 8) // Assumes randomId() is available
	numFilters := rand.Intn(5-3) + 3  // Chain of 3 to 5 filters
	var filterChain strings.Builder

	// Create a randomized pipeline of filters.
	rand.Shuffle(len(filterGenerators), func(i, j int) {
		filterGenerators[i], filterGenerators[j] = filterGenerators[j], filterGenerators[i]
	})

	// Build the filter chain from the shuffled list.
	for i := 0; i < numFilters && i < len(filterGenerators); i++ {
		filterChain.WriteString(filterGenerators[i]())
	}

	// Define a simple shape to apply the complex filter to.
	rectWidth := rand.Intn(300-200) + 200
	rectHeight := rand.Intn(250-150) + 150
	rectX := (svgWidth - rectWidth) / 2
	rectY := (svgHeight - rectHeight) / 2

	// Assemble the final SVG, defining the filter chain in <defs> and applying it.
	svg := fmt.Sprintf(
		`<svg width="%d" height="%d" viewBox="0 0 %d %d" xmlns="http://www.w3.org/2000/svg">`+
			`<defs>`+
			`<filter id="%s" x="-50%%" y="-50%%" width="200%%" height="200%%">`+ // Expand filter area to avoid clipping
			`%s`+
			`</filter>`+
			`</defs>`+
			`<rect x="%d" y="%d" width="%d" height="%d" fill="%s" filter="url(#%s)" />`+
			`</svg>`,
		svgWidth, svgHeight, svgWidth, svgHeight,
		filterID,
		filterChain.String(),
		rectX, rectY, rectWidth, rectHeight,
		randomColor(), // Assumes randomColor() is available
		filterID,
	)
	return svg
}

// generateFeGaussianBlur
func generateFeGaussianBlur() string {
	deviation := rand.Intn(6-2) + 2
	return fmt.Sprintf(`<feGaussianBlur stdDeviation="%d" />`, deviation)
}

// generateFeMorphology
func generateFeMorphology() string {
	operators := []string{"erode", "dilate"}
	operator := operators[rand.Intn(len(operators))]
	radius := rand.Intn(5-2) + 2
	return fmt.Sprintf(`<feMorphology operator="%s" radius="%d" />`, operator, radius)
}

// generateFeTurbulence
func generateFeTurbulence() string {
	baseFreq := 0.01 + rand.Float64()*(0.08-0.01)
	numOctaves := rand.Intn(4-2) + 2
	turbulenceTypes := []string{"fractalNoise", "turbulence"}
	turbulenceType := turbulenceTypes[rand.Intn(len(turbulenceTypes))]
	return fmt.Sprintf(`<feTurbulence type="%s" baseFrequency="%.4f" numOctaves="%d" result="noise" />`, turbulenceType, baseFreq, numOctaves)
}

// generateFeDisplacementMap
func generateFeDisplacementMap() string {
	scale := rand.Intn(75-20) + 20
	return fmt.Sprintf(`<feDisplacementMap in="SourceGraphic" in2="noise" scale="%d" />`, scale)
}

// generateFeColorMatrix
func generateFeColorMatrix() string {
	types := []string{"saturate", "hueRotate"}
	matrixType := types[rand.Intn(len(types))]
	var value float64
	if matrixType == "hueRotate" {
		value = float64(rand.Intn(360))
	} else { // saturate
		value = rand.Float64() // a value between 0.0 and 1.0
	}
	return fmt.Sprintf(`<feColorMatrix type="%s" values="%.2f" />`, matrixType, value)
}

// generateFeConvolveMatrix
func generateFeConvolveMatrix() string {
	var kernelMatrix strings.Builder
	for i := 0; i < 9; i++ {
		// Value between -1.5 and 1.5
		v := -1.5 + rand.Float64()*(1.5-(-1.5))
		kernelMatrix.WriteString(fmt.Sprintf("%.2f ", v))
	}
	order := 3 // 3x3 matrix
	bias := rand.Float64() * 0.3
	return fmt.Sprintf(`<feConvolveMatrix order="%d" kernelMatrix="%s" bias="%.2f" />`, order, strings.TrimSpace(kernelMatrix.String()), bias)
}

// randomSVG is the function registered in the template FuncMap. It acts as a
// dispatcher, calling the appropriate generator based on the provided type.
// It returns template.HTML to prevent auto-escaping of the SVG markup.
func randomSVG(svgType string) template.HTML {
	switch svgType {
	case "fractal":
		return template.HTML(generateFractalSVG())
	case "filters":
		return template.HTML(generateFiltersSVG())
	default:
		// Return an error for an unknown type, allowing for robust error handling.
		return ""
	}
}

func randomCSSVars(count int) template.HTML {
	if count < 3 {
		count = 3
	}
	var builder strings.Builder
	builder.WriteString("<style>\n:root {\n")
	builder.WriteString(fmt.Sprintf("  --v1: %dpx;\n", rand.Intn(100)+1))
	builder.WriteString(fmt.Sprintf("  --v2: %dpx;\n", rand.Intn(100)+1))
	for i := 3; i <= count; i++ {
		refs := []string{}
		refCount := rand.Intn(3) + 1
		for j := 0; j < refCount; j++ {
			refs = append(refs, fmt.Sprintf("var(--v%d)", rand.Intn(i-1)+1))
		}
		op := utilities.RandomKeyword([]string{"+", "-", "*", "/", "min", "max"})
		expr := ""
		if op == "min" || op == "max" {
			expr = fmt.Sprintf("%s(%s, %dpx)", op, strings.Join(refs, ", "), rand.Intn(100)+1)
		} else {
			expr = refs[0]
			for _, r := range refs[1:] {
				expr += fmt.Sprintf(" %s %s", op, r)
			}
			expr += fmt.Sprintf(" %s %dpx", op, rand.Intn(100)+1)
		}
		builder.WriteString(fmt.Sprintf("  --v%d: calc(%s);\n", i, expr))
	}
	// Randomly insert a cycle (for chaos)
	if count > 5 && rand.Float64() < 0.3 {
		builder.WriteString(fmt.Sprintf("  --v1: calc(var(--v%d) + 1px);\n", count))
	}
	builder.WriteString("}\n</style>")
	return template.HTML(builder.String())
}

func jsInteractiveContent(typ, content string) template.HTML {
	placeholderID := "ph" + utilities.RandomStringFromCharset(12, utilities.LowerAlphabetChars)

	// Generate a random number of random expressions for the waste loop
	numExpr := rand.Intn(4) + 3 // 3–6 expressions
	var exprs []string
	for i := 0; i < numExpr; i++ {
		exprs = append(exprs, "waste+="+randomJSExpr(3)+";")
	}
	concatMath := strings.Join(exprs, "\n    ")

	// Waste CPU: use the composed math expressions in a nested loop
	cpuWasteJS := fmt.Sprintf(`
let waste=0;
for(let i=0;i<1e5;i++){
  for(let j=0;j<100;j++){
    %s
  }
}
`, concatMath)

	// Obfuscation strategies
	key := byte(rand.Intn(256))
	type obfStrategy struct {
		name   string
		encode func(string) (string, string)
	}
	strategies := []obfStrategy{
		{
			"base64",
			func(s string) (string, string) {
				enc := base64.StdEncoding.EncodeToString([]byte(s))
				return enc, "let data=atob(obf);"
			},
		},
		{
			"reversed base64",
			func(s string) (string, string) {
				enc := utilities.ReverseString(base64.StdEncoding.EncodeToString([]byte(s)))
				return enc, "let data=atob(obf.split('').reverse().join(''));"
			},
		},
		{
			"hex",
			func(s string) (string, string) {
				enc := utilities.HexEncode(s)
				return enc, `
let hex=obf;
let data='';
for(let i=0;i<hex.length;i+=2){
  data+=String.fromCharCode(parseInt(hex.substr(i,2),16));
}`
			},
		},
		{
			"double base64",
			func(s string) (string, string) {
				enc := base64.StdEncoding.EncodeToString([]byte(base64.StdEncoding.EncodeToString([]byte(s))))
				return enc, "let data=atob(atob(obf));"
			},
		},
		{
			"xor+base64",
			func(s string) (string, string) {
				enc := utilities.XorEncode(s, key)
				return enc, fmt.Sprintf(`
let b=atob(obf);
let data='';
for(let i=0;i<b.length;i++){data+=String.fromCharCode(b.charCodeAt(i)^%d);}
`, key)
			},
		},
	}

	// Randomly pick a strategy
	strategy := strategies[rand.Intn(len(strategies))]
	encoded, decodeJS := strategy.encode(content)

	// Compose the script
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("<%s id=\"%s\"></%s>\n", typ, placeholderID, typ))
	builder.WriteString("<script>(function(){\n")
	builder.WriteString(cpuWasteJS + "\n")
	builder.WriteString(fmt.Sprintf("let obf='%s';\n", encoded))
	builder.WriteString(decodeJS + "\n")
	builder.WriteString(fmt.Sprintf("document.getElementById('%s').innerHTML=data;\n", placeholderID))
	builder.WriteString("})();</script>")
	return template.HTML(builder.String())
}
