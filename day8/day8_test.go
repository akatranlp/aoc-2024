package main

import (
	"bytes"
	"testing"
)

var part1Test = `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`

func TestDay8(t *testing.T) {
	day8 := Day8{}
	t.Run("part 1", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 14
		actual := day8.Part1(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 34
		actual := day8.Part2(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})
}
