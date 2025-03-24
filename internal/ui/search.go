package ui

import (
	"regexp"
	"strings"
)

var (
	backgroundColor = "\033[43m"
	foregroundColor = "\033[30m" // Black foreground
)

// HighlightTerms highlights all occurrences of search terms in the input slice.
func HighlightTerms(input *[][]string, searchTerms []string) {
	if len(searchTerms) == 0 {
		return
	}

	pattern := "(" + strings.Join(searchTerms, "|") + ")"
	re := regexp.MustCompile(pattern)

	for i := range *input {
		for j := range (*input)[i] {
			(*input)[i][j] = re.ReplaceAllStringFunc((*input)[i][j], func(match string) string {
				return backgroundColor + foregroundColor + match + "\033[0m"
			})
		}
	}
}
