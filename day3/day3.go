package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/utils"
	"fmt"
	"io"
	"regexp"
	"slices"
	"strconv"
)

type Day3 struct{}

var _ aoc.Problem = (*Day3)(nil)

func (*Day3) Part1(r io.Reader) int {
	matcher := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	data := string(utils.Must(io.ReadAll(r)))
	matches := matcher.FindAllStringSubmatch(data, -1)
	return its.Reduce(slices.Values(matches), 0, func(acc int, v []string) int {
		l, r := utils.Must(strconv.Atoi(v[1])), utils.Must(strconv.Atoi(v[2]))
		return acc + l*r
	})
}

func (*Day3) Part2(r io.Reader) int {
	matcher := regexp.MustCompile(`(?:mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`)
	data := string(utils.Must(io.ReadAll(r)))
	matches := matcher.FindAllStringSubmatch(data, -1)

	type result struct {
		enabled bool
		sum     int
	}

	res := its.Reduce(slices.Values(matches), &result{enabled: true}, func(acc *result, v []string) *result {
		action := v[0]
		if action == "do()" {
			acc.enabled = true
		} else if action == "don't()" {
			acc.enabled = false
		} else if acc.enabled {
			var l, r int
			fmt.Sscanf(action, "mul(%d,%d)", &l, &r)
			acc.sum += l * r
		}
		return acc
	})
	return res.sum
}
