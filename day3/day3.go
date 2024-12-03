package main

import (
	"aoc-lib/aoc"
	"aoc-lib/utils"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

type Day3 struct{}

var _ aoc.Problem = (*Day3)(nil)

func (*Day3) Part1(r io.Reader) int {
	matcher := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	data := string(utils.Must(io.ReadAll(r)))
	matches := matcher.FindAllStringSubmatch(data, -1)
	var sum int
	for _, match := range matches {
		l, r := utils.Must(strconv.Atoi(match[1])), utils.Must(strconv.Atoi(match[2]))
		sum += l * r
	}

	return sum
}

func (*Day3) Part2(r io.Reader) int {
	matcher := regexp.MustCompile(`(?:mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`)
	data := string(utils.Must(io.ReadAll(r)))
	matches := matcher.FindAllStringSubmatch(data, -1)
	var sum int
	enabled := true
	for _, match := range matches {
		action := match[0]
		if action == "do()" {
			enabled = true
		} else if action == "don't()" {
			enabled = false
		} else if enabled {
			var l, r int
			fmt.Sscanf(action, "mul(%d,%d)", &l, &r)
			sum += l * r
		}
	}

	return sum
}
