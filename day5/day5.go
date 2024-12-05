package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/utils"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Day5 struct{}

var _ aoc.Problem = (*Day5)(nil)

type Order struct {
	l, r int
}

type Update struct {
	nList []int
	nMap  map[int]struct{}
}

func (*Day5) Part1(r io.Reader) int {
	orders := make([]Order, 0)
	updates := make([]*Update, 0)
	for i, block := range its.ReaderToIter(r, its.SplitByBlocks) {
		if i == 0 {
			for _, line := range its.Filter2(its.ReaderToIter(bytes.NewBufferString(block)), its.FilterEmptyLines) {
				var order Order
				utils.Must(fmt.Sscanf(line, "%d|%d", &order.l, &order.r))
				orders = append(orders, order)
			}
		} else if i == 1 {
			for _, line := range its.Filter2(its.ReaderToIter(bytes.NewBufferString(block)), its.FilterEmptyLines) {
				numbers := its.MapSlice(strings.Split(line, ","), func(s string) int { return utils.Must(strconv.Atoi(s)) })
				update := &Update{
					nList: numbers,
					nMap:  make(map[int]struct{}),
				}
				for _, n := range numbers {
					update.nMap[n] = struct{}{}
				}
				updates = append(updates, update)
			}
		}
	}

	orderMap := make(map[int][]int)
	for _, order := range orders {
		orderMap[order.r] = append(orderMap[order.r], order.l)
	}

	var sum int
	for _, update := range updates {
		length := len(update.nList)
		middleInt := update.nList[length/2]
		found := true

	outer:
		for _, order := range orders {
			_, okL := update.nMap[order.l]
			_, okR := update.nMap[order.r]
			if !okL || !okR {
				continue
			}
			for _, n := range update.nList {
				if n == order.l {
					break
				} else if n == order.r {
					found = false
					break outer
				}
			}
		}

		if found {
			sum += middleInt
		}
	}

	return sum
}

func (*Day5) Part2(r io.Reader) int {
	fmt.Println("Part2 not implemented")
	return -1
}
