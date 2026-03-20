package ebitenPlus

// A Point is an X, Y float64 coordinate pair. The axes increase right and down.
type Point struct {
	X, Y float64
}

// Pt is shorthand for Point{X, Y}.
func Pt(x, y float64) Point {
	return Point{X: x, Y: y}
}

// Coords returns the X and Y values of the point.
func (p Point) Coords() (float64, float64) {
	return p.X, p.Y
}
