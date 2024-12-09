package map2d

type Vector2 struct{ X, Y int }

func NewVector2(x, y int) Vector2 {
	return Vector2{x, y}
}

// sin90 := int(math.Sin(math.Pi / 2))
// cos90 := int(math.Cos(math.Pi / 2))
const (
	sin0   = 0
	cos0   = 1
	sin90  = 1
	cos90  = 0
	sin180 = 0
	cos180 = -1
	sin270 = -1
	cos270 = 0
)

func (v Vector2) RotateClockwise() Vector2 {
	return Vector2{
		X: v.X*cos90 - v.Y*sin90,
		Y: v.X*sin90 + v.Y*cos90,
	}
}

func (v *Vector2) RotateClockwiseMut() *Vector2 {
	x := v.X
	y := v.Y
	v.X = x*cos90 - y*sin90
	v.Y = x*sin90 + y*cos90
	return v
}

func (v Vector2) RotateCounterClockwise() Vector2 {
	return Vector2{
		X: v.X*cos270 - v.Y*sin270,
		Y: v.X*sin270 + v.Y*cos270,
	}
}

func (v *Vector2) RotateCounterClockwiseMut() *Vector2 {
	x := v.X
	y := v.Y
	v.X = x*cos270 - y*sin270
	v.Y = x*sin270 + y*cos270
	return v
}

func (v Vector2) RotateHalf() Vector2 {
	return Vector2{
		X: v.X*cos180 - v.Y*sin180,
		Y: v.X*sin180 + v.Y*cos180,
	}
}

func (v *Vector2) RotateHalfMut() *Vector2 {
	x := v.X
	y := v.Y
	v.X = x*cos180 - y*sin180
	v.Y = x*sin180 + y*cos180
	return v
}

func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{v.X - other.X, v.Y - other.Y}
}

func (v *Vector2) SubMut(other Vector2) *Vector2 {
	v.X -= other.X
	v.Y -= other.Y
	return v
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{v.X + other.X, v.Y + other.Y}
}

func (v *Vector2) AddMut(other Vector2) *Vector2 {
	v.X += other.X
	v.Y += other.Y
	return v
}

func (v Vector2) Scale(n int) Vector2 {
	return Vector2{v.X * n, v.Y * n}
}

func (v *Vector2) ScaleMut(n int) *Vector2 {
	v.X *= n
	v.Y *= n
	return v
}
