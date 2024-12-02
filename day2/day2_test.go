package main

import (
	"bytes"
	"testing"
)

var part2Test = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestDay2(t *testing.T) {
	day2 := Day2{}
	t.Run("part 1", func(t *testing.T) {
		input := bytes.NewBufferString(part2Test)

		expected := 2
		actual := day2.Part1(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		input := bytes.NewBufferString(part2Test)

		expected := 4
		actual := day2.Part2(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})
}
