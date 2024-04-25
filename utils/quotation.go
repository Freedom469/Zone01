package reloaded

import (
	"strings"
)

func PunctuationMark(s string) string {

	s = strings.ReplaceAll(s, "' ", " '")
	s = strings.ReplaceAll(s, " '", "'")

	return strings.TrimSpace(s)
}
