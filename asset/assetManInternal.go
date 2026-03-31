package asset

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// This method loads an Image into the map from the disk.
func (a *AssetMan) loadImage(i *Image, path string) *ebiten.Image {
	p, err := LoadImageFromFile(path)
	if err != nil {
		a.Logger.Println(err)
	}

	i.FinishedLoading(p)

	return p
}

// This method loads an Image from the FS of AssetMan, it will also panic if the FS is nil.
func (a *AssetMan) loadImageFS(i *Image, path string) *ebiten.Image {
	if a.FileSystem == nil {
		panic("FileSystem is nil")
	}

	p, err := LoadImageFromFS(a.FileSystem, path)
	if err != nil {
		a.Logger.Println(err)
	}

	i.FinishedLoading(p)

	return p
}
