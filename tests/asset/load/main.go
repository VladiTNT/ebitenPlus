package main

import (
	"fmt"
	"log"
	"os"

	"github.com/VladiTNT/ebitenPlus/asset"
	testassets "github.com/VladiTNT/ebitenPlus/test_assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	PLAYER_SPRITE = "PLAYER"
	PLANK_SPRITE  = "PLANK"
)

type Game struct {
	As *asset.AssetMan
}

func NewGame() *Game {
	l := log.New(os.Stderr, "ERROR: ", log.LstdFlags)

	g := &Game{
		As: asset.NewAssetMan(44100, l),
	}

	g.As.MountFileSystem(testassets.AssetFS)

	// Loading PNG from file on disk on main thread.
	g.As.LoadImage("./test_assets/sprite.png", PLAYER_SPRITE)
	// Loading JPEG from FS on worker thread.
	g.As.AsyncLoadImageFS("plank.jpeg", PLANK_SPRITE)

	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Player sprite
	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Translate(40, 40)
	screen.DrawImage(g.As.GetImage(PLAYER_SPRITE), op1)

	// Plank sprite
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(80, 40)
	screen.DrawImage(g.As.GetImage(PLANK_SPRITE), op2)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %f", ebiten.ActualTPS()))
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowTitle("Asset_load_TEST")
	ebiten.SetWindowSize(1280, 960)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
