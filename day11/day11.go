package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"aoc-lib/utils"
	"fmt"
	"io"
	"iter"
	"maps"
	"slices"
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

func test() {
	digits := make(map[int]map[int]int)
	// digits[0] = make(map[int]int)
	digits[1] = make(map[int]int)
	digits[2] = make(map[int]int)
	digits[3] = make(map[int]int)
	digits[4] = make(map[int]int)
	digits[5] = make(map[int]int)
	digits[6] = make(map[int]int)
	digits[7] = make(map[int]int)
	digits[8] = make(map[int]int)
	digits[9] = make(map[int]int)

	digits[0][0] = 0 // 0 -> 1
	digits[0][1] = 0 // digits[1][0]

	digits[1][0] = 1 // 1
	digits[1][1] = 1 // 2024
	digits[1][2] = 2 // 20 24
	digits[1][3] = 4 // digits[2][0] digits[0][0] digits[2][0] digits[4][0]

	digits[2][0] = 1 // 2
	digits[2][1] = 1 // 4048
	digits[2][2] = 1 // 40 48
	digits[2][3] = 1 // digits[4][0] digits[0][0] digits[4][0] digits[8][0]

	digits[3][0] = 1 // 3
	digits[3][1] = 1 // 6072
	digits[3][2] = 1 // 60 72
	digits[3][3] = 1 // 6 0 7 2

	digits[4][0] = 1 // 4
	digits[4][1] = 1 // 8096
	digits[4][2] = 1 // 80 96
	digits[4][3] = 1 // digits[8][0] digits[0][0] digits[9][0] digits[6][0]

	digits[5][0] = 1 // 5
	digits[5][1] = 1 // 10120
	digits[5][2] = 1 // 20.482.880
	digits[5][3] = 1 // 2048 2880
	digits[5][4] = 1 // 20 48 28 80
	digits[5][5] = 1 // digits[2][0] digits[0][0] digits[4][0] digits[8][0] digits[2][0] digits[8][0] digits[8][0] digits[0][0]

	digits[6][0] = 1 // 6
	digits[6][1] = 1 // 12144
	digits[6][2] = 1 // 12144
	digits[6][3] = 1 // 24.579.456
	digits[6][4] = 1 // 2457 9456
	digits[6][5] = 1 // 24 57 94 56
	digits[6][6] = 1 // digits[2][0] digits[4][0] digits[5][0] digits[7][0] digits[9][0] digits[4][0] digits[5][0] digits[6][0]

	digits[7][0] = 1 // 7
	digits[7][1] = 1 // 14168
	digits[7][2] = 1 // 28.676.032
	digits[7][2] = 1 // 2867 6032

	digits[8][0] = 1 // 8
	digits[8][1] = 1 // 16162
	digits[8][2] = 1 // 32.772.608
	digits[8][3] = 1 // 3277 2608
	digits[8][4] = 1 // 32 77 26 digits[8][0]
	digits[8][5] = 1 // digits[3][0] digits[2][0] digits[7][0] digits[7][0] digits[2][0] digits[6][0] digits[8][1]

	digits[9][0] = 1 // 9
	digits[9][1] = 1 // 18216
	digits[9][2] = 1 // 36.869.184
	digits[9][3] = 1 // 3686 9184

	// zeroes[0] = 1 // 0

}

type DigitMap map[int]map[int]Calculation

func (dm DigitMap) DebugPrint() {
	for digit := range slices.Sorted(maps.Keys(dm)) {
		fmt.Printf("%d: ", digit)
		for stage := range slices.Sorted(maps.Keys(dm[digit])) {
			calc := dm[digit][stage]
			fmt.Print(calc, ", ")
		}
		fmt.Println()
	}
}

func NewDigitsMap() DigitMap {
	dm := make(DigitMap)
	dm[0] = make(map[int]Calculation)
	dm[0][0] = Calculation{numbers: []int{0}, count: 1}
	dm[1] = make(map[int]Calculation)
	dm[1][0] = Calculation{numbers: []int{1}, count: 1}
	dm[2] = make(map[int]Calculation)
	dm[2][0] = Calculation{numbers: []int{2}, count: 1}
	dm[3] = make(map[int]Calculation)
	dm[3][0] = Calculation{numbers: []int{3}, count: 1}
	dm[4] = make(map[int]Calculation)
	dm[4][0] = Calculation{numbers: []int{4}, count: 1}
	dm[5] = make(map[int]Calculation)
	dm[5][0] = Calculation{numbers: []int{5}, count: 1}
	dm[6] = make(map[int]Calculation)
	dm[6][0] = Calculation{numbers: []int{6}, count: 1}
	dm[7] = make(map[int]Calculation)
	dm[7][0] = Calculation{numbers: []int{7}, count: 1}
	dm[8] = make(map[int]Calculation)
	dm[8][0] = Calculation{numbers: []int{8}, count: 1}
	dm[9] = make(map[int]Calculation)
	dm[9][0] = Calculation{numbers: []int{9}, count: 1}
	return dm
}

type Calculation struct {
	numbers []int
	count   int
	// if num -1 then its in digts and stages
	digits []int
	stages []int
}

func calcNextStage(digit, stage int, dm DigitMap) {
	v := dm[digit][stage]
	if last, ok := dm[digit][stage-1]; ok {
		last.digits = nil
		last.numbers = nil
		last.stages = nil
		dm[digit][stage-1] = last
	}

	newCalc := Calculation{}
	if len(v.numbers) == 1 {
		if v.numbers[0] == 0 {
			newCalc.digits = []int{1}
			newCalc.stages = []int{0}
		} else {
			numString := strconv.Itoa(v.numbers[0])
			if len(numString)%2 == 0 {
				firstHalf := utils.MapToInt(numString[:len(numString)/2])
				secondHalf := utils.MapToInt(numString[len(numString)/2:])
				newCalc.numbers = []int{firstHalf, secondHalf}
				newCalc.count = 2
			} else {
				newCalc.numbers = []int{v.numbers[0] * 2024}
				newCalc.count = 1
			}
		}
	} else if len(v.numbers) > 1 {
		newCalc.numbers = make([]int, 0, len(v.numbers)*2)

		var numA, numB string
		for _, n := range v.numbers {
			numString := strconv.Itoa(n)
			numA = numString[:len(numString)/2]
			numB = numString[len(numString)/2:]
			firstHalf := utils.MapToInt(numA)
			secondHalf := utils.MapToInt(numB)
			numB = strconv.Itoa(secondHalf)

			if len(numA) == 1 {
				newCalc.digits = append(newCalc.digits, firstHalf, secondHalf)
			} else if len(numB) == 1 {
				newCalc.numbers = append(newCalc.numbers, firstHalf)
				newCalc.digits = append(newCalc.digits, secondHalf)
			} else {
				newCalc.numbers = append(newCalc.numbers, firstHalf, secondHalf)
			}
		}

		newCalc.stages = make([]int, len(newCalc.digits))
		newCalc.digits = append(newCalc.digits, v.digits...)
		newCalc.stages = append(newCalc.stages, its.MapSlice(v.stages, MapInc)...)
		newCalc.count = len(newCalc.numbers)
	} else {
		newCalc.digits = v.digits
		newCalc.stages = its.MapSlice(v.stages, MapInc)
	}

	for i := range len(newCalc.digits) {
		digit := newCalc.digits[i]
		stage := newCalc.stages[i]
		newCalc.count += CalcStages(digit, stage, dm)
	}
	dm[digit][stage+1] = newCalc
}

func MapInc(n int) int { return n + 1 }

func CalcStages(digit, stage int, dm DigitMap) int {
	v, ok := dm[digit][stage]
	if ok {
		fmt.Println(v)
		return v.count
	}
	fmt.Println("Calc next stage for", digit, stage)
	calcNextStage(digit, stage-1, dm)
	return dm[digit][stage].count
}

func (*Day11) Part1(r io.Reader) int {
	numbers := its.MapSlice(strings.Fields(string(utils.Must(io.ReadAll(r)))), utils.MapToInt)
	dm := NewDigitsMap()

	zeroes := make([]int, 0)
	stages := make([]int, 0)
	var count int

	for range 25 {
		newSlice := make([]int, 0, len(numbers)*2)

		for _, v := range numbers {
			if v == 0 {
				zeroes = append(zeroes, 0)
				stages = append(stages, 0)
				continue
			}
			numString := strconv.Itoa(v)
			if len(numString)%2 == 0 {
				firstHalf := utils.MapToInt(numString[:len(numString)/2])
				secondHalf := utils.MapToInt(numString[len(numString)/2:])
				newSlice = append(newSlice, firstHalf)
				newSlice = append(newSlice, secondHalf)
				continue
			}
			newSlice = append(newSlice, v*2024)
		}

		numbers = newSlice
		count = 0
		for i := range len(zeroes) {
			count += CalcStages(zeroes[i], stages[i], dm)
		}
	}

	dm.DebugPrint()
	return len(numbers) + count
}

func (*Day11) Part2(r io.Reader) int {
	numbers := its.MapSlice(strings.Fields(string(utils.Must(io.ReadAll(r)))), utils.MapToInt)
	fmt.Println(numbers)
	// dll := NewDLLFromValues(numbers...)

	for range 75 {
		newSlice := make([]int, 0, len(numbers)*2)

		for _, v := range numbers {
			if v == 0 {
				newSlice = append(newSlice, 1)
				continue
			}
			numString := strconv.Itoa(v)
			if len(numString)%2 == 0 {
				firstHalf := utils.MapToInt(numString[:len(numString)/2])
				secondHalf := utils.MapToInt(numString[len(numString)/2:])
				newSlice = append(newSlice, firstHalf)
				newSlice = append(newSlice, secondHalf)
				continue
			}
			newSlice = append(newSlice, v*2024)
		}

		numbers = newSlice

		// for node := range dll.Iter() {
		// 	if node.value == 0 {
		// 		node.value = 1
		// 		continue
		// 	}
		// 	numString := strconv.Itoa(node.value)
		// 	if len(numString)%2 == 0 {
		// 		firstHalf := utils.MapToInt(numString[:len(numString)/2])
		// 		secondHalf := utils.MapToInt(numString[len(numString)/2:])
		// 		newNode := &Node{value: firstHalf}
		// 		node.value = secondHalf
		// 		dll.InsertBefore(newNode, node)
		// 		continue
		// 	}
		// 	node.value *= 2024
		// }
	}

	return len(numbers)
}
