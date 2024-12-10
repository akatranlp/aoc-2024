package main

import (
	"aoc-lib/aoc"
	"aoc-lib/map2d"
	"aoc-lib/slices"
	"io"
)

type Day10 struct{}

var _ aoc.Problem = (*Day10)(nil)

var allNumbers = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var neighbors = []map2d.Vector2{
	map2d.NewVector2(-1, 0),
	map2d.NewVector2(0, 1),
	map2d.NewVector2(1, 0),
	map2d.NewVector2(0, -1),
}

func walkMap(fields *map2d.CellMap[map2d.Cell], currPos map2d.Vector2, missingNumbers []byte, nineSet slices.Set[map2d.Vector2]) int {
	cell := fields.Get(currPos)
	if cell.Value != missingNumbers[0] {
		return 0
	}

	if len(missingNumbers) == 1 {
		nineSet.Set(currPos)
		return 1
	}

	var count int
	for _, neighbor := range neighbors {
		newPos := currPos.Add(neighbor)
		if !fields.InBounce(newPos) {
			continue
		}
		count += walkMap(fields, newPos, missingNumbers[1:], nineSet)
	}
	return count
}

func (*Day10) Part1(r io.Reader) int {
	fields := map2d.NewCellMap(r, func(x, y int, value byte) map2d.Cell { return map2d.Cell{X: x, Y: y, Value: value - '0'} })

	startingPoints := make([]map2d.Vector2, 0)

	for cell := range fields.Iter() {
		if cell.Value == 0 {
			startingPoints = append(startingPoints, cell.ExtractCoords())
		}
	}

	var count int
	for _, p := range startingPoints {
		nineSet := slices.NewSet[map2d.Vector2]()
		walkMap(fields, p, allNumbers, nineSet)
		count += len(nineSet)
	}

	return count
}

func (*Day10) Part2(r io.Reader) int {
	fields := map2d.NewCellMap(r, func(x, y int, value byte) map2d.Cell { return map2d.Cell{X: x, Y: y, Value: value - '0'} })

	startingPoints := make([]map2d.Vector2, 0)

	for cell := range fields.Iter() {
		if cell.Value == 0 {
			startingPoints = append(startingPoints, cell.ExtractCoords())
		}
	}

	var count int
	for _, p := range startingPoints {
		nineSet := slices.NewSet[map2d.Vector2]()
		count += walkMap(fields, p, allNumbers, nineSet)
	}

	return count
}
