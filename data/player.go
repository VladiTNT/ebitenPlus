package data

import "github.com/hajimehoshi/ebiten/v2/audio"

// Returns the sample rate of the audio context.
func (a *AssetManager) AudioSampleRate() int {
	return a.audioContext.SampleRate()
}

// Makes a new audio player.
func (a *AssetManager) NewAudioPlayer(bytes []byte) *audio.Player {
	return a.audioContext.NewPlayerFromBytes(bytes)
}

// Wrapper around the audio context's "IsReady()" method.
func (a *AssetManager) IsAudioReady() bool {
	return a.audioContext.IsReady()
}

// Makes a new F32 audio player.
func (a *AssetManager) NewF32AudioPlayer(bytes []byte) *audio.Player {
	return a.audioContext.NewPlayerF32FromBytes(bytes)
}
