package fonts

import (
	"fmt"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
)

type FontSettings struct {
	Name    string
	Size    float64
	DPI     float64
	Hinting font.Hinting
}

type FontManager struct {
	prepared map[*FontSettings][]byte
	loaded   map[*FontSettings]font.Face
}

func NewFontManager() *FontManager {
	return &FontManager{
		prepared: make(map[*FontSettings][]byte),
		loaded:   make(map[*FontSettings]font.Face),
	}
}

func (r *FontManager) Set(key *FontSettings, b []byte) {
	if r.prepared == nil {
		r.prepared = make(map[*FontSettings][]byte)
	}

	r.prepared[key] = b
}

func (r *FontManager) Get(key *FontSettings) (font.Face, bool) {
	if r.loaded == nil {
		r.loaded = make(map[*FontSettings]font.Face)
	}

	val, ok := r.loaded[key]
	return val, ok
}

func (r *FontManager) Load() {
	if r.prepared == nil {
		panic(fmt.Sprintf("attempted to load font manager without prepared fonts"))
	}

	for fontKey, fontBytes := range r.prepared {
		tt, err := opentype.Parse(fontBytes)
		if err != nil {
			log.Printf("load failed for %s: %v", fontKey.Name, err)
		}

		fontFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    fontKey.Size,
			DPI:     fontKey.DPI,
			Hinting: fontKey.Hinting,
		})
		if err != nil {
			log.Printf("load failed to %s: %v", fontKey.Name, err)
		}
		r.loaded[fontKey] = fontFace
	}
}
