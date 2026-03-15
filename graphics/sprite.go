package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Basic sprite is just a static image with X and Y coordinates.
type BasicSprite struct {
	Img *ebiten.Image
	Pos Point
}

// Makes a new basic sprite.
func NewBasicSprite(img *ebiten.Image, x, y float64) *BasicSprite {
	return &BasicSprite{
		Img: img,
		Pos: Pt(x, y),
	}
}

// Draw uses the transformation in p to draw the image at that position.
func (s *BasicSprite) Draw(screen *ebiten.Image, p Point) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(s.Img, op)
}
