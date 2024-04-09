package main

import (
	"fmt"
	"strings"

	reloaded "reloaded/functions"
	// "strconv"
)

func main() {
	input := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."

	words := strings.Split(input, " ")

	for i := 0; i < len(words); i++ {
		if words[i] != " " {
			if words[i] == "(cap)" && i > 0 {
				reloaded.Capitalize(words)

				// fmt.Println(num)
			} else if words[i] == "(up)" && i > 0 {
				reloaded.Upper(words)
			} else if words[i] == "(low)" && i > 0 {
				reloaded.Low(words)
			}
		}
	}

	fmt.Println(strings.Join(words, " "))
}
