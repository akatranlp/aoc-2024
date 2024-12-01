package utils

func IntAbs(l, r int) int {
	v := l - r
	if v < 0 {
		return -v
	}
	return v
}
