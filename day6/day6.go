package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/slices"
	"fmt"
	"io"
	"maps"
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
	pos       Vector2
	direction Vector2
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
				player.pos = Vector2{x, y}
				player.direction = Vector2{0, -1}
			}
			if char == '#' {
				obstacleSet.Set(Vector2{x, y})
			}
		}
	}

	steppedSet := make(slices.Set[Vector2])
	steppedSet.Set(player.pos)

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
		steppedSet.Set(player.pos)
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
				player.pos = Vector2{x, y}
				player.direction = Vector2{0, -1}
			}
			if char == '#' {
				obstacleSet.Set(Vector2{x, y})
			}
		}
	}

	playerSteps := make([]Vector2, 0)

	testPlayer := Player{pos: player.pos, direction: player.direction}
	for inBounce(rows, cols, testPlayer) {
		newPos := testPlayer.PeekStep()
		if obstacleSet.Has(newPos) {
			testPlayer.RotateClockwise()
			newPos = testPlayer.PeekStep()
			if obstacleSet.Has(newPos) {
				testPlayer.RotateClockwise()
			}
		}
		testPlayer.Step()
		playerSteps = append(playerSteps, testPlayer.pos)
	}

	obstacleTestSet := slices.NewSet[Vector2]()

	var count int
	for i, v := range playerSteps {
		if i == len(playerSteps)-1 {
			continue
		}
		if obstacleTestSet.Has(v) {
			continue
		}
		obstacleTestSet.Set(v)
		newObstacles := maps.Clone(obstacleSet)
		newObstacles.Set(v)

		if testForLoop(map2d, player, rows, cols, newObstacles) {
			count += 1
		}
	}

	return count
}

func testForLoop(map2d [][]byte, p Player, rows, cols int, obstacles slices.Set[Vector2]) bool {
	loopDetectionSet := slices.NewSet[Player]()
	steppedSet := make(slices.Set[Vector2])
	steppedSet.Set(p.pos)

	for inBounce(rows, cols, p) {
		newPos := p.PeekStep()
		if obstacles.Has(newPos) {
			p.RotateClockwise()
			newPos = p.PeekStep()
			if obstacles.Has(newPos) {
				p.RotateClockwise()
			}
		}
		p.Step()
		steppedSet.Set(p.pos)
		if loopDetectionSet.Has(p) {
			// PrintMap(map2d, steppedSet, obstacles)
			// var s string
			// fmt.Scanln(&s)
			return true
		}
		loopDetectionSet.Set(p)
	}

	return false
}

func PrintMap(map2d [][]byte, steppedSet slices.Set[Vector2], obstacles slices.Set[Vector2]) {
	for y, row := range map2d {
		for x := range row {
			pos := Vector2{x, y}
			if steppedSet.Has(pos) {
				fmt.Print("X")
			} else if obstacles.Has(pos) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
