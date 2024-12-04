package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"fmt"
	"io"
)

type Day4 struct {
	wordMap []string
}

var _ aoc.Problem = (*Day4)(nil)

func (d *Day4) searchWord(row, col int, word string) int {
	var count int
	neighbors := []int{-1, 0, 1}
	for _, dRow := range neighbors {
		for _, dCol := range neighbors {
			if dRow == 0 && dCol == 0 {
				continue
			}
			if d.searchWordDirection(row, col, dRow, dCol, word) {
				// fmt.Printf("FOUND XMAS at (%d %d) -> (%d, %d)\n", row, col, dCol, dRow)
				count++
			}
		}
	}
	return count
}

func (d *Day4) inBounce(row, col int) bool {
	rows := len(d.wordMap)
	cols := len(d.wordMap[0])
	return row >= 0 && row < rows && col >= 0 && col < cols
}

func (d *Day4) searchWordDirection(row, col, dRow, dCol int, word string) bool {
	for i := range word {
		newRow := row + dRow*i
		newCol := col + dCol*i
		if !d.inBounce(newRow, newCol) {
			return false
		}
		if d.wordMap[newRow][newCol] != word[i] {
			return false
		}
	}
	return true
}

func (d *Day4) Part1(r io.Reader) int {
	d.wordMap = make([]string, 0)
	for _, row := range its.Filter2(its.ReaderToIter(r), its.FilterEmptyLines) {
		d.wordMap = append(d.wordMap, row)
	}
	var count int

	// fmt.Println(d.searchWord(0, 5, "XMAS"))

	for i := range d.wordMap {
		for j := range d.wordMap[i] {
			element := d.wordMap[i][j]
			if element == 'X' {
				count += d.searchWord(i, j, "XMAS")
			}
		}
	}

	return count
}

func (*Day4) Part2(r io.Reader) int {
	fmt.Println("Part2 not implemented")
	return -1
}
