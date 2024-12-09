package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/utils"
	"io"
	"slices"
	"strings"
)

type Day2 struct{}

var _ aoc.Problem = (*Day2)(nil)

func isValid(list []int) bool {
	if len(list) < 1 {
		return true
	}

	order := list[0]-list[1] < 0
	for v1, v2 := range its.Window2(slices.Values(list)) {
		asc := v1-v2 < 0

		if asc && !order || !asc && order {
			return false
		}

		difference := utils.IntAbs(v1, v2)
		if difference < 1 || difference > 3 {
			return false
		}
	}
	return true
}

func (*Day2) Part1(r io.Reader) int {
	return its.Reduce(its.Filter(its.ReaderToIter(r), its.FilterEmptyLines), 0, func(acc int, row string) int {
		if isValid(its.MapSlice(strings.Split(row, " "), utils.MapToInt)) {
			return acc + 1
		}
		return acc
	})
}

func (*Day2) Part2(r io.Reader) int {
	return its.Reduce(its.Filter(its.ReaderToIter(r), its.FilterEmptyLines), 0, func(acc int, row string) int {
		numbers := its.MapSlice(strings.Split(row, " "), utils.MapToInt)
		for removeIndex := range numbers {
			if isValid(its.RemoveIndexNew(numbers, removeIndex)) {
				return acc + 1
			}
		}
		return acc
	})
}
