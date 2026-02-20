package data

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

/*
Returns an array of ebiten images from a sprite sheet:

n - number of sprites to parse;

spWidth, spHeight - size of each individual sprite;

sheetlength - horizontal length of the sheet measured in sprite widths;
*/
func LoadSpriteSheet(img *ebiten.Image, n int, spWidth, spHeight int, sheetlength int) []*ebiten.Image {
	sprites := make([]*ebiten.Image, n)
	for i := range n {
		col := i % sheetlength
		row := i / sheetlength

		x0 := col * spWidth
		y0 := row * spHeight
		x1 := x0 + spWidth
		y1 := y0 + spHeight
		sprites[i] = img.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
	}
	return sprites
}
