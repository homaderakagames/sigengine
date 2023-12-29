package input

import "github.com/hajimehoshi/ebiten/v2"

type KeyboardAdapter interface {
	IsKeyPressed(key ebiten.Key) bool
}

type EbitenKeyboardAdapter struct {
}

func (e *EbitenKeyboardAdapter) IsKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

type testKeyboardAdapter struct {
	curKey *ebiten.Key
}

func (t *testKeyboardAdapter) pushKey(key ebiten.Key) {
	t.curKey = &key
}

func (t *testKeyboardAdapter) releaseKey() {
	t.curKey = nil
}

func (t *testKeyboardAdapter) IsKeyPressed(key ebiten.Key) bool {
	if t.curKey == nil {
		return false
	}
	if *t.curKey == key {
		return true
	}
	return false
}
