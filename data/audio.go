package data

import (
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

// Loads MP3 data from a static file if not cached.
func (a *AssetManager) GetMP3FromFile(path string) []byte {
	if b, ok := a.audios[path]; ok {
		return b
	}

	file, err := os.Open(path)
	if err != nil {
		a.errors <- err
		return nil
	}

	d, err := mp3.DecodeWithSampleRate(a.audioContext.SampleRate(), file)
	if err != nil {
		a.errors <- err
		return nil
	}

	r, err := io.ReadAll(d)
	if err != nil {
		a.errors <- err
		return nil
	}

	a.audios[path] = r

	return r
}

// Loads Wav data from a static file if not cached.
func (a *AssetManager) GetWavFromFile(path string) []byte {
	if b, ok := a.audios[path]; ok {
		return b
	}

	file, err := os.Open(path)
	if err != nil {
		a.errors <- err
		return nil
	}

	d, err := wav.DecodeWithSampleRate(a.audioContext.SampleRate(), file)
	if err != nil {
		a.errors <- err
		return nil
	}

	r, err := io.ReadAll(d)
	if err != nil {
		a.errors <- err
		return nil
	}

	a.audios[path] = r

	return r
}

// Loads Vorbis data from a static file if not cached.
func (a *AssetManager) GetVorbisFromFile(path string) []byte {
	if b, ok := a.audios[path]; ok {
		return b
	}

	file, err := os.Open(path)
	if err != nil {
		a.errors <- err
		return nil
	}

	d, err := vorbis.DecodeWithSampleRate(a.audioContext.SampleRate(), file)
	if err != nil {
		a.errors <- err
		return nil
	}

	r, err := io.ReadAll(d)
	if err != nil {
		a.errors <- err
		return nil
	}

	a.audios[path] = r

	return r
}

// Loads MP3 data from the manager's file system if not cached.
func (a *AssetManager) GetMP3FromFS(path string) []byte {
	if b, ok := a.audios[path]; ok {
		return b
	}

	file, err := a.fileSystem.Open(path)
	if err != nil {
		a.errors <- err
		return nil
	}

	d, err := mp3.DecodeWithSampleRate(a.audioContext.SampleRate(), file)
	if err != nil {
		a.errors <- err
		return nil
	}

	r, err := io.ReadAll(d)
	if err != nil {
		a.errors <- err
		return nil
	}

	a.audios[path] = r

	return r
}

// Loads Wav data from the manager's file system if not cached.
func (a *AssetManager) GetWavFromFS(path string) []byte {
	if b, ok := a.audios[path]; ok {
		return b
	}

	file, err := a.fileSystem.Open(path)
	if err != nil {
		a.errors <- err
		return nil
	}

	d, err := wav.DecodeWithSampleRate(a.audioContext.SampleRate(), file)
	if err != nil {
		a.errors <- err
		return nil
	}

	r, err := io.ReadAll(d)
	if err != nil {
		a.errors <- err
		return nil
	}

	a.audios[path] = r

	return r
}

// Loads Vorbis data from the manager's file system if not cached.
func (a *AssetManager) GetVorbisFromFS(path string) []byte {
	if b, ok := a.audios[path]; ok {
		return b
	}

	file, err := a.fileSystem.Open(path)
	if err != nil {
		a.errors <- err
		return nil
	}

	d, err := vorbis.DecodeWithSampleRate(a.audioContext.SampleRate(), file)
	if err != nil {
		a.errors <- err
		return nil
	}

	r, err := io.ReadAll(d)
	if err != nil {
		a.errors <- err
		return nil
	}

	a.audios[path] = r

	return r
}
