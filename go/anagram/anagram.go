package anagram

import "strings"

func Detect(subject string, candidates []string) []string {

	var anagrams []string
	for _, candidate := range candidates {
		if isAnagram(strings.ToLower(subject), strings.ToLower(candidate)) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}

func isAnagram(subject string, candidate string) bool {
	if len(subject) != len(candidate) || subject == candidate {
		return false
	}

	for _, r := range subject {
		if strings.Count(subject, string(r)) != strings.Count(candidate, string(r)) {
			return false
		}
	}

	return true
}
