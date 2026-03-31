package asset

import "github.com/hajimehoshi/ebiten/v2"

// This method loads the image located in path from the disk with the given key.
// The image will also be returned after loading and be available with the GetImage method.
func (a *AssetMan) LoadImage(path, key string) *ebiten.Image {
	// We reserve space in the map.
	i := NewImage()
	a.Images[key] = i

	// We load the image into the map and also return it.
	return a.loadImage(i, path)
}

// This method does the same thing as LoadImage() but asyncronously.
//
// Keep in mind that the image that is loading may not be available right away, if GetImage is
// called on an image that is currently loading, it will block until the worker is done.
func (a *AssetMan) AsyncLoadImage(path, key string) {
	// We reserve space in the map.
	i := NewImage()
	a.Images[key] = i

	// We let the image load in the background.
	go a.loadImage(i, path)
}

// This method loads an image from the FS of the AssetMan, if the FS is nil, the function will panic.
// The loaded asset will be returned and will be available with the GetImage method.
func (a *AssetMan) LoadImageFS(path, key string) *ebiten.Image {
	i := NewImage()
	a.Images[key] = i

	return a.loadImageFS(i, path)
}

// This method does the same thing as LoadImageFS but does it asyncronously.
//
// Keep in mind that the image that is loading may not be available right away, if GetImage is
// called on an image that is currently loading, it will block until the worker is done.
func (a *AssetMan) AsyncLoadImageFS(path, key string) {
	// We reserve space in the map.
	i := NewImage()
	a.Images[key] = i

	// We let the image load in the background.
	go a.loadImageFS(i, path)
}

// This method fetches the image with the given key and it blocks if the image is still loading from
// an async function.
func (a *AssetMan) GetImage(key string) *ebiten.Image {
	// Retrive image from map.
	i, ok := a.Images[key]
	if !ok {
		a.Logger.Printf("Asset '%s' was not found. (Called by GetImage)\n", key)
	}

	// Blocks until asset is available.
	<-i.C

	return i.Img
}

// This method frees the image's internal state and removes it's entry from the map.
func (a *AssetMan) FreeImage(key string) {
	// Retrive image from map.
	i, ok := a.Images[key]
	if !ok {
		a.Logger.Printf("Asset '%s' was not found. (Called by FreeImage)\n", key)
	}

	// Free it.
	i.Img.Deallocate()
	delete(a.Images, key)
}
