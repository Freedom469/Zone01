package reloaded

import (
	"strings"
	"strconv"
	"fmt"
)

func Low(words []string){
for i := 0; i < len(words); i++ {
		if words[i] != " "{
			if words[i] == "(low)" && i > 0 { 
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...) 

			} else if words[i] == "(low," && i > 0{
				// fmt.Println(words[i:i+2])
				number := words[i:i+2][1]
				number = strings.Trim(number, ")")
				num, error := strconv.Atoi(number)

				if error != nil {
					fmt.Println("Error", error)
					continue
				}

				for j := 0; j < num && i-j-1 >= 0; j++ {
					words[i-j-1] = strings.ToLower(words[i-j-1])
					words = append(words[:i], words[i+1:]... )
				}
				// fmt.Println(num)
			}
		}
	}
}