package graphics

import "image/color"

// DrawOptions are attributes about images that can be changed to alter their appearence when drawn.
type DrawOptions struct {
	// Theta is the rotation angle of the sprite.
	Theta float64
	// ScaleX and ScaleY stretch the sprite.
	ScaleX, ScaleY float64
	// A, R, G and B are the premultiplied ColorScale values.
	A, R, G, B float32
	// Mipmaps help increase quality of shrinked sprites.
	Mipmaps bool
}

// Gives a DrawOptions struct with:
//
// Theta = 0
//
// ScaleX, ScaleY = 1 (no scaling)
//
// A, R, G, B = 1 (base colors)
//
// Mipmaps = false
func NewDrawOptions() DrawOptions {
	return DrawOptions{
		Theta:  0,
		ScaleX: 1, ScaleY: 1,
		A: 1, R: 1, G: 1, B: 1,
		Mipmaps: false,
	}
}

// Sets both ScaleX and ScaleY to the given float value.
func (d *DrawOptions) Scale(size float64) {
	d.ScaleX = size
	d.ScaleY = size
}

// Scales the ColorScale by the given color.
func (d *DrawOptions) ScaleColor(c color.Color) {
	r, g, b, a := c.RGBA()

	fR := float32(r)
	fG := float32(g)
	fB := float32(b)
	fA := float32(a)

	d.R, d.G, d.B, d.A = fR, fG, fB, fA
}
