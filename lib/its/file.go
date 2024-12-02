package its

import (
	"bufio"
	"io"
	"iter"
)

func ReaderToIter(r io.Reader, splits ...bufio.SplitFunc) iter.Seq2[int, string] {
	scanner := bufio.NewScanner(r)
	split := bufio.ScanLines
	if len(splits) > 0 {
		split = splits[0]
	}
	scanner.Split(split)
	return func(yield func(i int, s string) bool) {
		i := 0
		for scanner.Scan() {
			msg := scanner.Text()
			if !yield(i, msg) {
				return
			}
			i++
		}
	}
}
