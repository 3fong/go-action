package lock

import (
	"image"
	"sync"
)

var loadIconsOnce sync.Once
var icons map[string]image.Image

// Concurrency-safe.
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}
