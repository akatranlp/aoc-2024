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
	for _, row := range its.Filter2(its.ReaderToIter(r), its.FilterEmptyLines) {
		var number1, number2 int
		fmt.Sscanf(row, "%d   %d", &number1, &number2)
		leftList = append(leftList, number1)
		rightList = append(rightList, number2)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	sum := 0
	for l, r := range its.Zip(slices.Values(leftList), slices.Values(rightList)) {
		distance := utils.IntAbs(l, r)
		sum += distance
	}
	return sum
}

func (*Day1) Part2(r io.Reader) int {
	leftList := make([]int, 0)
	rightMap := make(map[int]int)

	for _, row := range its.ReaderToIter(r) {
		var number1, number2 int
		fmt.Sscanf(row, "%d   %d", &number1, &number2)
		leftList = append(leftList, number1)
		if v, ok := rightMap[number2]; ok {
			rightMap[number2] = v + 1
		} else {
			rightMap[number2] = 1
		}
	}

	sum := 0
	for _, v := range leftList {
		if c, ok := rightMap[v]; ok {
			sum += c * v
		}
	}
	return sum
}
