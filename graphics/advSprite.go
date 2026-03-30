package graphics

import (
	"github.com/VladiTNT/ebitenPlus"
	"github.com/hajimehoshi/ebiten/v2"
)

// An advanced sprite is just a static image with a position vector and DrawOptions proprieties.
type AdvancedSprite struct {
	Img *ebiten.Image
	Pos ebitenPlus.Point
	DrawOptions
}

// Gives an AdvancedSprite with the given image and position.
func NewAdvancedSprite(img *ebiten.Image, pos ebitenPlus.Point) *AdvancedSprite {
	return &AdvancedSprite{
		Img:         img,
		Pos:         pos,
		DrawOptions: NewDrawOptions(),
	}
}

// Draws the sprite. ebiten.DrawImageOptions{} is configured with the attributes from DrawOptions
// embedded struct.
func (s *AdvancedSprite) Draw(screen *ebiten.Image, p ebitenPlus.Point) {
	op := &ebiten.DrawImageOptions{}

	// Get bounds
	w := s.Img.Bounds().Dx()
	h := s.Img.Bounds().Dy()

	// Apply rotation.
	op.GeoM.Translate(float64(-w/2), float64(-h/2))
	op.GeoM.Rotate(s.Theta)
	op.GeoM.Translate(float64(w/2), float64(h/2))

	// Scale
	op.GeoM.Scale(s.ScaleX, s.ScaleY)

	// ColorScale
	op.ColorScale.Scale(s.R, s.G, s.B, s.A)

	// Mipmaps
	op.DisableMipmaps = !s.Mipmaps

	// Movement Translation
	op.GeoM.Translate(p.X, p.Y)

	// Final draw
	screen.DrawImage(s.Img, op)
}
