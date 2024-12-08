package main

import (
	"bytes"
	"testing"
)

var part1Test = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

func TestDay6(t *testing.T) {
	day6 := Day6{}
	t.Run("part 1", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 41
		actual := day6.Part1(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		input := bytes.NewBufferString(part1Test)

		expected := 6
		actual := day6.Part2(input)

		if expected != actual {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, actual)
		}
	})
}

func TestPlayer(t *testing.T) {
	t.Run("Rotate Clockwise", func(t *testing.T) {
		player := Player{direction: Vector2{0, -1}}
		expected := Vector2{1, 0}
		player.RotateClockwise()
		if player.direction != expected {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, player.direction)
		}

		player.RotateClockwise()
		expected = Vector2{0, 1}
		if player.direction != expected {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, player.direction)
		}

		player.RotateClockwise()
		expected = Vector2{-1, 0}
		if player.direction != expected {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, player.direction)
		}

		player.RotateClockwise()
		expected = Vector2{0, -1}
		if player.direction != expected {
			t.Fatalf("ERROR: expected %d actual %d\n", expected, player.direction)
		}
	})
}
