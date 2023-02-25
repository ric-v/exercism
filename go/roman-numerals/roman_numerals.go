package romannumerals

import (
	"fmt"
)

func ToRomanNumeral(input int) (output string, err error) {

	if input <= 0 || input >= 4000 {
		return "", fmt.Errorf("%d is out of range", input)
	}
	var numericUnits = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var romanNumerals = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i, v := range numericUnits {
		for input >= v {
			output += romanNumerals[i]
			input -= v
		}
	}

	return output, nil
}
