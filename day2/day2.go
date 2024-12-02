package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/utils"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Day2 struct{}

var _ aoc.Problem = (*Day2)(nil)

func isValid(list []int) bool {
	if len(list) < 1 {
		return true
	}

	order := list[0]-list[1] < 0
	for i := 0; i < len(list)-1; i++ {
		v1, v2 := list[i], list[i+1]
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
	sum := 0
	for _, row := range its.Filter2(its.ReaderToIter(r), its.FilterEmptyLines) {
		values := strings.Split(row, " ")
		numbers := its.Map(slices.Values(values), func(s string) int {
			return utils.Must(strconv.Atoi(s))
		})
		list := slices.Collect(numbers)
		if isValid(list) {
			sum += 1
		}
	}

	return sum
}

func (*Day2) Part2(r io.Reader) int {
	sum := 0
	for _, row := range its.Filter2(its.ReaderToIter(r), its.FilterEmptyLines) {
		values := strings.Split(row, " ")
		numbers := its.Map(slices.Values(values), func(s string) int {
			return utils.Must(strconv.Atoi(s))
		})

		list := slices.Collect(numbers)
		for removeIndex := range list {
			newList := make([]int, len(list))
			copy(newList, list)
			newList = append(newList[:removeIndex], newList[removeIndex+1:]...)
			if isValid(newList) {
				sum += 1
				break
			}
		}
	}

	return sum
}
