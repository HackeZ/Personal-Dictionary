package utils

import (
	"strconv"

	"github.com/russross/blackfriday"
)

// Author: HackerZ
// Time  : 2016-7-5 14:32

// Atoi64 is shorthand for ParseInt(s, 10, 0).
// @param s string
// @return int64, error
func Atoi64(s string) (int64, error) {
	i64, err := strconv.ParseInt(s, 10, 0)
	return i64, err
}

// StringsToJSON Format String to JSON.
// @param input string
// @return json string
func StringsToJSON(input string) string {
	rs := []rune(input)
	jsons := ""

	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}

// StringsToMarkdown Format String to Markdown.
// @param input string
// @return output string
func StringsToMarkdown(input string) (output string) {
	outputByte := blackfriday.MarkdownBasic([]byte(input))
	output = string(outputByte)
	return
}
