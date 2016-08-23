package utils

import (
	"strconv"
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
// @param s string
// @return json string
func StringsToJSON(str string) string {
	rs := []rune(str)
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
