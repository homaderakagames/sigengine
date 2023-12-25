package resources

import (
	fontManager "github.com/homaderakagames/sigengine/pkg/engine/resources/fonts"
)

type ResourceManager struct {
	FontManager *fontManager.FontManager
}

func NewResourceManager(
	fontManager *fontManager.FontManager,
) *ResourceManager {

	return &ResourceManager{
		FontManager: fontManager,
	}
}
