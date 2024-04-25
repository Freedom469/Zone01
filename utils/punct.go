package reloaded

import (
	"strings"
	"unicode"
)

func Punctuation(s string) string {
	sRune := []rune(s)
	var output []rune

	for i, char := range sRune {
		if (char != ' ') && char != '\'' && !(unicode.IsLetter(char)) && !(unicode.IsDigit(char)) {
			if sRune[i-1] == ' ' && i >= 0 {
				sRune[i], sRune[i-1] = sRune[i-1], sRune[i]
			}
		}
	}

	for i, char := range sRune {
		if char == ' ' && sRune[i-1] == ' ' {
			continue
		}
		output = append(output, sRune[i])
	}
	return strings.TrimSpace(string(output))
}
