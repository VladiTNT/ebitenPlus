package data_test

import (
	"ebitenPlus/data"
	"testing"
)

func TestManager(t *testing.T) {
	a := data.NewAssetManager(44100, nil)

	ok := a.IsAudioReady()

	if !ok {
		t.Errorf("Error audio context is not ok.\n")
	}

	i1 := a.LoadImageFromFile("/img")
	i2 := a.LoadImageFromFS("/assets/img")

	a1 := a.LoadMP3FromFile("/mp3")
	a2 := a.LoadMP3FromFS("/assets/mp3")

	if i1 != nil || i2 != nil || a1 != nil || a2 != nil {
		t.Errorf("Assets are not nil when they should be: %v %v %v %v\n", i1, i2, a1, a2)
	}

	errs := a.Errors()

	for _, err := range errs {
		t.Errorf("Asset manager error: %v\n", err)
	}
}
