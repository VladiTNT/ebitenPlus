package ui

import (
	"image/color"

	"github.com/VladiTNT/ebitenplus"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Button is just an interface with the "OnClick" and "WhenHeld" methods.
type Button interface {
	OnClick(ebiten.MouseButton, func())
	WhenHeld(ebiten.MouseButton, func())
}

// A BasicButton is just a rectangle with input checking. It is mainly a base for labeled
// and sprite buttons.
//
// It is compatible with both mouse and touch.
//
// The StrokeWidth and Col fields are for the debug "DrawBounds" function.
type BasicButton struct {
	X, Y int
	W, H int

	IsHeld  bool
	IsHover bool

	heldTouchIds        []ebiten.TouchID
	justPressedTouchIds []ebiten.TouchID

	StrokeWidth float32
	Col         color.Color
}

// Makes a BasicButton with the specified size. StrokeWidth is 1 and Col is
// color.White by default.
func NewBasicButton(x, y, w, h int) *BasicButton {
	return &BasicButton{
		X: x, Y: y, W: w, H: h,
		heldTouchIds:        make([]ebiten.TouchID, 0),
		justPressedTouchIds: make([]ebiten.TouchID, 0),
		StrokeWidth:         1,
		Col:                 color.White,
	}
}

// Draws the bounds of the button.
func (b *BasicButton) DrawBounds(screen *ebiten.Image) {
	vector.StrokeRect(screen, float32(b.X), float32(b.Y), float32(b.W), float32(b.H), b.StrokeWidth, b.Col, false)
}

// Updates the touchIds and IsHover, this function must be executed before the "OnClick" and
// "WhenHeld" methods, otherwise your input checking is going to be one tick behind.
func (b *BasicButton) Update() {
	x, y := ebiten.CursorPosition()
	b.heldTouchIds = ebiten.AppendTouchIDs(b.heldTouchIds[:0])
	b.justPressedTouchIds = inpututil.AppendJustPressedTouchIDs(b.justPressedTouchIds[:0])

	b.IsHover = ebitenplus.AABB(x, y, b.X, b.X+b.W, b.Y, b.Y+b.H)
}

// Detects single taps and executes the given function.
func (b *BasicButton) OnClick(btn ebiten.MouseButton, do func()) {
	touched := false
	for _, id := range b.justPressedTouchIds {
		x, y := ebiten.TouchPosition(id)
		if ebitenplus.AABB(x, y, b.X, b.X+b.W, b.Y, b.Y+b.H) {
			touched = true
			break
		}
	}

	if (b.IsHover && inpututil.IsMouseButtonJustPressed(btn)) || touched {
		do()
	}
}

// Detects constant presses and executes the given function every frame.
func (b *BasicButton) WhenHeld(btn ebiten.MouseButton, do func()) {
	touched := false
	for _, id := range b.heldTouchIds {
		x, y := ebiten.TouchPosition(id)
		if ebitenplus.AABB(x, y, b.X, b.X+b.W, b.Y, b.Y+b.H) {
			touched = true
			break
		}
	}

	if (b.IsHover && ebiten.IsMouseButtonPressed(btn)) || touched {
		b.IsHeld = true
		do()
	} else {
		b.IsHeld = false
	}
}
