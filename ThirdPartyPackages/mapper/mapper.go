package mapper

import (
	"github.com/rylans/getlang"
)

// Greet function expects a string and returns a string depending on language code info
func Greet(s string) string {
	info := getlang.FromString(s)
	switch info.LanguageCode() {
	case "en":
		return "Hello!"
	case "de":
		return "Guten Tag!"
	case "fr":
		return "Bonjour!"
	default:
		return "I dunno your language."
	}
}
