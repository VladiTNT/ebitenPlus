package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// LabeledButton is just a BasicButton with a label.
type LabeledButton struct {
	*BasicButton
	Label string
}

// Makes a new LabeledButton
func NewLabeledButton(btn *BasicButton, label string) *LabeledButton {
	return &LabeledButton{BasicButton: btn, Label: label}
}

// Draw's the button label in the middle of the buttons bounds.
func (b *LabeledButton) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, b.Label, b.X+b.W/2-len(b.Label)*6/2, b.Y+b.H/2-6)
}
