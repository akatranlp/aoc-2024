package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/map2d"
	"io"
	"slices"
)

type Day4 struct{}

var _ aoc.Problem = (*Day4)(nil)

var neighbors = []int{-1, 0, 1}

func searchWord(fields *map2d.CellMap[map2d.Cell], pos map2d.Vector2, word string) int {
	var count int
	for _, dRow := range neighbors {
		for _, dCol := range neighbors {
			if dRow == 0 && dCol == 0 {
				continue
			}
			dir := map2d.NewVector2(dCol, dRow)
			if searchWordDirection(fields, pos, dir, word) {
				count++
			}
		}
	}
	return count
}

func searchWordDirection(fields *map2d.CellMap[map2d.Cell], pos, dir map2d.Vector2, word string) bool {
	return its.All2(slices.All([]byte(word)), func(idx int, char byte) bool {
		newPos := pos.Add(dir.Scale(idx))
		if !fields.InBounce(newPos) {
			return false
		}
		if fields.Get(newPos).Value != char {
			return false
		}
		return true
	})
}

func (*Day4) Part1(r io.Reader) int {
	fields := map2d.NewCellMap(r, map2d.CellMapFn)
	return its.Reduce(fields.Iter(), 0, func(acc int, cell map2d.Cell) int {
		if cell.Value == 'X' {
			return acc + searchWord(fields, cell.ExtractCoords(), "XMAS")
		}
		return acc
	})
}

var (
	topLeft     = map2d.NewVector2(-1, -1)
	topRight    = map2d.NewVector2(-1, 1)
	bottomLeft  = map2d.NewVector2(1, -1)
	bottomRight = map2d.NewVector2(1, 1)
)

func checkX(fields *map2d.CellMap[map2d.Cell], pos map2d.Vector2) bool {
	var newPos1, newPos2 map2d.Vector2
	var cell1, cell2 map2d.Cell

	newPos1 = pos.Add(topLeft)
	newPos2 = pos.Add(bottomRight)

	if !fields.InBounce(newPos1) || !fields.InBounce(newPos2) {
		return false
	}
	cell1 = fields.Get(newPos1)
	cell2 = fields.Get(newPos2)
	if !(cell1.Value == 'M' && cell2.Value == 'S') && !(cell1.Value == 'S' && cell2.Value == 'M') {
		return false
	}

	newPos1 = pos.Add(topRight)
	newPos2 = pos.Add(bottomLeft)
	if !fields.InBounce(newPos1) || !fields.InBounce(newPos2) {
		return false
	}
	cell1 = fields.Get(newPos1)
	cell2 = fields.Get(newPos2)
	if !(cell1.Value == 'M' && cell2.Value == 'S') && !(cell1.Value == 'S' && cell2.Value == 'M') {
		return false
	}
	return true
}

func (d *Day4) Part2(r io.Reader) int {
	fields := map2d.NewCellMap(r, map2d.CellMapFn)

	return its.Reduce(fields.Iter(), 0, func(acc int, cell map2d.Cell) int {
		if cell.Value == 'A' && checkX(fields, cell.ExtractCoords()) {
			return acc + 1
		}
		return acc
	})
}
