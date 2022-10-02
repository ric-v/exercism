package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for _, c := range word {
		if c == '-' || c == ' ' {
			continue
		}
		if strings.Count(word, string(c)) > 1 {
			return false
		}
	}
	return true
}
