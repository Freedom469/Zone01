package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// QuotationMark replaces spaces around single quotes with no spaces.
func QuotationMark(s string) string {
	s = strings.ReplaceAll(s, "' ", " '")
	s = strings.ReplaceAll(s, " '", "'")
	return strings.TrimSpace(s)
}

// Punctuation moves punctuation marks to the preceding space.
func Punctuation(s string) string {
    sRune := []rune(s)
    var new []rune
//iterating over the srune
    for i, char := range sRune {
        if (char != ' ') && char != '\'' && !(unicode.IsLetter(char)) && !(unicode.IsDigit(char)) {
            if i > 0 && sRune[i-1] == ' ' {
                sRune[i], sRune[i-1] = sRune[i-1], sRune[i]
            }
        }
    }
    
    for i := 0; i < len(sRune); i++ {
        if i < len(sRune)-1 && sRune[i] == ' ' && sRune[i+1] == ' ' {
            continue
        }
        new = append(new, sRune[i])
    }
    
    return strings.TrimSpace(string(new))
}


// Vowels checks for the presence of vowels or h after 'a' and replaces it with 'an'.
func Vowels(s string) string {
    vowels := "aeiouAEIOUhH"
    words := strings.Fields(s)
    res := ""

    for i, word := range words {
        // Check if there's a next word
        if i+1 < len(words) {
            nextWord := words[i+1]
            // Check if the current word is "a" or "A" and the next word starts with a vowel
            if (word == "a" || word == "A") && strings.ContainsRune(vowels, rune(nextWord[0])) {
                // Replace "a" with "an" or "A" with "An"
                if word == "a" {
                    word = "an"
                } else {
                    word = "An"
                }
            }
        }
        res += word + " "
    }

    return strings.TrimSpace(res)
}

// WriteToOutputFile writes the processed string to the output file.
func WriteToOutputFile(file, s string) error {
	openFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer openFile.Close()

	_, err = openFile.WriteString(s)
	if err != nil {
		fmt.Println("failed to write to file: ", err)
	}

	fmt.Println("Data has been written to the file:", file)
	return nil
}

func ProcessWordCommands(words []string) string {
	// Process commands in reverse order
	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]
		switch word {
		case "(cap)":
			// Capitalize the preceding word
			if i > 0 {
				words[i-1] = strings.Title(words[i-1])
			}
			// Remove the command from the slice
			words = append(words[:i], words[i+1:]...)

		case "(up)":
			// Convert the preceding word to uppercase
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			// Remove the command from the slice
			words = append(words[:i], words[i+1:]...)

		case "(low)":
			// Convert the preceding word to lowercase
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}
			// Remove the command from the slice
			words = append(words[:i], words[i+1:]...)

		case "(hex)":
			// Convert the preceding word from hexadecimal to decimal
			if i > 0 {
				num, err := strconv.ParseInt(words[i-1], 16, 64)
				if err == nil {
					words[i-1] = fmt.Sprint(num)
				} else {
					fmt.Println("Error:", err)
				}
			}
			// Remove the command from the slice
			words = append(words[:i], words[i+1:]...)

		case "(bin)":
			// Convert the preceding word from binary to decimal
			if i > 0 {
				num, err := strconv.ParseInt(words[i-1], 2, 64)
				if err == nil {
					words[i-1] = fmt.Sprint(num)
				} else {
					fmt.Println("Error:", err)
				}
			}
			// Remove the command from the slice
			words = append(words[:i], words[i+1:]...)

		case "(cap,":
			// Capitalize multiple preceding words based on the following number
			if i > 0 {
				command := words[i] + " " + words[i+1]
				values := strings.Split(command, " ")
				num, err := strconv.Atoi(strings.TrimSpace(strings.Split(values[1], ")")[0]))

				if err != nil {
					fmt.Println("Error:", err)
					return "nil"
				}

				for k := 1; k <= num; k++ {
					index := i - k
					if index >= 0 {
						words[index] = strings.Title(words[index])
					} else {
						break
					}
				}
			}
			// Remove the command from the slice
			words = append(words[:i], words[i+2:]...)

		case "(low,":
			// Convert multiple preceding words to lowercase based on the following number
			if i > 0 {
				command := words[i] + " " + words[i+1]
				values := strings.Split(command, " ")
				num, err := strconv.Atoi(strings.TrimSpace(strings.Split(values[1], ")")[0]))

				if err != nil {
					fmt.Println(err)
					return "nil"
				}

				for k := 1; k <= num; k++ {
					index := i - k

					if index >= 0 {
						words[index] = strings.ToLower(words[index])
					} else {
						break
					}
				}
			}
			// Remove the command from the slice
			words = append(words[:i], words[i+2:]...)

		case "(up,":
			// Convert multiple preceding words to uppercase based on the following number
			if i > 0 {
				command := words[i] + " " + words[i+1]
				values := strings.Split(command, " ")
				number, err := strconv.Atoi(strings.TrimSpace(strings.Split(values[1], ")")[0]))

				if err != nil {
					fmt.Println("Error: ", err)
					return "nil"
				}

				for j := 1; j <= number; j++ {
					index := i - j

					if index >= 0 {
						words[index] = strings.ToUpper(words[index])
					}
				}
			}
			// Remove the command from the slice
			words = append(words[:i], words[i+2:]...)
		}
	}

	// Join the modified words into a single string
	return strings.Join(words, " ")
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: <inputfile> <outputfile>")
		return
	}
	inputfile := args[0]
	outputfile := args[1]
	// Read input file
	filedata, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("Error while reading file:", err)
		return
	}
	// Process the file data
	words := strings.Fields(Vowels(string(filedata)))
	response := ProcessWordCommands(words)
	// Write processed data to the output file
	WriteToOutputFile(outputfile, QuotationMark(Punctuation((response))))
}