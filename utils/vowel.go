package reloaded

import (
	"strings"
)

func Vowels(s string) string {
	vowels := "aeiouhAEIOUH"
	words := strings.Fields(s)
	res := ""

	for i, word := range words {
		if word == "a" || word == "A" && i+1 < len(words) {
			nextWord := words[i+1]

			if strings.ContainsRune(vowels, []rune(nextWord)[0]) {
				if word == "a" {
					word = "an"
				} else if word == "A" {
					word = "An"
				}
			}
		}

		res += word + " "
	}
	return res

}
