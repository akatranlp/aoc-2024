package its

import "iter"

func Zip[T, K any](seq1 iter.Seq[T], seq2 iter.Seq[K]) iter.Seq2[T, K] {
	return func(yield func(T, K) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()
		next2, stop2 := iter.Pull(seq2)
		defer stop2()
		for {
			v1, ok1 := next1()
			v2, ok2 := next2()

			if !ok1 || !ok2 {
				return
			}

			if !yield(v1, v2) {
				return
			}
		}
	}
}
