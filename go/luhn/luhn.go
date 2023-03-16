package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {

	// remove all white space
	id = strings.ReplaceAll(id, " ", "")

	// if len is less than 2 characters, invalid
	if len(id) < 2 {
		return false
	}

	// set sum as 0 and position of 2nd char as len(id)%2
	var sum, doublePos int32 = 0, int32(len(id)) % 2
	for _, n := range id {

		// if not a valid digit, invalid
		if !unicode.IsDigit(n) {
			return false
		}

		// if current char is the valid 2nd character in series,
		// do n * 2 and if n*2 is > 9, subtract 9 from the result and add to sum
		// else simply add n to sum
		if doublePos%2 == 0 {
			sum += double(n - 48)
		} else {
			sum += n - 48
		}

		// update the double position
		doublePos++
	}

	// if sum is divisible by 10, its valid, else invalid
	return sum%10 == 0
}

// double the number, if its > 9, subtract the result by 9 and return
func double(i int32) int32 {

	// double the digit
	i = i * 2

	// if its > 9, subtract 9 from the result
	if i > 9 {
		i -= 9
	}
	return i
}
