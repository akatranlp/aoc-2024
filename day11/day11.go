package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/utils"
	"io"
	"strconv"
	"strings"
)

type Day11 struct{}

var _ aoc.Problem = (*Day11)(nil)

func (*Day11) Part1(r io.Reader) int {
	numbers := its.MapSlice(strings.Fields(string(utils.Must(io.ReadAll(r)))), utils.MapToInt)

	for range 25 {
		newSlice := make([]int, 0, len(numbers)*2)

		for _, v := range numbers {
			if v == 0 {
				newSlice = append(newSlice, 1)
				continue
			}
			numString := strconv.Itoa(v)
			if len(numString)%2 == 0 {
				firstHalf := utils.MapToInt(numString[:len(numString)/2])
				secondHalf := utils.MapToInt(numString[len(numString)/2:])
				newSlice = append(newSlice, firstHalf)
				newSlice = append(newSlice, secondHalf)
				continue
			}
			newSlice = append(newSlice, v*2024)
		}

		numbers = newSlice
	}

	return len(numbers)
}

func (*Day11) Part2(r io.Reader) int {
	numbers := its.MapSlice(strings.Fields(string(utils.Must(io.ReadAll(r)))), utils.MapToInt)
	numbersMap := make(map[int]int)
	for _, n := range numbers {
		numbersMap[n]++
	}

	for range 75 {
		newNumbers := make(map[int]int)

		for number, count := range numbersMap {
			if number == 0 {
				newNumbers[1] += count
				continue
			}
			numString := strconv.Itoa(number)
			if len(numString)%2 == 0 {
				half := len(numString) / 2
				firstHalf := utils.MapToInt(numString[:half])
				secondHalf := utils.MapToInt(numString[half:])
				newNumbers[firstHalf] += count
				newNumbers[secondHalf] += count
				continue
			}
			newNumbers[number*2024] += count
		}

		numbersMap = newNumbers
	}

	var total int
	for _, count := range numbersMap {
		total += count
	}
	return total
}
