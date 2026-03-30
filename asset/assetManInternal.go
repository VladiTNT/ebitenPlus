package asset

import (
	"errors"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

// This function loads an Image into the map and returns it.
func (a *AssetMan) loadImage(i *Image, path, key string) *ebiten.Image {
	var p *ebiten.Image
	var err error

	// If the AssetMan has a file system we try to load it from there.
	if a.FileSystem != nil {
		p, err = LoadImageFromFS(a.FileSystem, path)
		// If the asset isn't there than we try to load it from a file.
		if errors.Is(err, fs.ErrNotExist) {
			a.Logger.Printf("Couldn't find '%s' in FS, fetching it from disk...\n", path)
			p, err = LoadImageFromFile(path)
		}
	} else {
		// If the AssetMan has no file system we load it from a file.
		p, err = LoadImageFromFile(path)
	}

	// We check if we encountered an error.
	if err != nil {
		a.Logger.Println(err)
	}

	// We insert the image into the map and mark it as ready for use.
	i.FinishedLoading(p)

	return p
}
