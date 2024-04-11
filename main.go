package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Vowel(s string) string {
	// Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h.
	//  (Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!").

	// Define vowels and h as a string
	vowels := "aeiouAEIOUhH"
	// Split the input string into words
	words := strings.Fields(s)
	// fmt.Println(words)
	// Initialize the result string
	result := ""

	// Iterate over each word in the input
	for i, word := range words {
		// Check if the word is "a" and there's a next word
		if strings.EqualFold(word, "a") && i+1 < len(words) {
			nextWord := words[i+1]
			// Check if the first character of the next word is a vowel or "h"
			if strings.ContainsRune(vowels, []rune(nextWord)[0]) {
				// If so, change "a" to "an"
				word = "an"
			}
		}
		// Append the modified or original word to the result string
		result += word + " "
	}

	// Trim any trailing whitespace and return the result
	return strings.TrimSpace(result)
}

func Punctuation(s string) string {
	// The punctuation mark ' will always be found with another instance of it and they should be placed to the
	// right and left of the word in the middle of them, without any spaces.
	//  (Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'")

	words := strings.Split(s, "'")
	for i := range words {
		words[i] = strings.TrimSpace(words[i])
		// fmt.Println(words[i])
	}
	return strings.Join(words, "'")
}

func main() {
	input := "it (up) was the best of TIMES (up, 3) , it was the worst of times (cap, 7), it was the age of wisdom, it was the age of foolishness Hello World  (up, 10) "
	// input := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
	// input := "There is no greater agony than bearing a untold story inside you."
	// input := "I am exactly how they describe me: ' awesome  don't awesome'"

	EditedInput := Punctuation(Vowel(input))

	words := strings.Split(EditedInput, " ")

	for i, word := range words {
		if word != "" {
			if word == "(cap)" {
				words[i-1] = strings.Title(words[i-1])
				words = append(words[:i], words[i+1:]...)
			} else if word == "(low)" {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
			} else if word == "(up)" {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
			} else if word == "(up," && i < len(words) {
				w := strings.Trim(words[i+1], ")")
				// fmt.Println(w)

				num, error := strconv.Atoi(w)

				if error != nil {
					panic(error)
				}

				// 			// if num > i {
				// 			// 	len :=
				// 			// }

				for j := 0; j <= num; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}

				words = append(words[:i], words[i+2:]...)

				// 			// fmt.Println(w)
			} else if word == "(low," {
						w := strings.Trim(strings.Trim(words[i+1], words[i][1:]), ")")
						num, error := strconv.Atoi(w)

						if error != nil {
							panic(error)
						}

						// if num > i {
						// 	len :=
						// }

						for j := 0; j <= num; j++ {
							if words[i-j] != ","{
								words[i-j] = strings.ToLower(words[i-j])
							}
		}

					words = append(words[:i], words[i+2:]... )

					// fmt.Println(w)
				} else if word == "(cap," {
					w := strings.Trim(strings.Trim(words[i+1], words[i][1:]), ")")
					num, error := strconv.Atoi(w)

					if error != nil {
						panic(error)
					}

					// if num > i {
					// 	len :=
					// }

					for j := 0; j <= num; j++ {
						if words[i-j] != ","{
							words[i-j] = strings.Title(words[i-j])
						}

					}

					words = append(words[:i], words[i+2:]... )

					// fmt.Println(w)
				} else if word == "(hex)" {
					num, error := strconv.ParseInt(words[i-1], 16, 64)

					if error != nil {
						panic(error)
					}

					words[i-1] = fmt.Sprint(num)

					words = append(words[:i], words[i+1:]... )

				}  else if word == "(bin)" {
					num, error := strconv.ParseInt(words[i-1], 2, 64)

					if error != nil {
						panic(error)
					}

					words[i-1] = fmt.Sprint(num)

					words = append(words[:i], words[i+1:]... )
				}
			}
	}

	fmt.Println(strings.Join(words, " "))
}
