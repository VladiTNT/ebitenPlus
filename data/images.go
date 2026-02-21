package data

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (a *AssetManager) LoadImageFromFile(path string) *ebiten.Image {
	if i, ok := a.images[path]; ok {
		return i
	}

	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		a.errors <- err
		return nil
	}

	a.images[path] = img

	return img
}

func (a *AssetManager) LoadImageFromFS(path string) *ebiten.Image {
	if i, ok := a.images[path]; ok {
		return i
	}

	img, _, err := ebitenutil.NewImageFromFileSystem(a.fileSystem, path)
	if err != nil {
		a.errors <- err
		return nil
	}

	a.images[path] = img

	return img
}
