package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[(INF|ERR|TRC|DBG|WRN|FTL)\])`)
	return re.FindString(text) != ""
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile("(<([-=*~]*)>)")
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var c int
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			c++
		}
	}
	return c
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`(end-of-line\d+)`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var searchRe = regexp.MustCompile(`(User\s+[A-Za-z\d]+)`)
	var nameRe = regexp.MustCompile(`([^User\s]+[A-Za-z\d]+)`)

	for i := range lines {
		searchStr := searchRe.FindString(lines[i])
		if searchStr != "" {
			name := nameRe.FindString(searchStr)
			fmt.Println(name)
			lines[i] = `[USR] ` + name + " " + lines[i]
		}
	}
	return lines
}
