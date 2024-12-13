// Source: https://stackoverflow.com/questions/32081808/strip-all-whitespace-from-a-string
package string

import (
	"strings"
	"unicode"
)

func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func SpaceFieldsJoin(str string) string {
	return strings.Join(strings.Fields(str), "")
}

func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
