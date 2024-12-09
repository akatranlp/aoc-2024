package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/map2d"
	set "aoc-lib/slices"
	"io"
	"maps"
	"regexp"
	"slices"
)

type Day8 struct{}

var _ aoc.Problem = (*Day8)(nil)

const counterSignal = '#'

var signalRegex = regexp.MustCompile("[a-zA-Z0-9]")

func (*Day8) Part1(r io.Reader) int {
	field := map2d.NewCellMap(r, map2d.CellMapFn)

	signalMap := make(map[byte]set.Set[map2d.Vector2])
	signalMap[counterSignal] = set.NewSet[map2d.Vector2]()

	for p := range its.Filter(field.Iter(), func(p map2d.Cell) bool { return signalRegex.MatchString(string(p.Value)) }) {
		v, ok := signalMap[p.Value]
		if !ok {
			v = set.NewSet[map2d.Vector2]()
		}
		v.Set(p.ExtractCoords())
		signalMap[p.Value] = v
	}

	keys := slices.Sorted(maps.Keys(signalMap))

	for _, c := range keys {
		set := signalMap[c]
		if len(set) < 2 {
			continue
		}

		keys := slices.SortedFunc(maps.Keys(set), func(a, b map2d.Vector2) int {
			if a.Y < b.Y {
				return -1
			} else if a.X < b.X {
				return -1
			}
			return 0
		})

		for c := range its.AllCombinations(keys, false) {
			first, second := c.L, c.R
			difference := first.Sub(second)
			c1 := first.Add(difference)

			if field.InBounce(c1) {
				signalMap[counterSignal].Set(c1)
			}
			c2 := second.Sub(difference)
			if field.InBounce(c2) {
				signalMap[counterSignal].Set(c2)
			}
		}
	}

	return len(signalMap[counterSignal])
}

func (*Day8) Part2(r io.Reader) int {
	field := map2d.NewCellMap(r, map2d.CellMapFn)

	signalMap := make(map[byte]set.Set[map2d.Vector2])
	signalMap[counterSignal] = make(set.Set[map2d.Vector2])

	for p := range its.Filter(field.Iter(), func(p map2d.Cell) bool { return signalRegex.MatchString(string(p.Value)) }) {
		v, ok := signalMap[p.Value]
		if !ok {
			v = make(set.Set[map2d.Vector2])
		}
		v.Set(p.ExtractCoords())
		signalMap[p.Value] = v
	}

	keys := slices.Sorted(maps.Keys(signalMap))

	for _, k := range keys {
		set := signalMap[k]
		if len(set) < 2 {
			continue
		}

		keys := slices.SortedFunc(maps.Keys(set), func(a, b map2d.Vector2) int {
			if a.Y < b.Y {
				return -1
			} else if a.X < b.X {
				return -1
			}
			return 0
		})

		for c := range its.AllCombinations(keys, false) {
			first, second := c.L, c.R
			difference := first.Sub(second)
			signalMap[counterSignal].Set(first)
			signalMap[counterSignal].Set(second)

			for c1 := first.Add(difference); field.InBounce(c1); c1 = c1.Add(difference) {
				signalMap[counterSignal].Set(c1)
			}

			for c2 := second.Sub(difference); field.InBounce(c2); c2 = c2.Sub(difference) {
				signalMap[counterSignal].Set(c2)
			}
		}
	}

	return len(signalMap[counterSignal])
}
