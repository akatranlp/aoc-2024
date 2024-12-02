package main

import (
	"aoc-lib/fs"
	"aoc-lib/its"
	"aoc-lib/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const (
	ASC = iota
	DESC
	UNKNOWN
)

func main() {
	reports := make([][]int, 0)
	fs.ApplyToLines("day2/input.txt", func(msg string) {
		values := strings.Split(msg, " ")
		fmt.Println(values)
		seq := slices.Values(values)
		numbers := its.Map(seq, func(s string) int {
			return utils.Must(strconv.Atoi(s))
		})
		reports = append(reports, slices.Collect(numbers))
	})

	sum := 0
	for _, list := range reports {
		worked := true
		strict := UNKNOWN
		var value1, value2 int
		for i := range list {
			j := i + 1
			if j == len(list) {
				break
			}
			value1, value2 = list[i], list[j]
			asc := value1-value2 < 0
			if strict == UNKNOWN {
				if asc {
					strict = ASC
				} else {
					strict = DESC
				}
			} else {
				if asc && strict == DESC {
					worked = false
					break
				} else if !asc && strict == ASC {
					worked = false
					break
				}
			}

			difference := utils.IntAbs(value1, value2)
			if difference < 1 || difference > 3 {
				worked = false
			}
		}
		if worked {
			sum += 1
		}
	}
	fmt.Println("Sum1:", sum)
}
