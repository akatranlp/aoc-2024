package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/utils"
	"fmt"
	"io"
	"slices"
)

type Day1 struct{}

var _ aoc.Problem = (*Day1)(nil)

func (*Day1) Part1(r io.Reader) int {
	leftList := make([]int, 0)
	rightList := make([]int, 0)
	for row := range its.Filter(its.ReaderToIter(r), its.FilterEmptyLines) {
		var number1, number2 int
		fmt.Sscanf(row, "%d   %d", &number1, &number2)
		leftList = append(leftList, number1)
		rightList = append(rightList, number2)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	return its.Reduce2(its.Zip(slices.Values(leftList), slices.Values(rightList)), 0, func(acc, l, r int) int {
		return acc + utils.IntAbs(l, r)
	})
}

func (*Day1) Part2(r io.Reader) int {
	leftList := make([]int, 0)
	rightMap := make(map[int]int)

	for row := range its.Filter(its.ReaderToIter(r), its.FilterEmptyLines) {
		var number1, number2 int
		fmt.Sscanf(row, "%d   %d", &number1, &number2)
		leftList = append(leftList, number1)
		if v, ok := rightMap[number2]; ok {
			rightMap[number2] = v + 1
		} else {
			rightMap[number2] = 1
		}
	}

	return its.Reduce(slices.Values(leftList), 0, func(acc, v int) int {
		if c, ok := rightMap[v]; ok {
			return acc + c*v
		}
		return acc
	})
}
