package asset

import "github.com/hajimehoshi/ebiten/v2"

// This function loads an image from the file system of the AssetMan and caches it, if there is no
// file system it loads it from the disk, if the AssetMan has a file system but it can't find the image
// inside it, it tries to load it from the disk (it also logs that it's doing that). After it has finished
// loading it, it returns it and it will also be available in the future with the GetImage method.
func (a *AssetMan) LoadImage(path, key string) *ebiten.Image {
	// We reserve space in the map.
	i := NewImage()
	a.Images[key] = i

	// We load the image into the map and also return it.
	return a.loadImage(i, path, key)
}

// This function does the same thing as LoadImage, however this one returns to execution after
// it is called.
//
// Keep in mind that the asset that is being loaded may not be available right away, this isn't usually
// a big issue as the GetImage method waits for assets that are currently loading.
func (a *AssetMan) AsyncLoadImage(path, key string) {
	// We reserve space in the map.
	i := NewImage()
	a.Images[key] = i

	// We let the image load in the background.
	go a.loadImage(i, path, key)
}
