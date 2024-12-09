package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/map2d"
	set "aoc-lib/slices"
	"io"
	"maps"
	"slices"
)

type Day6 struct{}

type Player struct {
	pos       map2d.Vector2
	direction map2d.Vector2
}

func (p *Player) RotateClockwise() {
	p.direction = p.direction.RotateClockwise()
}

func (p *Player) PeekStep() map2d.Vector2 {
	return p.pos.Add(p.direction)
}

func (p *Player) Step() {
	p.pos.AddMut(p.direction)
}

var _ aoc.Problem = (*Day6)(nil)

func (*Day6) Part1(r io.Reader) int {
	obstacleSet := make(set.Set[map2d.Vector2])
	fields := map2d.NewCellMap(r, map2d.CellMapFn)
	var player Player

	for cell := range fields.Iter() {
		if cell.Value == '^' {
			player.pos = cell.ExtractCoords()
			player.direction = map2d.NewVector2(0, -1)
		} else if cell.Value == '#' {
			obstacleSet.Set(cell.ExtractCoords())
		}
	}

	steppedSet := set.NewSetWithValues(player.pos)

	for fields.InBounce(player.pos) {
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

	return len(steppedSet) - 1
}

func (*Day6) Part2(r io.Reader) int {
	obstacleSet := make(set.Set[map2d.Vector2])
	var player Player

	fields := map2d.NewCellMap(r, map2d.CellMapFn)

	for cell := range fields.Iter() {
		if cell.Value == '^' {
			player.pos = cell.ExtractCoords()
			player.direction = map2d.NewVector2(0, -1)
		} else if cell.Value == '#' {
			obstacleSet.Set(cell.ExtractCoords())
		}
	}

	playerSteps := make([]map2d.Vector2, 0)

	testPlayer := Player{pos: player.pos, direction: player.direction}
	for fields.InBounce(testPlayer.pos) {
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

	obstacleTestSet := set.NewSet[map2d.Vector2]()

	return its.Reduce2(slices.All(playerSteps), 0, func(acc int, i int, step map2d.Vector2) int {
		if i == len(playerSteps)-1 || obstacleTestSet.Has(step) {
			return acc
		}

		obstacleTestSet.Set(step)
		newObstacles := maps.Clone(obstacleSet)
		newObstacles.Set(step)

		if testForLoop(fields, player, newObstacles) {
			return acc + 1
		}
		return acc
	})
}

func testForLoop(fields *map2d.CellMap[map2d.Cell], p Player, obstacles set.Set[map2d.Vector2]) bool {
	loopDetectionSet := set.NewSet[Player]()
	// steppedSet := set.NewSetWithValues(p.pos)

	for fields.InBounce(p.pos) {
		newPos := p.PeekStep()
		if obstacles.Has(newPos) {
			p.RotateClockwise()
			newPos = p.PeekStep()
			if obstacles.Has(newPos) {
				p.RotateClockwise()
			}
		}
		p.Step()
		// steppedSet.Set(p.pos)
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

func PrintMap(fields *map2d.CellMap[map2d.Cell], steppedSet set.Set[map2d.Vector2], obstacles set.Set[map2d.Vector2]) {
	fields.DebugPrint(func(c map2d.Cell) string {
		if steppedSet.Has(c.ExtractCoords()) {
			return "X"
		} else if obstacles.Has(c.ExtractCoords()) {
			return "#"
		} else {
			return "."
		}
	})
}
