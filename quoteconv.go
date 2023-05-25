package inkwell

import (
	"fmt"
	"regexp"
	"strings"
)

func QouteConv(text string) string {
	qouteRegex := regexp.MustCompile(`'[^']*'`)
	return qouteRegex.ReplaceAllStringFunc(text, FormatQoute)
}
func FormatQoute(match string) string {
	return fmt.Sprintf("'%s'", strings.TrimSpace(match[1:len(match)-1]))
}
