package main

import (
	"github.com/VladiTNT/ebitenplus/data"
	"github.com/VladiTNT/ebitenplus/graphics"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 240
	ScreenHeight = 180
)

type Game struct {
	As  *data.AssetManager
	Cam *graphics.Camera

	P *graphics.BasicSprite
	B *graphics.BasicSprite
}

func NewGame() *Game {
	g := new(Game)

	g.As = data.NewAssetManager(44100, nil)
	g.Cam = graphics.NewCamera(ScreenWidth/2, ScreenHeight/2, 15, 15)

	g.P = graphics.NewBasicSprite(g.As.GetImageFromFile("./tests/graphics/Sprite01.png"), 0, 0)
	g.B = graphics.NewBasicSprite(g.As.GetImageFromFile("./tests/graphics/Sprite02.png"), 100, 100)

	return g
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.P.Pos.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.P.Pos.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.P.Pos.Y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.P.Pos.X += 2
	}

	g.Cam.Follow(graphics.Pt(g.P.Pos.X+15, g.P.Pos.Y+15), 0.05)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.B.Draw(screen, g.Cam.ApplyWorldTranslation(g.B.Pos))

	g.P.Draw(screen, g.Cam.ApplyWorldTranslation(g.P.Pos))
}

func (g *Game) Layout(w, h int) (int, int) {
	return 240, 180
}

func main() {
	ebiten.SetWindowTitle("Graphics test")
	ebiten.SetWindowSize(480, 360)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
