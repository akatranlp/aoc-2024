package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"fmt"
	"io"
	"slices"
)

type Day9 struct{}

var _ aoc.Problem = (*Day9)(nil)

func (*Day9) Part1(r io.Reader) int {
	var numbers []int
	for row := range its.Filter(its.ReaderToIter(r), its.FilterEmptyLines) {
		numbers = its.MapSlice([]byte(row), func(char byte) int { return int(char - '0') })
	}

	disk := its.Reduce2(its.Enumerate(slices.Chunk(numbers, 2)), make([]int, 0), func(acc []int, idx int, chunk []int) []int {
		length := len(chunk)
		if length == 0 {
			return acc
		}
		if length > 0 {
			file := chunk[0]

			for range file {
				acc = append(acc, idx)
			}
		}
		if length > 1 {
			free := chunk[1]
			for range free {
				acc = append(acc, -1)
			}
		}
		return acc
	})

	leftIdx, rightIdx := 0, len(disk)-1
	for leftIdx != rightIdx {
		leftValue := disk[leftIdx]
		rightValue := disk[rightIdx]

		for leftValue != -1 && leftIdx != rightIdx {
			leftIdx++
			leftValue = disk[leftIdx]
		}
		for rightValue == -1 && leftIdx != rightIdx {
			rightIdx--
			rightValue = disk[rightIdx]
		}

		if leftIdx == rightIdx {
			break
		}

		disk[leftIdx], disk[rightIdx] = disk[rightIdx], -1
		leftIdx++
		rightIdx--
	}

	var sum int
	for i, fileBlock := range disk {
		if fileBlock == -1 {
			break
		}
		sum += i * fileBlock
	}

	return sum
}

type Block struct {
	id, fileSize, freeSize int
}

func (*Day9) Part2(r io.Reader) int {
	var numbers []int
	for row := range its.Filter(its.ReaderToIter(r), its.FilterEmptyLines) {
		numbers = its.MapSlice([]byte(row), func(char byte) int { return int(char - '0') })
	}

	disk := its.Reduce2(its.Enumerate(slices.Chunk(numbers, 2)), make([]*Block, 0), func(acc []*Block, idx int, chunk []int) []*Block {
		length := len(chunk)
		if length == 0 {
			return acc
		}
		block := &Block{id: idx}
		if length > 0 {
			file := chunk[0]
			block.fileSize = file
		}
		if length > 1 {
			free := chunk[1]
			block.freeSize = free
		}
		return append(acc, block)
	})

	for idx := len(disk) - 1; idx != 0; {
		tryToMoveBlock := disk[idx]

		var found bool
		for i := 0; i < idx; i++ {
			leftBlock := disk[i]
			if leftBlock.freeSize < tryToMoveBlock.fileSize {
				continue
			}
			leftFromRemovingBlock := disk[idx-1]
			leftFromRemovingBlock.freeSize += tryToMoveBlock.fileSize + tryToMoveBlock.freeSize
			tryToMoveBlock.freeSize = leftBlock.freeSize - tryToMoveBlock.fileSize
			leftBlock.freeSize = 0
			disk = its.RemoveIndex(slices.Insert(disk, i+1, tryToMoveBlock), idx+1)
			found = true

			break
		}
		if !found {
			idx--
		}
	}

	capacity := its.Reduce(slices.Values(disk), 0, func(acc int, block *Block) int {
		return acc + block.fileSize + block.freeSize
	})
	flatDisk := make([]int, 0, capacity)
	for _, block := range disk {
		for range block.fileSize {
			flatDisk = append(flatDisk, block.id)
		}
		for range block.freeSize {
			flatDisk = append(flatDisk, -1)
		}
	}

	var sum int
	for i, fileBlock := range flatDisk {
		if fileBlock == -1 {
			continue
		}
		sum += i * fileBlock
	}

	return sum
}

func PrintDisk(disk []*Block) {
	fmt.Print("[")
	for _, block := range disk {
		fmt.Printf("%v ", block)
	}
	fmt.Println("]")
}
