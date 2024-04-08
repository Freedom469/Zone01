package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: <inputfile>")
		return
	}

	inputfile := args[0]

	data, err := ioutil.ReadFile(inputfile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var words []string
	word := ""

	for _, char := range data {
		if unicode.IsSpace(rune(char)) {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	if word != "" {
		words = append(words, word)
	}

	fmt.Println(words[0])
}
