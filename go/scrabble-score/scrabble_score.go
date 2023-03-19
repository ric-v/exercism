package scrabble

import (
	"strings"
)

func Score(word string) (output int) {

	for _, letter := range strings.ToUpper(word) {
		switch letter {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			output += 1
		case 'D', 'G':
			output += 2
		case 'B', 'C', 'M', 'P':
			output += 3
		case 'F', 'H', 'V', 'W', 'Y':
			output += 4
		case 'K':
			output += 5
		case 'J', 'X':
			output += 8
		case 'Q', 'Z':
			output += 10
		}
	}
	return
}
