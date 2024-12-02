package main

import (
	"bytes"
	"testing"
)

var part1Test = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func TestDay1(t *testing.T) {
	day1 := Day1{}
	t.Run("part 1", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 11
		actual := day1.Part1(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 31
		actual := day1.Part2(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})
}
