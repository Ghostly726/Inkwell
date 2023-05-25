package inkwell

import (
	"regexp"
	"strconv"
)

func ExtractNumber(tag string) int {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(tag)
	if match == "" {
		return 1
	}
	num, err := strconv.Atoi(match)
	if err != nil {
		return 0
	}
	return num
}

func IsTag(input string) (string, int, bool) {
	re := regexp.MustCompile(`\((bin|hex|(up|low|cap)(,\s*\d+)?)\)`)
	match := re.FindString(input)
	if match == "" {
		return "", 1, false
	}
	tag := match[1 : len(match)-1]
	num := ExtractNumber(tag)
	return tag, num, true
}
