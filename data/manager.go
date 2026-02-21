package data

import (
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

// An AssetManager keeps track of your assets and makes sure you only load them once.
// AssetManager contains the audio context for a given application, thus you should never have
// more than one per game.
//
// AssetManager can also use data from io.fs file systems.
//
// Your audio and image data is cached in maps, so use the "Load..." methods as much as you want
// they won't cause new allocations.
//
// If anything goes wrong, errors will be sent into a channel and you can use the "Errors()"
// method to drain the errors in a slice.
type AssetManager struct {
	images       map[string]*ebiten.Image
	audios       map[string][]byte
	audioContext *audio.Context

	fileSystem fs.FS
	errors     chan error
}

// Makes a new AssetManager, if you are only gonna load data from static assets, you can just
// pass nil for the fileSystem.
//
// Sample rate should ideally be 44100 or 48000.
func NewAssetManager(sampleRate int, fileSystem fs.FS) *AssetManager {
	return &AssetManager{
		images:       make(map[string]*ebiten.Image),
		audios:       make(map[string][]byte),
		audioContext: audio.NewContext(sampleRate),

		fileSystem: fileSystem,
		errors:     make(chan error, 20),
	}
}

// Drains the errors channel into an array.
func (a *AssetManager) Errors() []error {
	var errs []error

	for {
		select {
		case e := <-a.errors:
			errs = append(errs, e)
		default:
			return errs
		}
	}
}
