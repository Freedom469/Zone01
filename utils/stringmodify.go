package reloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func ModifyString(words []string) string {

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
			num, err := strconv.Atoi(strings.TrimSpace(strings.Split(words[i+1], ")")[0]))

			if err != nil {
				fmt.Println("Error:", err)
				return ""
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
			num, err := strconv.Atoi(strings.TrimSpace(strings.Split(words[i+1], ")")[0]))

			if err != nil {
				fmt.Println(err)
				return ""
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

		case "(up,":
			number, err := strconv.Atoi(strings.TrimSpace(strings.Split(words[i+1], ")")[0]))

			if err != nil {
				fmt.Println("Error: ", err)
				return ""
			}

			for j := 1; j <= number; j++ {
				index := i - j

				if index >= 0 {
					words[index] = strings.ToUpper(words[index])
				}
			}

			words = append(words[:i], words[i+2:]...)
		}
	}

	return strings.Join(words, " ")
}
