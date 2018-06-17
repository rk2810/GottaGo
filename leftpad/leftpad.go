package leftpad

import (
	"strings"
	"unicode/utf8"
)

var defaultChar = ' '

// Format takes in a string and an intand returns the string
// left-paddedwith spaces to the length of the int. If the
// string is already longer than specified length, the
// original string is returned.
func Format(s string, size int) string {
	return FormatRune(s, size, defaultChar)
}

// FormatRune takes in a string, an int, and a rune and returns the string
// left-padded with specified rune to length of the int. If the
// string is already longer than specified length, the
// original string is returned.
func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}
	out := strings.Repeat(string(r), size - l) + s
	return out
}
