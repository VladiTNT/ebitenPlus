package ebitenPlus

// AABB returns true if the point defined by X, Y is present in the rectangle defined by x0, x1
// y0 and y1.
func AABB(X, Y int, x0, x1, y0, y1 int) bool {
	return X >= x0 && X <= x1 && Y >= y0 && Y <= y1
}
