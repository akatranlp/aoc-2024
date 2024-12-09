package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	set "aoc-lib/slices"
	"aoc-lib/utils"
	"fmt"
	"io"
	"slices"
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

func (u *Update) isValid(orders []Order) bool {
	for _, order := range orders {
		_, okL := u.nMap[order.l]
		_, okR := u.nMap[order.r]
		if !okL || !okR {
			continue
		}
		for _, n := range u.nList {
			if n == order.l {
				break
			} else if n == order.r {
				return false
			}
		}
	}
	return true
}

func (*Day5) Part1(r io.Reader) int {
	orders := make([]Order, 0)
	updates := make([]*Update, 0)

	first := true
	for row := range its.ReaderToIter(r) {
		if row == "" {
			first = false
			continue
		}
		if first {
			var order Order
			utils.Must(fmt.Sscanf(row, "%d|%d", &order.l, &order.r))
			orders = append(orders, order)
		} else {
			numbers := its.MapSlice(strings.Split(row, ","), utils.MapToInt)
			update := &Update{
				nList: numbers,
				nMap:  set.NewSetWithValues(numbers...),
			}
			updates = append(updates, update)
		}
	}

	orderMap := make(map[int][]int)
	for _, order := range orders {
		orderMap[order.r] = append(orderMap[order.r], order.l)
	}

	return its.Reduce(slices.Values(updates), 0, func(acc int, u *Update) int {
		if u.isValid(orders) {
			return acc + u.nList[len(u.nList)/2]
		}
		return acc
	})

}

func (*Day5) Part2(r io.Reader) int {
	orders := make([]Order, 0)
	updates := make([]*Update, 0)
	first := true
	for row := range its.ReaderToIter(r) {
		if row == "" {
			first = false
			continue
		}
		if first {
			var order Order
			utils.Must(fmt.Sscanf(row, "%d|%d", &order.l, &order.r))
			orders = append(orders, order)
		} else {
			numbers := its.MapSlice(strings.Split(row, ","), utils.MapToInt)
			update := &Update{
				nList: numbers,
				nMap:  set.NewSetWithValues(numbers...),
			}
			updates = append(updates, update)
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

	falseUpdates := its.Filter(slices.Values(updates), func(u *Update) bool { return !u.isValid(orders) })

	type tuple struct {
		f, s int
	}

	return its.Reduce(falseUpdates, 0, func(acc int, update *Update) int {
		order := its.MapSlice(update.nList, func(n int) tuple {
			orderSet := orderMap[n]
			numbersBefore := orderSet.Intersect(update.nMap)
			return tuple{n, len(numbersBefore)}
		})
		slices.SortFunc(order, func(a tuple, b tuple) int { return a.s - b.s })
		return acc + order[len(order)/2].f
	})
}
