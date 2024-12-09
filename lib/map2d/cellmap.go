package map2d

import (
	"aoc-lib/its"
	"fmt"
	"io"
	"iter"
)

type CellMap[T any] struct {
	data [][]T
	rows int
	cols int
}

func NewCellMap[T any](r io.Reader, mapFn func(x, y int, value byte) T) *CellMap[T] {
	map2d := new(CellMap[T])
	map2d.data = make([][]T, 0)
	for i, row := range its.Enumerate(its.Filter(its.ReaderToIter(r), its.FilterEmptyLines)) {
		rowBytes := []byte(row)
		rowSlice := make([]T, len(rowBytes))
		for j, char := range rowBytes {
			rowSlice[j] = mapFn(j, i, char)
		}
		map2d.data = append(map2d.data, rowSlice)

	}
	map2d.rows = len(map2d.data)
	map2d.cols = len(map2d.data[0])
	return map2d
}

func CellMapFn(x, y int, value byte) Cell { return Cell{X: x, Y: y, Value: value} }

func (d *CellMap[T]) InBounce(v Vector2) bool {
	return v.Y >= 0 && v.Y < d.rows && v.X >= 0 && v.X < d.cols
}

func (d *CellMap[T]) Get(v Vector2) T {
	return d.data[v.Y][v.X]
}

func (m *CellMap[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, row := range m.data {
			for _, v := range row {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func (d *CellMap[T]) DebugPrint(f ...func(T) string) {
	for _, row := range d.data {
		for _, cell := range row {
			if len(f) > 0 {
				fmt.Print(f[0](cell))
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Println()
	}
}

type Cell struct {
	X, Y  int
	Value byte
}

func (p Cell) ExtractCoords() Vector2 {
	return Vector2{p.X, p.Y}
}
