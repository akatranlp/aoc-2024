package utils

import "strconv"

func IntAbs(l, r int) int {
	v := l - r
	if v < 0 {
		return -v
	}
	return v
}

func MapToInt(s string) int { return Must(strconv.Atoi(s)) }
