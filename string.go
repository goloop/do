package g

import "strings"

const (
	// Numbers is a string of all numbers.
	Numbers = "1234567890"

	// Letters is a string of all ASCII letters.
	Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Symbols is a string of all special ASCII symbols.
	Symbols = ".,:;!?/\\|`~@#$%^&*()_+-=<>"

	// Quotes is a string of all quotes.
	Quotes = "'\"`"

	// Brackets is a string of all breakers.
	Brackets = "()[]<>{}"

	// Whitespaces is a set of characters that are used as whitespaces.
	Whitespaces = " \t\b"

	// Breakers is a set of characters that are used as breakers.
	Breakers = "\n\r\v\f"

	// Hidden contains part of Whitespaces group and all Breakers.
	Hidden = "\t\b\n\r\v\f"
)

// Weed removes characters by the patterns from the whole string.
//
// It is a utility function that helps you to 'clean up' your strings,
// removing any unwanted characters (weeds). It's allegory, just as a
// gardener would go through a field removing any unwanted plants,
// this function iterates over the string and plucks out any characters
// specified as patterns.
//
// By default, if no patterns are specified, it removes the most common
// breakers characters (e.g. newline, tab etc.). However, you can specify
// your own patterns to fit your needs. The function uses an efficient
// mapping approach to achieve this, making it effective for clearing
// large strings.
//
// Example usage:
//
//	g.Weed("Hello\t World")                  // Output: "Hello World"
//	g.Weed(" i@ goloop.one", g.Whitespaces)  // Output: "i@goloop.one"
//	g.Weed("+380 (96) 123 4567", " +()")     // Output: "380961234567"
func Weed(s string, patterns ...string) string {
	var sb strings.Builder

	if len(patterns) == 0 {
		patterns = []string{Hidden}
	}

	// Characters to be deleted.
	cbd := make(map[rune]struct{})
	for _, pattern := range patterns {
		for _, ch := range pattern {
			cbd[ch] = struct{}{}
		}
	}

	for _, r := range s {
		if _, ok := cbd[r]; !ok {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

// Trim removes all leading and trailing occurrences of specified characters
// from the string. If no characters are provided, it removes leading and
// trailing whitespace.
//
// It can be used to tidy up user input or to normalize strings for
// consistent processing.
//
// The function utilizes the TrimFunc function from the standard strings
// package, which makes it efficient for clearing large strings.
//
// Example usage:
//
//	g.Trim(" Hello\t World\r\n")              // Output: "Hello\t World"
//	g.Trim("    Go Loop   ")                  // Output: "Go Loop"
//	g.Trim(" i@ goloop.one ", g.Whitespaces)  // Output: "i@ goloop.one"
func Trim(s string, patterns ...string) string {
	if len(patterns) == 0 {
		patterns = []string{Whitespaces, Breakers}
	}

	// Characters to be deleted.
	cbd := make(map[rune]struct{})
	for _, pattern := range patterns {
		for _, ch := range pattern {
			cbd[ch] = struct{}{}
		}
	}

	return strings.TrimFunc(s, func(r rune) bool {
		_, exists := cbd[r]
		return exists
	})
}
