package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func PunctuationMark(s string) string{

    s = strings.ReplaceAll(s, "' ", " '")
    s = strings.ReplaceAll(s, " '", "'")
    
    return strings.TrimSpace(s)
}

func Punctuation(s string) string {
	sRune := []rune(s)

	for i, char := range sRune {
		if (char != ' ') && char != '\'' && !(unicode.IsLetter(char)) && !(unicode.IsDigit(char)) {
			if sRune[i-1] == ' ' {
				sRune[i], sRune[i-1] = sRune[i-1], sRune[i]
			}
		}
	}

	return strings.TrimSpace(string(sRune))
}

func Vowels(s string) string {
	vowels := "aeiouh"
	words := strings.Fields(s)
	res := ""

	for i, word := range words {
		if strings.EqualFold(word, "a") && i+1 < len(words) {
			nextWord := strings.ToLower(words[i+1])

			if strings.ContainsRune(vowels, []rune(nextWord)[0]) {
				word = "an"
			}
		}
		res += word + " "
	}

	return strings.TrimSpace(res)

}

func WriteToOutputFile(file, s string) error {
	openFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer openFile.Close()

	_, err = openFile.WriteString(s)
	if err != nil {
		return fmt.Errorf("failed to write to file: ", err)
	}

	fmt.Println("Data has been written to the file:", file)
	return nil
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
					words[i-1] = fmt.Sprint(num)
				} else {
					fmt.Println("Error:", err)
				}
			}
			words = append(words[:i], words[i+1:]...)

		case "(bin)":
			if i > 0 {
				num, err := strconv.ParseInt(words[i-1], 2, 64)
				if err == nil {
					words[i-1] = fmt.Sprint(num)
				} else {
					fmt.Println("Error:", err)
				}
			}
			words = append(words[:i], words[i+1:]...)

		case "(cap,":
			command := words[i] + " " + words[i+1]
			values := strings.Split(command, " ")
			num, err := strconv.Atoi(strings.TrimSpace(strings.Split(values[1], ")")[0]))

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

		case "(low,":
			command := words[i] + " " + words[i+1]
			values := strings.Split(command, " ")
			num, err := strconv.Atoi(strings.TrimSpace(strings.Split(values[1], ")")[0]))

			if err != nil {
				fmt.Println(err)
				return
			}

			for k := 1; k <= num; k++ {
				index := i - k

				if index >= 0 {
					words[index] = strings.ToLower(words[index])
				} else {
					break
				}
			}

			words = append(words[:i], words[i+2:]...)

		case "(up," :
			command := words[i] + " " + words[i+1]
			values := strings.Split(command, " ")
			number, err :=  strconv.Atoi(strings.TrimSpace(strings.Split(values[1], ")")[0]))

			if err != nil {
				fmt.Errorf("Error: ", err)
				return
			}

			for j := 1; j <= number; j++ {
				index := i-j

				if index >= 0 {
					words[index] = strings.ToUpper(words[index])
				}
			}

			words = append(words[:i], words[i+2:]...)
		}
	}

	WriteToOutputFile(outputfile, Punctuation(PunctuationMark((strings.Join(words, " ")))))
}
