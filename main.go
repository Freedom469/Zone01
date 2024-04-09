package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "it (cap) was the BEST (low) of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap)"

	words := strings.Split(input, " ")

	for i := 0; i < len(words); i++ { 
		if words[i] == "(cap)" && i > 0 { 
			words[i-1] = strings.Title(words[i-1])
			
			words = append(words[:i], words[i+1:]...) 
		} else if words[i] == "(up)" && i > 0 { 
			words[i-1] = strings.ToUpper(words[i-1]) 
			words = append(words[:i], words[i+1:]...) 
		} else if words[i] == "(low)" && i > 0 { 
			words[i-1] = strings.ToLower(words[i-1]) 
			words = append(words[:i], words[i+1:]...) 
		}
	}

	fmt.Println(strings.Join(words, " ")) 

}