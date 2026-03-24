package main

import (
	"time"

	"github.com/VladiTNT/ebitenPlus/data"
	"github.com/VladiTNT/ebitenPlus/graphics"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	A *graphics.AnimationSystem

	X, Y float64
}

func NewGame() *Game {
	img, _, err := ebitenutil.NewImageFromFile("./tests/graphics/SpriteSheet.png")
	if err != nil {
		panic(err)
	}
	sps := data.LoadSpriteSheet(img, 13, 32, 32, 4)

	delay := time.Millisecond * 250

	a := graphics.NewAnimationSystem()
	a.Register("Idle", graphics.NewAnimation(sps[0:1], delay))
	a.Register("Right", graphics.NewAnimation(sps[1:4], delay))
	a.Register("Up", graphics.NewAnimation(sps[4:7], delay))
	a.Register("Left", graphics.NewAnimation(sps[7:10], delay))
	a.Register("Down", graphics.NewAnimation(sps[10:13], delay))

	return &Game{

		A: a,

		X: 0, Y: 0,
	}
}

func (g *Game) Update() error {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		g.A.Play("Up")
		g.Y -= 2
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		g.A.Play("Left")
		g.X -= 2
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		g.A.Play("Down")
		g.Y += 2
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		g.A.Play("Right")
		g.X += 2
	default:
		g.A.Play("Idle")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.X, g.Y)
	screen.DrawImage(g.A.Sprite(), op)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowTitle("Animation test")
	ebiten.SetWindowSize(960, 720)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
