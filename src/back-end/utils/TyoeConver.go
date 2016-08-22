package utils

import "strconv"

// Author: HackerZ
// Time  : 2016-7-5 14:32

// StrToInt64 Change String to Int64
// @param 	string
// @return 	int64s
func StrToInt64(s string) (i64 int64) {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return
}
