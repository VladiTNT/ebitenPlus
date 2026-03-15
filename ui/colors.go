package ui

import "image/color"

// Basic colors.
var (
	ColorRED          = color.RGBA{255, 0, 0, 255}
	ColorBLUE         = color.RGBA{0, 0, 255, 255}
	ColorGREEN        = color.RGBA{0, 255, 0, 255}
	ColorORANGE       = color.RGBA{255, 128, 0, 255}
	ColorYELLOW       = color.RGBA{255, 255, 0, 255}
	ColorLIME         = color.RGBA{128, 255, 0, 255}
	ColorSPRING_GREEN = color.RGBA{0, 255, 128, 255}
	ColorAQUA         = color.RGBA{0, 255, 255, 255}
	ColorAZURE        = color.RGBA{0, 128, 255, 255}
	ColorPURPLE       = color.RGBA{127, 0, 255, 255}
	ColorPINK         = color.RGBA{255, 0, 255, 255}
)

// Shades of gray with the specified RGB values.
// 224 = birghter and 64 = darker.
var (
	Gray224 = color.RGBA{224, 244, 244, 255}
	Gray192 = color.RGBA{192, 192, 192, 255}
	Gray160 = color.RGBA{160, 160, 160, 255}
	Gray128 = color.RGBA{128, 128, 128, 255}
	Gray96  = color.RGBA{96, 96, 96, 255}
	Gray64  = color.RGBA{64, 64, 64, 255}
)

// Some transparent gray colors.
var (
	Gray96_ALPHA96   = color.RGBA{96, 96, 96, 96}
	Gray128_ALPHA128 = color.RGBA{128, 128, 128, 128}
	Gray160_ALPHA160 = color.RGBA{160, 160, 160, 160}
)
