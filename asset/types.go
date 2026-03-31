package asset

import "github.com/hajimehoshi/ebiten/v2"

// An Image is a wrapper around an ebiten.Image with a channel indicating if the image is ready to use.
//
// It's mostly ment for internal use, I made it public just in case anyone wanted to mess with it.
type Image struct {
	Img *ebiten.Image
	C   chan struct{}
}

// Initializes an Image with an empty asset, meant to reserve space in the map for an image while
// it's loading asynchronously.
func NewImage() *Image {
	return &Image{
		Img: nil,
		C:   make(chan struct{}),
	}
}

func (i *Image) FinishedLoading(img *ebiten.Image) {
	i.Img = img
	close(i.C)
}

// AduioData is a wrapper around the raw byte stream of an audio with a channel indicating if the
// audio is ready to use.
//
// It's mostly ment for internal use, I made it public just in case anyone wanted to mess with it.
type AudioData struct {
	Data []byte
	C    chan struct{}
}

// Initializes an AudioData with no asset, meant to reserve space in the map for an image while it's
// loading asynchronously.
func NewAudioData() *AudioData {
	return &AudioData{
		Data: nil,
		C:    make(chan struct{}),
	}
}

func (a *AudioData) FinishedLoading(data []byte) {
	a.Data = data
	close(a.C)
}
