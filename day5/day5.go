package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	set "aoc-lib/slices"
	"aoc-lib/utils"
	"bytes"
	"fmt"
	"io"
	"maps"
	"slices"
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
	nMap  set.Set[int]
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
					nMap:  set.NewSet[int](),
				}
				for _, n := range numbers {
					update.nMap.Set(n)
				}
				updates = append(updates, update)
			}
		}
	}

	orderMap := make(map[int]set.Set[int])
	for _, order := range orders {
		m, ok := orderMap[order.r]
		if !ok {
			m = set.NewSet[int]()
		}
		m.Set(order.l)
		orderMap[order.r] = m
	}

	var falseUpdates []*Update

	var sum int
	for _, update := range updates {
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

		if !found {
			falseUpdates = append(falseUpdates, update)
		}
	}

	type tuple struct {
		f, s int
	}

	for _, update := range falseUpdates {
		numbersBeforeMap := make(map[int][]int)
		for _, n := range update.nList {
			orderSet := orderMap[n]
			numbersBefore := orderSet.Intersect(update.nMap)
			numbersBeforeMap[n] = slices.Collect(maps.Keys(numbersBefore))
		}

		order := make([]tuple, 0)
		for k, v := range numbersBeforeMap {
			order = append(order, tuple{k, len(v)})
		}
		slices.SortFunc(order, func(a tuple, b tuple) int { return a.s - b.s })
		final := its.MapSlice(order, func(o tuple) int { return o.f })
		sum += final[len(final)/2]
	}

	return sum
}
