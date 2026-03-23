package ebitenPlus

import "strconv"

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

// Returns a string representation of p like "(3.14, 2.71)"
func (p Point) String() string {
	return "(" + strconv.FormatFloat(p.X, 'f', -1, 64) + "," + strconv.FormatFloat(p.Y, 'f', -1, 64) + ")"
}

// Returns the vector p+q.
func (p Point) Add(q Point) Point {
	return Pt(p.X+q.X, p.Y+q.Y)
}

// Returns the vector p-q.
func (p Point) Sub(q Point) Point {
	return Pt(p.X-q.X, p.Y-q.Y)
}

// Mul returns the vector p*k.
func (p Point) Mul(k float64) Point {
	return Pt(p.X*k, p.Y*k)
}

// Div returns the vector p/k.
func (p Point) Div(k float64) Point {
	return Pt(p.X/k, p.Y/k)
}

// Returns true if p and q are equal.
func (p Point) Eq(q Point) bool {
	return p == q
}
