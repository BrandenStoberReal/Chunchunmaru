# Chunchunmaru
*A sophisticated, Go-based web scraper tarpit.*

Chunchunmaru is a defensive tool designed to protect websites from unauthorized scraping. It dynamically generates a web of convincing, believable, and resource-intensive pages that are delivered with adaptive throttling. The primary goal is to waste a scraper's resources (CPU, time, network), pollute its dataset with plausible-sounding junk, and ultimately make scraping a given site prohibitively expensive—all without being easily detected as an artificial trap.
![Crashed Tab Image](https://files.catbox.moe/gimu54.png)
## Core Principles
- **Resource Asymmetry:** The server is lightweight; content models are loaded once at startup, and page generation is computationally cheap. The system is designed to impose a high cost on the client (scraper) relative to the low cost on the server.
- **Plausible by Design:** Content is generated from thematic models to appear human-written, preventing the site from being easily identified and blacklisted.
- **Adaptive Throttling:** The system "drip-feeds" responses by sending small HTML chunks over time, tying up a scraper's connection pool and draining its resources.
- **Zero Dependencies:** All generated pages are single, self-contained HTML files. All styling (CSS) and logic (JS) are inlined.
## The Dynamic Aggression System
Chunchunmaru employs a dynamic aggression model that escalates its response based on client behavior. The `Aggression` score (0–100) is passed into templates, allowing conditional inclusion of more resource-intensive or deceptive components as the score rises.
---
# Macro Library
Macros are available in Go templates and grouped by category. All macros are registered in the template engine and can be used directly in HTML templates.
## Category 1: Content Generation
| Macro Signature                                                                        | Description |
|:---------------------------------------------------------------------------------------| :--- |
| `markovSentence length`                                                        | Generates a thematic sentence from a Markov chain model with `length` tokens. |
| `markovParagraphs count minSentences maxSentences minSentenceLength maxSentenceLength` | Generates `count` paragraphs using the Markov model. |
| `randomParagraphs count minSentences maxSentences minSentenceLength maxSentenceLength` | Generates `count` paragraphs of filler text using random words. |
| `randomSentence length`                                                                | Generates a nonsensical sentence from `length` random words. |
| `randomWord`                                                                           | Returns a single random word. |
| `randomString "type" length`                                                           | Generates a random string. Types: `username`, `email`, `uuid`, `hex`, `alphanum`. |
| `randomDate "layout" "start" "end"`                                                    | Generates a random, formatted date within a range. |
## Category 2: Structure & Composition
| Macro Signature | Description |
| :--- | :--- |
| `randomForm count styleCount` | Generates a `<form>` with `count` varied input fields and random inline styles. |
| `randomDefinitionData count length` | Returns a slice of `{Term, Def}` structs for use in definition lists. |
## Category 3: Styling
| Macro Signature | Description |
| :--- | :--- |
| `randomColor` | Returns a random hex color code. |
| `randomId "prefix" length` | Generates a unique, random ID string with the given prefix and length. |
| `randomClasses count` | Returns a string of `count` random utility-style class names. |
| `randomInlineStyle length` | Returns a `style` attribute string with `length` random CSS properties. |
| `randomCSSStyle length` | Returns a string of `length` random CSS property declarations (not wrapped in `style=""`). |
## Category 4: Link & Navigation
| Macro Signature | Description |
| :--- | :--- |
| `randomLink` | Generates a plausible relative URL path. |
| `randomQueryLink keyCount` | Generates a relative URL path and appends `keyCount` random query parameters. |
| `randomJSON depth maxElements maxStringLength` | Generates a random, nested JSON object string. |
## Category 5: Logic & Control
| Macro Signature | Description |
| :--- | :--- |
| `randomInt min max` | Returns a random integer in `[min, max)`. |
| `repeat count` | Returns a slice for use with `range`. |
| `list item1 item2 ...` | Returns a slice from the provided arguments. |
| `randomChoice slice` | Returns a random item from a slice. |
| `add a b` | Returns `a + b`. |
| `sub a b` | Returns `a - b`. |
| `div a b` | Returns `a / b` (integer division, returns 0 if b==0). |
| `mult a b` | Returns `a * b`. |
| `max a b` | Returns the maximum of `a` and `b`. |
| `min a b` | Returns the minimum of `a` and `b`. |
## Category 6: Computationally Expensive & Anti-Extraction
| Macro Signature | Description |
| :--- | :--- |
| `nestDivs count` | Generates `count` deeply nested `<div>`s with random classes and inline styles. |
| `randomComplexTable rows cols` | Generates an irregular `<table>` with random `colspan` and `rowspan` attributes. |
| `randomStyleBlock "type" count` | Generates a `<style>` block with `count` complex CSS rules under a unique parent class. Types: `utility`, `nested`, `complex`. Returns a struct with `{Style, Class}`. |
| `randomSVG "type"` | Generates an inline `<svg>` designed to be computationally expensive. Types: `fractal`, `filters`. |
| `randomCSSVars count` | Generates a `<style>` block defining a chain of interdependent CSS custom properties. |
| `jsInteractiveContent "type" content` | Generates a placeholder element and an inline script. The script performs a CPU-intensive calculation, then decodes and injects the `content` into the placeholder. `type` is the tag (e.g., `div`, `span`). |
---
## Template System
Templates (see `/templates/*.html`) use these macros to generate dynamic, aggression-scaled pages. The `.Aggression` variable (0–100) is passed to each template and can be used to conditionally scale up the complexity and resource cost of the generated page.
Example usage in a template:
```gotemplate
{{- $aggression := .Aggression -}}
{{- $mainDivs := add 4 (div $aggression 20) -}}
<body class="{{randomClasses (add 6 (div $aggression 10))}}" {{randomInlineStyle (add 6 (div $aggression 8))}}>
    {{nestDivs $mainDivs}}
    <h1>{{markovSentence "model.json" (add 12 (div $aggression 8))}}</h1>
    ...
</body>
```
---
## API Schema
The API exposes endpoints for server info and template management. See [internal/utilities/api.go](file://Chunchunmaru/internal/utilities/api.go) for struct definitions.

## Credits
**CTAG07** - Minor Math Contributions + Template Engine + Initial Concept