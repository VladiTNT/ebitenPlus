package graphics

import "github.com/VladiTNT/ebitenPlus"

// A camera is an object that helps you with GeoM transformations for sprites.
type Camera struct {
	// These coordinates represent the center of the camera relative to the screen.
	Cx, Cy float64
	// These coordinates represent the position of the camera relative to the world.
	Wx, Wy float64
}

// Makes a camera with the specified coordinates.
func NewCamera(cx, cy, wx, wy float64) *Camera {
	return &Camera{Cx: cx, Cy: cy, Wx: wx, Wy: wy}
}

// Set's the camera's Wx and Wy coordinates to the specified point.
func (c *Camera) SetWorldCoords(x, y float64) {
	c.Wx = x
	c.Wy = y
}

// Set's the camera's Cx and Cy coordinates to the specified point.
func (c *Camera) SetScreenCoords(x, y float64) {
	c.Cx = x
	c.Cy = y
}

// Translates p relative to the camera's world position.
func (c *Camera) ApplyWorldTranslation(p ebitenPlus.Point) ebitenPlus.Point {
	return ebitenPlus.Pt(p.X+c.Cx-c.Wx, p.Y+c.Cy-c.Wy)
}

// Returns the translation for the camera's center.
// Use for sprites that stay glued to the screen's center (kinda like the player in Terraria).
func (c *Camera) ApplyCenterTranslation() ebitenPlus.Point {
	return ebitenPlus.Pt(c.Cx, c.Cy)
}

// Makes the camera's world position smoothly follow the point p.
// The implementation is simple using exponential smoothing.
//
// Slower speed's are smoother.
func (c *Camera) Follow(p ebitenPlus.Point, speed float64) {
	c.Wx += (p.X - c.Wx) * speed
	c.Wy += (p.Y - c.Wy) * speed
}
