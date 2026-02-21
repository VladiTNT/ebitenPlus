package data

import (
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

func (a *AssetManager) LoadMP3FromFile(path string) []byte {
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

func (a *AssetManager) LoadWavFromFile(path string) []byte {
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

func (a *AssetManager) LoadVorbisFromFile(path string) []byte {
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

func (a *AssetManager) LoadMP3FromFS(path string) []byte {
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

func (a *AssetManager) LoadWavFromFS(path string) []byte {
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

func (a *AssetManager) LoadVorbisFromFS(path string) []byte {
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
