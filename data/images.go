package data

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Loads an ebiten.Image from a static file if not cached.
func (a *AssetManager) GetImageFromFile(path string) *ebiten.Image {
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

// Loads an ebiten.Image from the manager's file system if not cached.
func (a *AssetManager) GetImageFromFS(path string) *ebiten.Image {
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
