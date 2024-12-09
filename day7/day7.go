package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/utils"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Day7 struct{}

var _ aoc.Problem = (*Day7)(nil)

func part1Walk(result int, numbers []int) bool {
	return part1WalkCalculation(numbers[0], result, numbers[1:])
}

func part1WalkCalculation(curr int, result int, numbers []int) bool {
	if len(numbers) == 0 {
		return false
	}

	next := numbers[0]

	mult := curr * next
	add := curr + next
	if mult == result || add == result {
		return true
	}
	if mult > result && add > result {
		return false
	}

	if mult < result && part1WalkCalculation(mult, result, numbers[1:]) {
		return true
	}
	if add < result && part1WalkCalculation(add, result, numbers[1:]) {
		return true
	}

	return false
}

func (*Day7) Part1(r io.Reader) int {
	calculations := its.Map1To2(its.Filter(its.ReaderToIter(r), its.FilterEmptyLines), func(row string) (int, []int) {
		calc := strings.Split(row, ": ")
		result := utils.Must(strconv.Atoi(calc[0]))
		numbers := its.MapSlice(strings.Split(calc[1], " "), utils.MapToInt)
		return result, numbers
	})

	results := its.Filter2(calculations, func(result int, numbers []int) bool { return part1Walk(result, numbers) })

	return its.Reduce2(results, 0, func(acc int, result int, _ []int) int { return acc + result })
}

func part2Walk(result int, numbers []int) bool {
	return part2WalkCalculation(numbers[0], result, numbers[1:])
}

func part2WalkCalculation(curr int, result int, numbers []int) bool {
	if len(numbers) == 0 {
		return false
	}

	next := numbers[0]
	mult := curr * next
	add := curr + next

	combination := utils.MapToInt(fmt.Sprintf("%d%d", curr, next))

	if mult == result || add == result || combination == result {
		return true
	}
	if mult > result && add > result && combination > result {
		return false
	}

	if mult < result && part2WalkCalculation(mult, result, numbers[1:]) {
		return true
	}
	if add < result && part2WalkCalculation(add, result, numbers[1:]) {
		return true
	}
	if combination < result && part2WalkCalculation(combination, result, numbers[1:]) {
		return true
	}

	return false
}

func (*Day7) Part2(r io.Reader) int {
	calculations := its.Map1To2(its.Filter(its.ReaderToIter(r), its.FilterEmptyLines), func(row string) (int, []int) {
		calc := strings.Split(row, ": ")
		result := utils.Must(strconv.Atoi(calc[0]))
		numbers := its.MapSlice(strings.Split(calc[1], " "), func(s string) int { return utils.Must(strconv.Atoi(s)) })
		return result, numbers
	})

	results := its.Filter2(calculations, func(result int, numbers []int) bool { return part2Walk(result, numbers) })

	return its.Reduce2(results, 0, func(acc int, result int, _ []int) int { return acc + result })
}
