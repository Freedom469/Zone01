package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func punctuation(s string) string {
	sRune := []rune(s)

	for i, char := range sRune {
		if (char != ' ') && !(unicode.IsLetter(char)) && !(unicode.IsDigit(char)) {
			if sRune[i-1] == ' ' {
				sRune[i], sRune[i-1] = sRune[i-1], sRune[i]
			}
		}
	}

	return strings.TrimSpace(string(sRune))
}


func Vowels(s string) string {
	vowels := "aeiouAEIOUhH"
	words := strings.Fields(s)
	res := ""

	for i, word := range words {
		if strings.EqualFold(word, "a") && i+1 < len(words) {
			nextWord := words[i+1]

			if strings.ContainsRune(vowels, []rune(nextWord)[0]) {
				word = "an"
			}
		}
		res += word + " "
	}

	return strings.TrimSpace(res)
	
}
func WriteToOutputFile(file, s string) {
	openFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer openFile.Close()

	_, err = openFile.WriteString(s)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Data has been written to the file", file)
}

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

	words := strings.Fields(Vowels(string(filedata)))

	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]
		switch word {
		case "(cap)":
			if i > 0 {
				words[i-1] = strings.Title(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
		case "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
		case "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
		case "(hex)":
			if i > 0 {
				num, err := strconv.ParseInt(words[i-1], 16, 64)
				if err == nil {
					words[i-1] = fmt.Sprintf("%d", num)
				} else {
					fmt.Println("Error:", err)
				}
			}
			words = append(words[:i], words[i+1:]...)
		case "(bin)":
			if i > 0 {
				num, err := strconv.ParseInt(words[i-1], 2, 64)
				if err == nil {
					words[i-1] = fmt.Sprintf("%d", num)
				} else {
					fmt.Println("Error:", err)
				}
			}
			words = append(words[:i], words[i+1:]...)
		case "(cap,":
			var value string
		
			for j := i; j < len(words); j++ {
				if strings.HasSuffix(words[j], ")") {
					value += words[j]
					break
				} else {
					value += words[j] + " "
				}
				
			}
		
			val := strings.Split(value, ",")
			val = strings.Split(val[1], ")")
			Strnum := val[0]
		
			num, err := strconv.Atoi(strings.TrimSpace(Strnum))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		
			for k := 1; k <= num; k++ {
				index := i - k
				if index >= 0 {
					words[index] = strings.Title(words[index])
				} else {
					break
				}
			}
		
			words = append(words[:i], words[i+2:]...)
		
		case "(low," :
			var value string
			for j := i; j < len(words); j++ {
				if strings.HasSuffix(words[j], ")") {
					value += words[j]
				} else {
					value += words[j] + " "
				}
			}
			val := strings.Split(value, ",")
			val = strings.Split(val[1], ")")
			strNum := val[0]

			num, error := strconv.Atoi(strings.TrimSpace(strNum))

			if error != nil {
				fmt.Println(error)
				return
			}

			for k := 1; k <= num; k++ {
				index := i -k

				if index >= 0 {
					words[index] = strings.ToLower(words[index])
				} else {
					break
				}
			}

			words = append(words[:i], words[i+2:]...)
		}
	}

	WriteToOutputFile(outputfile, punctuation(strings.Join(words, " ")))
}