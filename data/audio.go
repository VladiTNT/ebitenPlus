package data

import (
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

func (a *AssetManager) LoadMP3(path string) []byte {
	b, ok := a.Audios[path]

	if ok {
		return b
	}

	file, err := os.Open(path)
	if err != nil {
		a.Err = err
		return nil
	}

	d, err := mp3.DecodeWithSampleRate(a.AudioContext.SampleRate(), file)
	if err != nil {
		a.Err = err
		return nil
	}

	r, err := io.ReadAll(d)
	if err != nil {
		a.Err = err
		return nil
	}

	a.Audios[path] = r

	return r
}

func (a *AssetManager) LoadWav(path string) []byte {
	b, ok := a.Audios[path]

	if ok {
		return b
	}

	file, err := os.Open(path)
	if err != nil {
		a.Err = err
		return nil
	}

	d, err := wav.DecodeWithSampleRate(a.AudioContext.SampleRate(), file)
	if err != nil {
		a.Err = err
		return nil
	}

	r, err := io.ReadAll(d)
	if err != nil {
		a.Err = err
		return nil
	}

	a.Audios[path] = r

	return r
}

func (a *AssetManager) LoadOgg(path string) []byte {
	b, ok := a.Audios[path]

	if ok {
		return b
	}

	file, err := os.Open(path)
	if err != nil {
		a.Err = err
		return nil
	}

	d, err := vorbis.DecodeWithSampleRate(a.AudioContext.SampleRate(), file)
	if err != nil {
		a.Err = err
		return nil
	}

	r, err := io.ReadAll(d)
	if err != nil {
		a.Err = err
		return nil
	}

	a.Audios[path] = r

	return r
}
