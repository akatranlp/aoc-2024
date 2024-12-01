package main

import (
	"aoc-lib/fs"
	"aoc-lib/its"
	"aoc-lib/utils"
	"fmt"
	"slices"
)

func main() {
	leftList := make([]int, 0)
	rightList := make([]int, 0)
	rightMap := make(map[int]int)

	fs.ApplyToLines("day1/input.txt", func(msg string) {
		var number1, number2 int
		fmt.Sscanf(msg, "%d   %d", &number1, &number2)
		leftList = append(leftList, number1)
		rightList = append(rightList, number2)
		if v, ok := rightMap[number2]; ok {
			rightMap[number2] = v + 1
		} else {
			rightMap[number2] = 1
		}
	})

	slices.Sort(leftList)
	slices.Sort(rightList)

	sumPart1 := 0
	sumPart2 := 0
	for l, r := range its.Zip(slices.Values(leftList), slices.Values(rightList)) {
		distance := utils.IntAbs(l, r)
		sumPart1 += distance

		if v, ok := rightMap[l]; ok {
			sumPart2 += l * v
		}
	}
	fmt.Println("Part1:", sumPart1)
	fmt.Println("Part2:", sumPart2)
}
