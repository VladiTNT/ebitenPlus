package data

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type AssetManager struct {
	Images map[string]*ebiten.Image
	Audios map[string][]byte

	AudioContext *audio.Context
	Err          error
}

func NewAssetManager(sampleRate int) *AssetManager {
	return &AssetManager{
		Images: make(map[string]*ebiten.Image),
		Audios: make(map[string][]byte),

		AudioContext: audio.NewContext(sampleRate),
	}
}

func (a *AssetManager) LoadImage(path string) *ebiten.Image {
	i, ok := a.Images[path]

	if ok {
		return i
	}

	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		a.Err = err
		return nil
	}

	a.Images[path] = img

	return img
}
