package main

import (
	"ebitenPlus/data"
	"ebitenPlus/tests/data/assets"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Width, Height int

	A  *audio.Player
	As *data.AssetManager
}

func NewGame() *Game {
	return &Game{
		Width: 320, Height: 240,
		As: data.NewAssetManager(44100, assets.Assets),
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		err := g.A.Rewind()
		if err != nil {
			return err
		}

		g.A.Play()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.Width)/2-32, float64(g.Height)/2-32)

	screen.DrawImage(g.As.GetImageFromFS("sprite.png"), op)
}

func (g *Game) Layout(w, h int) (int, int) {
	return g.Width, g.Height
}

func main() {
	g := NewGame()
	g.A = g.As.NewAudioPlayer(g.As.GetWavFromFS("jump.wav"))

	ebiten.SetWindowTitle("Asset manager test")
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(g); err != nil {
		fmt.Println("Game crash:", err)
		return
	}
}
