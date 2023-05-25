package inkwell

import (
	"regexp"
	"strings"
)

func PuncConv(text string) string {
	puncRegex := regexp.MustCompile(` *(\,|\.|\!|\?|\:|\;)+ *`)
	matches := puncRegex.FindAllStringIndex(text, -1)
	// reverse loop cuz we using ranges here
	for i := len(matches) - 1; i >= 0; i-- {
		text = HandlePunc(matches[i], text)
	}
	return text
}

func HandlePunc(rng []int, text string) string {
	var replc string
	startIndex, endIndex := rng[0], rng[1]
	match := text[startIndex:endIndex]
	// var replc string
	if (endIndex == len(text)) || (endIndex < len(text) && text[endIndex] == '\n') {
		replc = strings.TrimSpace(match)
	} else {
		replc = strings.TrimSpace(match) + " "
	}
	return text[:startIndex] + replc + text[endIndex:]
}
