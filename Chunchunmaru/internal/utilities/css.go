package utilities

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// randomLength generates a random length value with a unit.
func randomLength(min, max int, units ...string) string {
	if len(units) == 0 {
		units = []string{"px", "em", "%", "rem", "vh", "vw"}
	}
	value := rand.Intn(max-min+1) + min
	unit := units[rand.Intn(len(units))]
	return strconv.Itoa(value) + unit
}

// randomFloat generates a random float value for properties like opacity or line-height.
func randomFloat(min, max float64) string {
	val := min + rand.Float64()*(max-min)
	return fmt.Sprintf("%.2f", val)
}

// randomShorthandLength generates a random shorthand value for margin or padding.
func randomShorthandLength() string {
	count := rand.Intn(4) + 1 // 1 to 4 values
	values := make([]string, count)
	for i := 0; i < count; i++ {
		values[i] = randomLength(0, 50, "px", "%", "em")
	}
	return strings.Join(values, " ")
}

// randomShadow generates a random box-shadow or text-shadow.
func randomShadow() string {
	hOffset := randomLength(-10, 10, "px")
	vOffset := randomLength(-10, 10, "px")
	blur := randomLength(0, 20, "px")
	spread := ""
	inset := ""

	// box-shadow can have spread and inset
	if rand.Intn(2) == 0 {
		spread = " " + randomLength(0, 15, "px")
		if rand.Intn(3) == 0 {
			inset = " inset"
		}
	}
	color := RandomHexColor()
	return fmt.Sprintf("%s %s %s%s %s%s", hOffset, vOffset, blur, spread, color, inset)
}

// randomGridTemplate generates a value for grid-template-columns/rows.
func randomGridTemplate() string {
	parts := []string{}
	for i := 0; i < rand.Intn(4)+1; i++ { // 1 to 4 tracks
		partType := rand.Intn(3)
		switch partType {
		case 0:
			parts = append(parts, randomLength(20, 200, "px", "em"))
		case 1:
			parts = append(parts, randomLength(10, 100, "%"))
		case 2:
			parts = append(parts, fmt.Sprintf("%dfr", rand.Intn(4)+1))
		}
	}
	return strings.Join(parts, " ")
}

// randomTransition generates a value for the transition shorthand property.
func randomTransition() string {
	property := RandomKeyword([]string{"all", "opacity", "transform", "background-color", "color"})
	duration := fmt.Sprintf("%.2fs", 0.1+rand.Float64()*1.9) // 0.1s to 2.0s
	timingFunction := RandomKeyword([]string{"ease", "ease-in", "ease-out", "ease-in-out", "linear", "step-start"})
	delay := fmt.Sprintf("%.2fs", rand.Float64()*0.5) // 0s to 0.5s
	return fmt.Sprintf("%s %s %s %s", property, duration, timingFunction, delay)
}

// randomClipPath generates a simple clip-path value.
func randomClipPath() string {
	shapes := []string{
		fmt.Sprintf("circle(%s at %s %s)", randomLength(25, 50, "%"), randomLength(25, 75, "%"), randomLength(25, 75, "%")),
		fmt.Sprintf("ellipse(%s %s at 50%% 50%%)", randomLength(25, 50, "%"), randomLength(25, 50, "%")),
		fmt.Sprintf("polygon(%s %s, %s %s, %s %s)", randomLength(0, 100, "%"), randomLength(0, 100, "%"), randomLength(0, 100, "%"), randomLength(0, 100, "%"), randomLength(0, 100, "%"), randomLength(0, 100, "%")),
		"inset(10% 20% 30% 40% round 10px)",
	}
	return RandomKeyword(shapes)
}

var cssProperties = map[string]func() string{
	// --- Text & Font Properties ---
	"color": func() string {
		return RandomKeyword([]string{"red", "blue", "green", "purple", "orange", "pink", "brown", "black", "white", RandomHexColor()})
	},
	"font-size": func() string { return randomLength(12, 48, "px", "em", "rem") },
	"font-family": func() string {
		return RandomKeyword([]string{`"Arial", sans-serif`, `"Georgia", serif`, `"Courier New", monospace`, `"Verdana", sans-serif`, `"Times New Roman", serif`})
	},
	"font-style": func() string { return RandomKeyword([]string{"normal", "italic", "oblique"}) },
	"font-weight": func() string {
		return RandomKeyword([]string{"normal", "bold", "bolder", "lighter", "100", "400", "700", "900"})
	},
	"font-variant": func() string { return RandomKeyword([]string{"normal", "small-caps"}) },
	"text-align":   func() string { return RandomKeyword([]string{"left", "right", "center", "justify"}) },
	"text-decoration": func() string {
		return fmt.Sprintf("%s %s %s", RandomKeyword([]string{"none", "underline", "overline", "line-through"}), RandomKeyword([]string{"solid", "wavy", "dotted"}), RandomHexColor())
	},
	"text-transform": func() string { return RandomKeyword([]string{"none", "capitalize", "uppercase", "lowercase"}) },
	"text-shadow":    randomShadow,
	"text-indent":    func() string { return randomLength(0, 100, "px", "%") },
	"line-height":    func() string { return randomFloat(1, 2.5) },
	"letter-spacing": func() string { return randomLength(-2, 10, "px") },
	"word-spacing":   func() string { return randomLength(0, 15, "px", "em") },
	"white-space":    func() string { return RandomKeyword([]string{"normal", "nowrap", "pre", "pre-wrap"}) },
	"writing-mode":   func() string { return RandomKeyword([]string{"horizontal-tb", "vertical-rl", "vertical-lr"}) },

	// --- Box Model ---
	"width":   func() string { return RandomKeyword([]string{"auto", randomLength(20, 100, "%", "px", "vw")}) },
	"height":  func() string { return RandomKeyword([]string{"auto", randomLength(20, 100, "%", "px", "vh")}) },
	"margin":  randomShorthandLength,
	"padding": randomShorthandLength,
	"border": func() string {
		return fmt.Sprintf("%s %s %s", randomLength(1, 12, "px"), RandomKeyword([]string{"solid", "dotted", "dashed", "double"}), RandomHexColor())
	},
	"border-radius": func() string { return randomLength(0, 50, "%", "px") },
	"box-shadow":    randomShadow,
	"box-sizing":    func() string { return RandomKeyword([]string{"content-box", "border-box"}) },

	// --- Positioning ---
	"position": func() string { return RandomKeyword([]string{"static", "relative", "absolute", "fixed", "sticky"}) },
	"top":      func() string { return randomLength(-50, 100, "px", "%") },
	"right":    func() string { return randomLength(-50, 100, "px", "%") },
	"bottom":   func() string { return randomLength(-50, 100, "px", "%") },
	"left":     func() string { return randomLength(-50, 100, "px", "%") },
	"overflow": func() string { return RandomKeyword([]string{"visible", "hidden", "scroll", "auto"}) },
	"z-index":  func() string { return strconv.Itoa(rand.Intn(1000) - 500) },

	// --- Background ---
	"background-color": RandomHexColor,
	"background-image": func() string {
		return fmt.Sprintf("linear-gradient(%sdeg, %s, %s)", strconv.Itoa(rand.Intn(361)), RandomHexColor(), RandomHexColor())
	},
	"background-size":     func() string { return RandomKeyword([]string{"auto", "cover", "contain"}) },
	"background-repeat":   func() string { return RandomKeyword([]string{"repeat", "no-repeat", "repeat-x", "repeat-y"}) },
	"background-position": func() string { return fmt.Sprintf("%s %s", randomLength(0, 100, "%"), randomLength(0, 100, "%")) },
	"background-blend-mode": func() string {
		return RandomKeyword([]string{"normal", "multiply", "screen", "overlay", "darken", "lighten", "color-dodge"})
	},
	"background-attachment": func() string { return RandomKeyword([]string{"scroll", "fixed", "local"}) },

	// --- Flexbox (Container & Items) ---
	"display": func() string {
		return RandomKeyword([]string{"block", "inline", "inline-block", "flex", "grid", "none"})
	},
	"flex-direction": func() string { return RandomKeyword([]string{"row", "row-reverse", "column", "column-reverse"}) },
	"justify-content": func() string {
		return RandomKeyword([]string{"flex-start", "flex-end", "center", "space-between", "space-around", "space-evenly"})
	},
	"align-items": func() string {
		return RandomKeyword([]string{"stretch", "flex-start", "flex-end", "center", "baseline"})
	},
	"flex-wrap":   func() string { return RandomKeyword([]string{"nowrap", "wrap", "wrap-reverse"}) },
	"flex-grow":   func() string { return strconv.Itoa(rand.Intn(5)) },
	"flex-shrink": func() string { return strconv.Itoa(rand.Intn(5)) },
	"order":       func() string { return strconv.Itoa(rand.Intn(10) - 5) },

	// --- Grid (Container & Items) ---
	"grid-template-columns": randomGridTemplate,
	"grid-template-rows":    randomGridTemplate,
	"grid-gap":              func() string { return randomLength(0, 50, "px", "em") },
	"grid-auto-flow":        func() string { return RandomKeyword([]string{"row", "column", "dense"}) },
	"grid-column":           func() string { return fmt.Sprintf("span %d", rand.Intn(3)+1) },
	"grid-row":              func() string { return fmt.Sprintf("span %d", rand.Intn(3)+1) },

	// --- Transitions ---
	"transition": func() string { return randomTransition() },

	// --- Visual Effects & Miscellaneous ---
	"opacity": func() string { return randomFloat(0.1, 1.0) },
	"cursor": func() string {
		return RandomKeyword([]string{"pointer", "default", "wait", "text", "move", "help", "crosshair"})
	},
	"visibility": func() string { return RandomKeyword([]string{"visible", "hidden", "collapse"}) },
	"transform": func() string {
		return fmt.Sprintf("rotate(%ddeg) scale(%.2f) skewX(%ddeg) translateX(%s)", rand.Intn(361)-180, rand.Float64()*1.5+0.5, rand.Intn(91)-45, randomLength(-100, 100, "px"))
	},
	"filter": func() string {
		return fmt.Sprintf("blur(%dpx) brightness(%.1f) saturate(%d%%)", rand.Intn(10), rand.Float64()*1.5+0.5, rand.Intn(201))
	},
	"clip-path":       randomClipPath,
	"object-fit":      func() string { return RandomKeyword([]string{"fill", "contain", "cover", "none", "scale-down"}) },
	"resize":          func() string { return RandomKeyword([]string{"none", "both", "horizontal", "vertical"}) },
	"scroll-behavior": func() string { return RandomKeyword([]string{"auto", "smooth"}) },

	// --- SVG Specific ---
	"fill":         RandomHexColor,
	"stroke":       RandomHexColor,
	"stroke-width": func() string { return strconv.Itoa(rand.Intn(10) + 1) },
}

var keys []string

func init() {
	keys = make([]string, 0, len(cssProperties))
	for k := range cssProperties {
		keys = append(keys, k)
	}
}

// --- Main CSS Generation Function ---

// GenerateRandomInlineCSS creates a single random CSS property and value guaranteed to be inline-compatible.
func GenerateRandomInlineCSS() string {
	// Select a random property
	property := keys[rand.Intn(len(keys))]

	// Generate a random value for the selected property
	value := cssProperties[property]()

	return property + ": " + value + ";"
}
