package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
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

func (d *Day4) checkX(row, col int) bool {
	{
		newRow1 := row - 1
		newCol1 := col - 1
		newRow2 := row + 1
		newCol2 := col + 1
		if !d.inBounce(newRow1, newCol1) || !d.inBounce(newRow2, newCol2) {
			return false
		}
		char1 := d.wordMap[newRow1][newCol1]
		char2 := d.wordMap[newRow2][newCol2]
		if !(char1 == 'M' && char2 == 'S') && !(char1 == 'S' && char2 == 'M') {
			return false
		}
	}
	{
		newRow1 := row + 1
		newCol1 := col - 1
		newRow2 := row - 1
		newCol2 := col + 1
		if !d.inBounce(newRow1, newCol1) || !d.inBounce(newRow2, newCol2) {
			return false
		}
		char1 := d.wordMap[newRow1][newCol1]
		char2 := d.wordMap[newRow2][newCol2]
		if !(char1 == 'M' && char2 == 'S') && !(char1 == 'S' && char2 == 'M') {
			return false
		}
	}
	return true
}

func (d *Day4) Part2(r io.Reader) int {
	d.wordMap = make([]string, 0)
	for _, row := range its.Filter2(its.ReaderToIter(r), its.FilterEmptyLines) {
		d.wordMap = append(d.wordMap, row)
	}
	var count int

	for i := range d.wordMap {
		for j := range d.wordMap[i] {
			element := d.wordMap[i][j]
			if element == 'A' {
				if d.checkX(i, j) {
					count++
				}
			}
		}
	}

	return count
}
