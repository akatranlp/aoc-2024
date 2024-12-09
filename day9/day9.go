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
	for _, row := range its.Filter2(its.ReaderToIter(r), its.FilterEmptyLines) {
		numbers = its.MapSlice([]byte(row), func(char byte) int { return int(char - '0') })
	}

	var disk = make([]int, 0)
	var id int
	for chunk := range slices.Chunk(numbers, 2) {
		length := len(chunk)
		if length == 0 {
			break
		}
		if length > 0 {
			file := chunk[0]

			for range file {
				disk = append(disk, id)
			}
		}
		if length > 1 {
			free := chunk[1]
			for range free {
				disk = append(disk, -1)
			}
		}
		id++
	}

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
	for _, row := range its.Filter2(its.ReaderToIter(r), its.FilterEmptyLines) {
		numbers = its.MapSlice([]byte(row), func(char byte) int { return int(char - '0') })
	}

	var disk = make([]*Block, 0)
	var id int
	for chunk := range slices.Chunk(numbers, 2) {
		length := len(chunk)
		if length == 0 {
			break
		}
		block := &Block{id: id}
		if length > 0 {
			file := chunk[0]
			block.fileSize = file
		}
		if length > 1 {
			free := chunk[1]
			block.freeSize = free
		}
		disk = append(disk, block)
		id++
	}

	for lastFileId := id - 1; lastFileId != 0; lastFileId-- {
		idx := slices.IndexFunc(disk, func(block *Block) bool { return block.id == lastFileId })
		tryToMoveBlock := disk[idx]

		for i := 0; i < idx; i++ {
			leftBlock := disk[i]
			if leftBlock.freeSize < tryToMoveBlock.fileSize {
				continue
			}
			leftFromRemovingBlock := disk[idx-1]
			leftFromRemovingBlock.freeSize += tryToMoveBlock.fileSize + tryToMoveBlock.freeSize
			tryToMoveBlock.freeSize = leftBlock.freeSize - tryToMoveBlock.fileSize
			leftBlock.freeSize = 0
			newDisk := make([]*Block, 0)
			newDisk = append(newDisk, disk[:i+1]...)
			newDisk = append(newDisk, tryToMoveBlock)
			newDisk = append(newDisk, disk[i+1:idx]...)
			newDisk = append(newDisk, disk[idx+1:]...)
			disk = newDisk

			break
		}
	}

	flatDisk := make([]int, 0)
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
