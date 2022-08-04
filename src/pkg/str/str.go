package str

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var space = regexp.MustCompile(`\s+`)
var NameRegex = `^[a-zA-ZğüşöçİĞÜŞÖÇ]+$`

func UpperCaseFirstLetters(s string) string {
	c := cases.Title(language.Turkish)
	return c.String(strings.ToLowerSpecial(unicode.TurkishCase, strings.TrimSpace(space.ReplaceAllString(s, " "))))
}
