package helpers

import (
	"strings"
	"unicode"
)

func ToCapitalCase(s string) string {
	w := strings.Fields(s)

	for i, word := range w {
		if len(w) > 0{
		runes := []rune(strings.ToLower(word))
		runes[0] = unicode.ToUpper(runes[0])
		w[i] = string(runes)
		}
	}
	return strings.Join(w," ")
}