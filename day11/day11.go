package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/utils"
	"fmt"
	"io"
	"iter"
	"strconv"
	"strings"
)

type Day11 struct{}

var _ aoc.Problem = (*Day11)(nil)

type Node struct {
	next  *Node
	prev  *Node
	value int
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewDLLFromValues(values ...int) *DoubleLinkedList {
	ddl := &DoubleLinkedList{}
	for _, v := range values {
		ddl.Add(v)
	}
	return ddl
}

func (d *DoubleLinkedList) Add(value int) {
	node := &Node{value: value}
	d.size++
	if d.head == nil {
		d.head = node
		d.tail = node
		return
	}
	node.prev = d.tail
	d.tail.next = node
	d.tail = node
}

func (d *DoubleLinkedList) Iter() iter.Seq[*Node] {
	return func(yield func(*Node) bool) {
		curr := d.head
		for curr != nil {
			if !yield(curr) {
				return
			}
			curr = curr.next
		}
	}
}

func (d *DoubleLinkedList) InsertBefore(node *Node, after *Node) {
	d.size++
	if after == d.head {
		d.head = node
		node.next = after
		after.prev = node
		return
	}
	before := after.prev
	before.next = node
	node.prev = before
	node.next = after
	after.prev = node
}

func (*Day11) Part1(r io.Reader) int {
	numbers := its.MapSlice(strings.Fields(string(utils.Must(io.ReadAll(r)))), utils.MapToInt)
	fmt.Println(numbers)
	dll := NewDLLFromValues(numbers...)

	for range 25 {
		for node := range dll.Iter() {
			if node.value == 0 {
				node.value = 1
				continue
			}
			numString := strconv.Itoa(node.value)
			if len(numString)%2 == 0 {
				firstHalf := utils.MapToInt(numString[:len(numString)/2])
				secondHalf := utils.MapToInt(numString[len(numString)/2:])
				newNode := &Node{value: firstHalf}
				node.value = secondHalf
				dll.InsertBefore(newNode, node)
				continue
			}
			node.value *= 2024
		}
	}

	return dll.size
}

func (*Day11) Part2(r io.Reader) int {
	fmt.Println("Part2 not implemented")
	return -1
}
