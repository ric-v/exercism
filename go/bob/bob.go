// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	if isShoutingQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	}

	if isShouting(remark) {
		return "Whoa, chill out!"
	}

	if isQuestion(remark) {
		return "Sure."
	}

	return "Whatever."
}

func isShoutingQuestion(remark string) bool {
	return isShouting(remark) && isQuestion(remark)
}

func isShouting(remark string) (yelling bool) {
	var letters = 0
	var upperLetters = 0
	for _, c := range remark {
		if c >= 'A' && c <= 'Z' {
			upperLetters++
		}
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			letters++
		}
	}
	return letters > 0 && upperLetters == letters
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}
