package ui

import (
	"regexp"
	"strings"

	"github.com/fatih/color"
)

// HighlightTerms highlights all occurrences of search terms in the input slice.
func HighlightTerms(input *[][]string, searchTerms []string) {
	if len(searchTerms) == 0 {
		return
	}

	fgColor := color.New(color.FgBlack)
	bkColor := color.New(color.BgYellow)

	pattern := "(" + strings.Join(searchTerms, "|") + ")"
	re := regexp.MustCompile(pattern)

	for i := range *input {
		for j := range (*input)[i] {
			(*input)[i][j] = re.ReplaceAllStringFunc((*input)[i][j], func(match string) string {
				return bkColor.Sprint(fgColor.Sprint(match))
			})
		}
	}
}
