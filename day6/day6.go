package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/slices"
	"fmt"
	"io"
	"math"
)

type Day6 struct{}

type Vector2 struct{ x, y int }

func (v *Vector2) RotateClockwise() *Vector2 {
	v.x = int(float64(v.x)*math.Cos(math.Pi/2) - (float64(v.y) * (math.Sin(math.Pi / 2))))
	v.y = int(float64(v.x)*math.Sin(math.Pi/2) - (float64(v.y) * (math.Cos(math.Pi / 2))))
	return v
}

type Player struct {
	pos       *Vector2
	direction *Vector2
}

func (p *Player) RotateClockwise() {
	if p.direction.x == 1 {
		p.direction.x = 0
		p.direction.y = 1
	} else if p.direction.x == -1 {
		p.direction.x = 0
		p.direction.y = -1
	} else if p.direction.y == 1 {
		p.direction.x = -1
		p.direction.y = 0
	} else {
		p.direction.x = 1
		p.direction.y = 0
	}
}

func (p *Player) PeekStep() Vector2 {
	return Vector2{p.pos.x + p.direction.x, p.pos.y + p.direction.y}
}

func (p *Player) Step() {
	p.pos.x += p.direction.x
	p.pos.y += p.direction.y
}

var _ aoc.Problem = (*Day6)(nil)

func inBounce(rows, cols int, player Player) bool {
	return player.pos.y >= 0 && player.pos.y < rows && player.pos.x >= 0 && player.pos.x < cols
}

func (*Day6) Part1(r io.Reader) int {
	obstacleSet := make(slices.Set[Vector2])
	map2d := make([][]byte, 0)
	var player Player

	for _, row := range its.Filter2(its.ReaderToIter(r), its.FilterEmptyLines) {
		map2d = append(map2d, []byte(row))
	}
	rows := len(map2d)
	cols := len(map2d[0])

	for y, row := range map2d {
		for x, char := range row {
			if char == '^' {
				player.pos = &Vector2{x, y}
				player.direction = &Vector2{0, -1}
			}
			if char == '#' {
				obstacleSet.Set(Vector2{x, y})
			}
		}
	}

	steppedSet := make(slices.Set[Vector2])
	steppedSet.Set(*player.pos)

	for inBounce(rows, cols, player) {
		newPos := player.PeekStep()
		if obstacleSet.Has(newPos) {
			player.RotateClockwise()
			newPos = player.PeekStep()
			if obstacleSet.Has(newPos) {
				player.RotateClockwise()
			}
		}
		player.Step()
		steppedSet.Set(*player.pos)
	}

	for y, row := range map2d {
		for x := range row {
			pos := Vector2{x, y}
			if steppedSet.Has(pos) {
				// fmt.Print("X")
			} else if obstacleSet.Has(pos) {
				// fmt.Print("#")
			} else {
				// fmt.Print(".")
			}
		}
		// fmt.Println()
	}

	return len(steppedSet) - 1
}

func (*Day6) Part2(r io.Reader) int {
	fmt.Println("Part2 not implemented")
	return -1
}
