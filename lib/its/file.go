package its

import (
	"bufio"
	"bytes"
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

func SplitByBlocks(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 2, data[:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}
