package map2d

import (
	"aoc-lib/its"
	"fmt"
	"iter"
)

type Map2D struct {
	data [][]byte
	rows int
	cols int
}

func (d *Map2D) InBounce(v Vector2) bool {
	return v.Y >= 0 && v.Y < d.rows && v.X >= 0 && v.X < d.cols
}
func (d *Map2D) DebugPrint() {
	for _, row := range d.data {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func (d *Map2D) Set(p Point) {
	d.data[p.Y][p.X] = p.Value
}

type Vector2 struct{ X, Y int }

func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{v.X - other.X, v.Y - other.Y}
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{v.X + other.X, v.Y + other.Y}
}

type Point struct {
	X, Y  int
	Value byte
}

func (p Point) Extract() Vector2 {
	return Vector2{p.X, p.Y}
}

func NewMap2DFromStrings(data []string) *Map2D {
	return &Map2D{
		data: its.MapSlice(data, func(row string) []byte { return []byte(row) }),
		rows: len(data),
		cols: len(data[0]),
	}
}

func NewMap2D() *Map2D {
	return &Map2D{
		data: make([][]byte, 0),
		rows: 0,
		cols: 0,
	}
}

func (m *Map2D) AppendString(row string) *Map2D {
	m.data = append(m.data, []byte(row))
	m.cols = len(row)
	m.rows += 1
	return m
}

func (m *Map2D) Append(row []byte) *Map2D {
	m.data = append(m.data, row)
	m.cols = len(row)
	m.rows += 1
	return m
}

func (m *Map2D) IterEachField() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for y, row := range m.data {
			for x, v := range row {
				if !yield(Point{x, y, v}) {
					return
				}
			}
		}
	}
}
