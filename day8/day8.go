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
	field := map2d.NewMap2D()
	for _, row := range its.Filter2(its.ReaderToIter(r), its.FilterEmptyLines) {
		field.AppendString(row)
	}

	signalMap := make(map[byte]set.Set[map2d.Vector2])
	signalMap[counterSignal] = make(set.Set[map2d.Vector2])

	for p := range its.Filter(field.IterEachField(), func(p map2d.Point) bool { return signalRegex.MatchString(string(p.Value)) }) {
		v, ok := signalMap[p.Value]
		if !ok {
			v = make(set.Set[map2d.Vector2])
		}
		v.Set(p.Extract())
		signalMap[p.Value] = v
	}

	keys := slices.Collect(maps.Keys(signalMap))
	slices.Sort(keys)

	for _, k := range keys {
		set := signalMap[k]
		if len(set) < 2 {
			continue
		}

		keys := slices.Collect(maps.Keys(set))
		slices.SortFunc(keys, func(a, b map2d.Vector2) int {
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
			// fmt.Printf("%+v -> %+v = %+v\n", first, second, difference)
			c1 := first.Add(difference)

			if field.InBounce(c1) {
				// fmt.Printf("C1: %+v\n", c1)
				signalMap[counterSignal].Set(c1)
				// field.Set(map2d.Point{c1.X, c1.Y, counterSignal})
			}
			c2 := second.Sub(difference)
			if field.InBounce(c2) {
				// fmt.Printf("C2: %+v\n", c2)
				signalMap[counterSignal].Set(c2)
				// field.Set(map2d.Point{c2.X, c2.Y, counterSignal})
			}
		}
	}

	// field.DebugPrint()

	return len(signalMap[counterSignal])
}

func (*Day8) Part2(r io.Reader) int {
	field := map2d.NewMap2D()
	for _, row := range its.Filter2(its.ReaderToIter(r), its.FilterEmptyLines) {
		field.AppendString(row)
	}

	signalMap := make(map[byte]set.Set[map2d.Vector2])
	signalMap[counterSignal] = make(set.Set[map2d.Vector2])

	for p := range its.Filter(field.IterEachField(), func(p map2d.Point) bool { return signalRegex.MatchString(string(p.Value)) }) {
		v, ok := signalMap[p.Value]
		if !ok {
			v = make(set.Set[map2d.Vector2])
		}
		v.Set(p.Extract())
		signalMap[p.Value] = v
	}

	keys := slices.Collect(maps.Keys(signalMap))
	slices.Sort(keys)

	for _, k := range keys {
		set := signalMap[k]
		if len(set) < 2 {
			continue
		}

		keys := slices.Collect(maps.Keys(set))
		slices.SortFunc(keys, func(a, b map2d.Vector2) int {
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

			// fmt.Printf("%+v -> %+v = %+v\n", first, second, difference)
			for c1 := first.Add(difference); field.InBounce(c1); c1 = c1.Add(difference) {
				// fmt.Printf("C1: %+v\n", c1)
				signalMap[counterSignal].Set(c1)
				// field.Set(map2d.Point{c1.X, c1.Y, counterSignal})
			}

			for c2 := second.Sub(difference); field.InBounce(c2); c2 = c2.Sub(difference) {
				// fmt.Printf("C2: %+v\n", c2)
				signalMap[counterSignal].Set(c2)
				// field.Set(map2d.Point{c2.X, c2.Y, counterSignal})
			}
		}
	}

	// field.DebugPrint()
	return len(signalMap[counterSignal])
}
