package asset

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/fs"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// Decodes the image from r using image.Decode().
func decodeIntoEbitenImage(r io.Reader) (*ebiten.Image, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

func LoadImageFromFile(path string) (*ebiten.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	i, err := decodeIntoEbitenImage(f)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func LoadImageFromFS(FS fs.FS, pattern string) (*ebiten.Image, error) {
	f, err := FS.Open(pattern)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	i, err := decodeIntoEbitenImage(f)
	if err != nil {
		return nil, err
	}

	return i, nil
}
