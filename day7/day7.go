package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/utils"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Day7 struct{}

var _ aoc.Problem = (*Day7)(nil)

type Calculation struct {
	result  int
	numbers []int
}

func startWalk(c Calculation) bool {
	first := c.numbers[0]
	c.numbers = c.numbers[1:]
	return walkCalculation(first, c)
}

func walkCalculation(curr int, c Calculation) bool {
	if len(c.numbers) == 0 {
		return false
	}

	next := c.numbers[0]

	mult := curr * next
	add := curr + next
	if mult == c.result || add == c.result {
		return true
	}
	if mult > c.result && add > c.result {
		return false
	}

	if mult < c.result {
		newC := Calculation{c.result, c.numbers[1:]}
		if walkCalculation(mult, newC) {
			return true
		}
	}
	if add < c.result {
		newC := Calculation{c.result, c.numbers[1:]}
		if walkCalculation(add, newC) {
			return true
		}
	}

	return false
}

func (*Day7) Part1(r io.Reader) int {
	calculations := make([]Calculation, 0)
	for row := range its.Filter(its.ReaderToIter(r), its.FilterEmptyLines) {
		calc := strings.Split(row, ": ")
		result := utils.Must(strconv.Atoi(calc[0]))
		numbers := its.MapSlice(strings.Split(calc[1], " "), func(s string) int { return utils.Must(strconv.Atoi(s)) })
		calculations = append(calculations, Calculation{result: result, numbers: numbers})
	}

	results := its.Filter(slices.Values(calculations), func(calc Calculation) bool { return startWalk(calc) })

	return its.Reduce(results, 0, func(acc int, c Calculation) int { return acc + c.result })
}

func startWalkPart2(c Calculation) bool {
	first := c.numbers[0]
	c.numbers = c.numbers[1:]
	return walkCalculationPart2(first, c)
}

func walkCalculationPart2(curr int, c Calculation) bool {
	if len(c.numbers) == 0 {
		return false
	}

	next := c.numbers[0]
	mult := curr * next
	add := curr + next

	combination := utils.Must(strconv.Atoi(fmt.Sprintf("%d%d", curr, next)))

	if mult == c.result || add == c.result || combination == c.result {
		return true
	}
	if mult > c.result && add > c.result && combination > c.result {
		return false
	}

	if mult < c.result {
		newC := Calculation{c.result, c.numbers[1:]}
		if walkCalculationPart2(mult, newC) {
			return true
		}
	}
	if add < c.result {
		newC := Calculation{c.result, c.numbers[1:]}
		if walkCalculationPart2(add, newC) {
			return true
		}
	}
	if combination < c.result {
		newC := Calculation{c.result, c.numbers[1:]}
		if walkCalculationPart2(combination, newC) {
			return true
		}
	}

	return false
}

func (*Day7) Part2(r io.Reader) int {
	calculations := make([]Calculation, 0)
	for row := range its.Filter(its.ReaderToIter(r), its.FilterEmptyLines) {
		calc := strings.Split(row, ": ")
		result := utils.Must(strconv.Atoi(calc[0]))
		numbers := its.MapSlice(strings.Split(calc[1], " "), func(s string) int { return utils.Must(strconv.Atoi(s)) })
		calculations = append(calculations, Calculation{result: result, numbers: numbers})
	}

	results := its.Filter(slices.Values(calculations), func(calc Calculation) bool { return startWalkPart2(calc) })

	calcs := slices.Collect(results)

	return its.Reduce(slices.Values(calcs), 0, func(acc int, c Calculation) int { return acc + c.result })
}
