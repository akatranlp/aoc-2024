package aoc

import (
	"aoc-lib/utils"
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

type Problem interface {
	Part1(r io.Reader) int
	Part2(r io.Reader) int
}

func Run(fileName string, p Problem, problems ...int) {
	file := utils.Must(os.Open(fileName))
	defer file.Close()
	buf := bytes.NewReader(utils.Must(io.ReadAll(file)))

	if len(problems) == 0 {
		run1(buf, p)
		buf.Seek(0, io.SeekStart)
		run2(buf, p)
	} else {
		if problems[0] == 1 {
			run1(buf, p)
		} else if problems[0] == 2 {
			run2(buf, p)
		} else {
			panic("false problem number")
		}
	}
}

func run1(r io.Reader, p Problem) {
	start := time.Now()
	answer := p.Part1(r)
	duration := time.Since(start)
	fmt.Printf("Part1: %d - elapsed: %s\n", answer, duration)
}

func run2(r io.Reader, p Problem) {
	start := time.Now()
	answer := p.Part2(r)
	duration := time.Since(start)
	fmt.Printf("Part2: %d - elapsed: %s\n", answer, duration)
}
