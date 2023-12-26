package images

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png" // тк сейчас используются png
)

type ImageSettings struct {
	Name string
}

type ImageManager struct {
	prepared map[*ImageSettings][]byte
	loaded   map[*ImageSettings]*ebiten.Image
}

func NewImageManager() *ImageManager {
	return &ImageManager{
		prepared: make(map[*ImageSettings][]byte),
		loaded:   make(map[*ImageSettings]*ebiten.Image),
	}
}

func (r *ImageManager) Set(key *ImageSettings, b []byte) {
	if r.prepared == nil {
		r.prepared = make(map[*ImageSettings][]byte)
	}

	r.prepared[key] = b
}

func (r *ImageManager) Get(key *ImageSettings) (*ebiten.Image, bool) {
	if r.loaded == nil {
		r.loaded = make(map[*ImageSettings]*ebiten.Image)
	}

	val, ok := r.loaded[key]
	return val, ok
}

func (r *ImageManager) Load() {
	if r.prepared == nil {
		panic(fmt.Sprintf("attempted to load image manager without prepared images"))
	}

	for imageKey, imageBytes := range r.prepared {
		ebitenImage, _, err := ebitenutil.NewImageFromReader(bytes.NewBuffer(imageBytes))
		if err != nil {
			panic(fmt.Sprintf("failed to load image by key=%s: %v", imageKey, err))
		}
		r.loaded[imageKey] = ebitenImage
	}
}
