package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// LabeledButton is just a BasicButton with a label.
type LabeledButton struct {
	*BasicButton
	Label        string
	ColWhenHover color.Color
	ColWhenHeld  color.Color
}

// Makes a new LabeledButton.
// The colors for the "hover" and "when held" actions are Gray96_ALPHA96 and Gray224 by default.
func NewLabeledButton(btn *BasicButton, label string) *LabeledButton {
	return &LabeledButton{
		BasicButton:  btn,
		Label:        label,
		ColWhenHover: Gray96_ALPHA96,
		ColWhenHeld:  Gray224,
	}
}

// Draw's the button label in the middle of the button's bounds.
func (b *LabeledButton) Draw(screen *ebiten.Image) {
	if b.IsHover {
		vector.FillRect(screen, float32(b.X), float32(b.Y), float32(b.W), float32(b.H),
			b.ColWhenHover, false)
	}

	if b.IsHeld || b.IsClicked {
		vector.FillRect(screen, float32(b.X), float32(b.Y), float32(b.W), float32(b.H),
			b.ColWhenHeld, false)
	}

	ebitenutil.DebugPrintAt(screen, b.Label, b.X+b.W/2-len(b.Label)*6/2, b.Y+b.H/2-6)
}
