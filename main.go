package main

import (
	"fmt"
	"os"
	"reloaded/utils"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: <inputfile> <outputfile>")
		return
	}
	inputfile := args[0]
	outputfile := args[1]
	filedata, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("Error while reading file:", err)
		return
	}
	words := strings.Fields(reloaded.Vowels(string(filedata)))
	modified_string := reloaded.ModifyString(words)
	reloaded.WriteToOutputFile(outputfile, reloaded.PunctuationMark(reloaded.Punctuation(modified_string)))
}
