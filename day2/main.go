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

type Order int

const (
	UNKNOWN Order = iota
	ASC
	DESC
)

func main() {
	reports := make([][]int, 0)
	fs.ApplyToLines("day2/input.txt", func(msg string) {
		values := strings.Split(msg, " ")
		seq := slices.Values(values)
		numbers := its.Map(seq, func(s string) int {
			return utils.Must(strconv.Atoi(s))
		})
		reports = append(reports, slices.Collect(numbers))
	})

	sum := 0
	for _, list := range reports {
		someWorked := false
		for removeIndex := range list {
			worked := true
			strict := UNKNOWN
			newList := make([]int, len(list))
			copy(newList, list)
			newList = its.RemoveIndex(newList, removeIndex)
			for i := range newList {
				j := i + 1
				if j == len(newList) {
					break
				}
				value1, value2 := newList[i], newList[j]
				asc1 := value1-value2 < 0
				if strict == UNKNOWN {
					if asc1 {
						strict = ASC
					} else {
						strict = DESC
					}
				} else {
					if asc1 && strict == DESC {
						worked = false
						break
					} else if !asc1 && strict == ASC {
						worked = false
						break
					}
				}

				difference := utils.IntAbs(value1, value2)
				if difference < 1 || difference > 3 {
					worked = false
					break
				}
			}
			if worked {
				someWorked = true
				break
			}
		}
		if someWorked {
			sum += 1
		}
	}
	fmt.Println("Sum1:", sum)
}
