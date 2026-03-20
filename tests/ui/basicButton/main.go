package main

import (
	"fmt"

	"github.com/VladiTNT/ebitenPlus/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	B       *ui.LabeledButton
	Bb      *ui.BasicButton
	Counter int
}

func NewGame() *Game {
	return &Game{
		B:       ui.NewLabeledButton(ui.NewBasicButton(20, 80, 100, 40), "Vlad"),
		Bb:      ui.NewBasicButton(20, 160, 100, 40),
		Counter: 0,
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	g.B.Update()
	g.Bb.Update()

	g.B.OnClick(ebiten.MouseButton0, func() {
		g.Counter++
	})

	g.Bb.WhenHeld(ebiten.MouseButton0, func() {
		g.Counter++
	})

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Counter: %d", g.Counter), 0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("B: Hover- %t Held- %t", g.B.IsHover, g.B.IsHeld), 0, 12)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Bb: Hover- %t Held- %t", g.Bb.IsHover, g.Bb.IsHeld), 0, 24)

	g.B.Draw(screen)

	g.B.DrawBounds(screen)
	g.Bb.DrawBounds(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowTitle("Ui test")
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(NewGame()); err != nil {
		fmt.Println(err)
	}
}
