# Chunchunmaru
*A sophisticated, Go-based web scraper tarpit.*

Chunchunmaru is a defensive tool designed to protect websites from unauthorized scraping. It dynamically generates a web of convincing, believable, and resource-intensive pages that are delivered with adaptive throttling. The primary goal is to waste a scraper's resources (CPU, time, network), pollute its dataset with plausible-sounding junk, and ultimately make scraping a given site prohibitively expensive—all without being easily detected as an artificial trap.

![Crashed Tab Image](https://files.catbox.moe/gimu54.png)

### Core Principles

*   **Resource Asymmetry:** The server is exceptionally lightweight. Content models are loaded once at startup, and page generation is computationally cheap. The system is designed to impose a high cost on the client (scraper) relative to the low cost on the server.
*   **Plausible by Design:** Believability is paramount. Content is generated from thematic models to appear human-written, preventing the site from being easily identified and blacklisted.
*   **Adaptive Throttling:** Inspired by honeypots like Nepenthes, the system "drip-feeds" responses by sending small HTML chunks over time. This technique ties up a scraper's connection pool, exploits short timeouts, and drains its resources.
*   **Zero Dependencies:** All generated pages are single, self-contained HTML files. All styling (CSS) and logic (JS) are inlined to frustrate scrapers that ignore external assets and to ensure each page is an atomic unit of work.

### The Dynamic Aggression System

Instead of serving from a fixed pool of HTML pages, Chunchunmaru employs a dynamic aggression model that escalates its response based on client behavior. This allows for a graceful, targeted defense that punishes persistent scrapers while having a minimal impact on casual or legitimate visitors.

#### How It Works

1.  **IP Reputation Tracking:** The Go backend maintains a simple, in-memory ledger of incoming IP addresses and their request counts.
2.  **Aggression Score Calculation:** For each request, the system calculates an **`Aggression`** score (e.g., an integer from 0 to 100). This score is derived from the IP's request frequency and history. A new visitor has an aggression score of 0; a persistent scraper that has requested hundreds of pages will have a much higher score.
3.  **Context-Aware Templating:** This `Aggression` score is passed directly into the Go template as a variable.
4.  **Conditional Page Generation:** Template authors can use simple Go template logic (`{{if}}`, `{{range}}`, etc.) to conditionally include more resource-intensive or deceptive components as the `Aggression` score rises. A single template can therefore produce a wide spectrum of pages, from lightweight and simple to a complex, CPU-melting nightmare.

This creates a feedback loop: the more that a scraper requests, the more hostile and expensive each subsequent response becomes.

---

# The Macro Library

## Category 1: Content Generation
| Function Signature                                                                             | Description                                                                                                                                                                                              |
|:-----------------------------------------------------------------------------------------------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `markovSentence "model" length`                                                                | Generates a single, thematic sentence from a specified Markov chain model with `length` tokens.                                                                                                          |
| `markovParagraphs "model" count minSentences maxSentences minSentenceLength maxSentenceLength` | Generates `count` paragraphs of thematic text using the specified Markov model. Returns raw text without `<p>` tags.                                                                                     |
| `randomParagraphs count minSentences maxSentences minSentenceLength maxSentenceLength`         | Generates `count` paragraphs of lower-quality filler text using random words. Returns raw text without `<p>` tags.                                                                                       |
| `randomSentence length`                                                                        | Generates a single nonsensical sentence from `length` random words.                                                                                                                                      |
| `randomWord`                                                                                   | Returns a single random word.                                                                                                                                                                            |
| `randomString "type" length`                                                                   | Generates a random string based on a pre-defined format or wordlist. Length affects types that are just random generation of a charset<br/>Types: `name`, `username`, `email`, `uuid`, `hex`, `alphanum` |
| `randomDate "layout" "start" "end"`                                                            | Generates a random, formatted date within a range to provide realistic timestamps.                                                                                                                       |

## Category 2: Structure & Composition
| Function Signature           | Description                                                                                                                                                                                                                                                                       |
|:-----------------------------|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `randomForm count`           | Generates a complete `<form>` element with `count` varied input fields.                                                                                                                                                                                                           |
| `randomDefinitionData count` | **Returns a data structure, not HTML.** Generates a slice of data (e.g., `[]struct{Term string; Definition string}`).                                                                                                                                                             |

## Category 3: Styling
| Function Signature         | Description                                                                                                                        |
|:---------------------------|:-----------------------------------------------------------------------------------------------------------------------------------|
| `randomColor`              | Returns a random hex color code.                                                                                                   |
| `randomId "prefix" count`  | Generates a unique, random ID attribute string  with `length` characters (e.g., `id="prefix-a8b2c"`) for linking elements.         |
| `randomClasses count`      | Returns a string of `count` space-separated, random utility-style class names.                                                     |
| `randomInlineStyle length` | Returns a `style` attribute string with `length` random CSS properties.                                                            |
| `randomCSSStyle length`    | Returns a randomly generated CSS style property. Note this does not include the tag or any formatting, purely the property setter. |

## Category 4: Link & Navigation
| Function Signature                             | Description                                                                                                                                                                          |
|:-----------------------------------------------|:-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `randomLink`                                   | Generates a plausible relative URL path, with a prefix chosen from a pre-configured list in the Go backend.                                                                          |
| `randomQueryLink keyCount`                     | Generates a relative URL path (via `randomLink`) and appends `keyCount` random query parameters. E.g., `/path?a4f1=b9d2&c3a8=e7f6`.                                                  |
| `randomJSON depth maxElements maxStringLength` | Generates a random, nested JSON object string with a depth of `depth`, elements until it reaches `maxElements`, and random string keys or values of strength up to `maxStringLength` |

## Category 5: Logic & Control
| Function Signature     | Description                                  |
|:-----------------------|:---------------------------------------------|
| `randomInt min max`    | Returns a random integer.                    |
| `repeat count`         | Returns a slice for use with `range`.        |
| `list item1 item2 ...` | Returns a slice from the provided arguments. |
| `randomChoice slice`   | Returns a random item from a slice.          |
| `add a b`              | Returns a + b.                               |
| `sub a b`              | Returns a - b.                               |
| `div a b`              | Returns a / b. (Or 0 if b == 0)              |
| `mult a b`             | Returns a * b.                               |

## Category 6: Computationally Expensive & Anti-Extraction
| Function Signature                    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
|:--------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `nestDivs count`                      | <br/>Generates `count` deeply nested `<div>`s, each with random classes.<br/>Attacks both DOM parser performance (due to depth) and CSS selector matching engines (due to class permutations).<br/>Low bandwidth, high client CPU cost.                                                                                                                                                                                                                                      |
| `randomComplexTable rows cols`        | <br/>Generates an irregular `<table>` with random `colspan` and `rowspan` attributes.<br/>Stresses a scraper's layout engine and data extraction logic, which often assumes regular table structures.                                                                                                                                                                                                                                                                        |
| `randomStyleBlock "type" count`       | <br/>Generates a `<style>` block with `count` complex CSS rules scoped under a unique parent class, and returns a struct with data {Style, Class}.<br/>Types: `utility` (many simple classes), `nested` (deep child/sibling selectors), `complex` (attribute and pseudo-class selectors).                                                                                                                                                                                    |
| `randomSVG "type"`                    | <br/>Generates an inline `<svg>` designed to be computationally expensive for the rendering engine, while keeping the data payload reasonable.<br/>Types: `fractal` (path data for a recursive fractal), `filters` (a stack of complex SVG filters applied to a simple element).                                                                                                                                                                                             |
| `randomCSSVars count`                 | <br/>Generates a `<style>` block defining a long chain of interdependent CSS Custom Properties (e.g., `--v2: calc(var(--v1) * 1.5)`).<br/>Forces a recursive calculation workload on the CSS engine.                                                                                                                                                                                                                                                                         |
| `jsInteractiveContent "type" content` | <br/>Generates a placeholder element and an inline script. The script first performs a CPU-intensive calculation, then decodes the `content` (e.g., from Base64) and uses standard DOM manipulation to inject it into the placeholder.<br/>This mimics a modern "hydrating" web component, making the CPU drain and content obfuscation a single, plausible, and necessary step for the scraper to get the content.<br/>Type dictates the placeholder (e.g., `div`, `span`). |

### Credits
CTAG07 - Minor Math Contributions + Initial Concept
