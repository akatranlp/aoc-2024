package main

import (
	"bytes"
	"testing"
)

var part1Test = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`

func TestDay10(t *testing.T) {
	day10 := Day10{}
	t.Run("part 1", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 36
		actual := day10.Part1(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 81
		actual := day10.Part2(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})
}
