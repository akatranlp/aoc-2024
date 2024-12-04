package map2d

type Map2D struct {
	data []string
	rows int
	cols int
}

type Point struct {
	X, Y int
}

func NewMap2D(data []string) *Map2D {
	return &Map2D{
		data: data,
		rows: len(data),
		cols: len(data[0]),
	}
}
