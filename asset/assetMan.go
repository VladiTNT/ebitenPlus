package asset

import (
	"io/fs"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

// AssetMan is an asset manager. It loads assets and caches them for long term use, making sure that you
// don't load assets more than once.
//
// AssetMan can load assets from fs.FS file systems, keep in mind that if you give the asset manager
// one, it will always check it first when you call a LoadX method, so you should either keep all your
// assets in an fs.FS or on disk.
//
// NOTE: AssetMan tries to keep a simple API, in order to do this most of it's methods don't return
// errors, instead, errors are logged using whatever logger the user wants to give it. Keep in mind
// that some errors (such as failing to load an asset) will cause the runtime to panic shortly
// after occuring because the program will likely try to access that resource even if it has a nil value.
//
// Compatible data:
//
// - Images: PNG, JPEG
//
// - Audio: mp3, wav, vorbis (.ogg)
//
// Data retrieval options:
//
// - Static files
//
// - Virtual filesystems (io.FS)
type AssetMan struct {
	// Images is the map of cached images.
	Images map[string]*Image
	// Audios is the map of cached audio data.
	Audios       map[string]*AudioData
	AudioContext *audio.Context

	FileSystem fs.FS

	// Logger is used to log any errors that occur when assets are loaded.
	// Everything is logged with the Println method.
	Logger *log.Logger
}

// Initializes an empty AssetMan with an AudioContext with the given sample rate.
//
// Sample rate should always be 44100 or 48000.
//
// NOTE: FileSystem is set to nil by default, be careful how you access this value to avoid
// runtime panics. Use the MountFileSystem method to mount an FS.
func NewAssetMan(sampleRate int, logger *log.Logger) *AssetMan {
	return &AssetMan{
		Images:       make(map[string]*Image),
		Audios:       make(map[string]*AudioData),
		AudioContext: audio.NewContext(sampleRate),
		FileSystem:   nil,
		Logger:       logger,
	}
}

// This function sets the FileSystem of the AssetMan to the given one.
func (a *AssetMan) MountFileSystem(fs fs.FS) {
	a.FileSystem = fs
}
