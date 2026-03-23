package graphics

import (
	"github.com/VladiTNT/ebitenPlus"
	"github.com/hajimehoshi/ebiten/v2"
)

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

// Makes the camera's world position follow ahead of the point p.
// This function also takes in the point's velocity to compute how far ahead it should be.
//
// maxDist represents how far ahead the camera will try to be, a value of 10-20 is usually good for most
// games.
func (c *Camera) FollowAhead(p, vel ebitenPlus.Point, maxDist, speed float64) {
	targetX := p.X + (vel.X * maxDist)
	targetY := p.Y + (vel.Y * maxDist)

	c.Wx += (targetX - c.Wx) * speed
	c.Wy += (targetY - c.Wy) * speed
}

// Makes the camera's world position follow the point p with an offset based of the cursor position.
//
// maxDist is the radius around p after which the camera won't follow the cursor anymore.
func (c *Camera) FollowCursor(p ebitenPlus.Point, maxDist, speed float64) {
	x, y := ebiten.CursorPosition()

	targetX := p.X + float64(x) - c.Cx
	targetY := p.Y + float64(y) - c.Cy

	offX := targetX - p.X
	offY := targetY - p.Y

	if offX > maxDist {
		offX = maxDist
	}
	if offX < -maxDist {
		offX = -maxDist
	}
	if offY > maxDist {
		offY = maxDist
	}
	if offY < -maxDist {
		offY = -maxDist
	}

	c.Wx += (p.X + offX - c.Wx) * speed
	c.Wy += (p.Y + offY - c.Wy) * speed
}
