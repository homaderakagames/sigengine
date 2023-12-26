package resources

import (
	fontManager "github.com/homaderakagames/sigengine/pkg/engine/resources/fonts"
	imageManager "github.com/homaderakagames/sigengine/pkg/engine/resources/images"
)

type ResourceManager struct {
	FontManager  *fontManager.FontManager
	ImageManager *imageManager.ImageManager
}

func NewResourceManager(
	fontManager *fontManager.FontManager,
	imageManager *imageManager.ImageManager,
) *ResourceManager {

	return &ResourceManager{
		FontManager:  fontManager,
		ImageManager: imageManager,
	}
}

func (rm *ResourceManager) Load() {
	rm.ImageManager.Load()
	rm.FontManager.Load()
}
