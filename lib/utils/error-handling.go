package utils

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Must2[T, K any](v1 T, v2 K, err error) (T, K) {
	if err != nil {
		panic(err)
	}
	return v1, v2
}
