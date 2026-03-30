package graphics

import (
	"github.com/VladiTNT/ebitenPlus"
	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite is just a static image with a position vector.
type Sprite struct {
	Img *ebiten.Image
	Pos ebitenPlus.Point
}

// Makes a new basic sprite.
func NewSprite(img *ebiten.Image, pos ebitenPlus.Point) *Sprite {
	return &Sprite{
		Img: img,
		Pos: pos,
	}
}

// Draw uses the transformation in p to draw the image at that position.
func (s *Sprite) Draw(screen *ebiten.Image, p ebitenPlus.Point) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(s.Img, op)
}
