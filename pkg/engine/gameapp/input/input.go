package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type KeyState int

func (st *KeyState) nextKeyState() (prevState KeyState, curState KeyState) {
	switch *st {
	case KeyStateNotPressed:
		curState = KeyStatePressed
	case KeyStatePressed:
		curState = KeyStateHold
	case KeyStateHold:
		curState = KeyStateReleased
	case KeyStateReleased:
		curState = KeyStateNotPressed
	}
	prevState = *st
	return
}

const (
	KeyStateNotPressed KeyState = iota
	KeyStatePressed
	KeyStateHold
	KeyStateReleased
)

// Input не потокобезопасный драйвер ввода
type Input struct {
	keyboardAdapter KeyboardAdapter
	keyStates       map[ebiten.Key]KeyState
}

func NewInput(keyboardAdapter KeyboardAdapter) *Input {
	return &Input{
		keyboardAdapter: keyboardAdapter,
		keyStates:       map[ebiten.Key]KeyState{},
	}
}

// GetButtonState возвращает состояние кнопки с проверкой нажатия
func (in *Input) GetButtonState(key ebiten.Key) KeyState {
	curKeyState, ok := in.keyStates[key]
	if !in.keyboardAdapter.IsKeyPressed(key) {
		if !ok {
			return KeyStateNotPressed
		}

		// released
		_, next := curKeyState.nextKeyState()
		in.keyStates[key] = KeyStateReleased
		return next
	}

	if curKeyState == KeyStateNotPressed {
		in.keyStates[key] = KeyStatePressed
		return KeyStatePressed
	}

	if curKeyState == KeyStatePressed {
		_, next := curKeyState.nextKeyState()
		in.keyStates[key] = next
		return next
	}

	return KeyStateHold
}

// IsButtonPressed возвращает true если кнопка была нажата и не удерживается
func (in *Input) IsButtonPressed(key ebiten.Key) bool {
	return in.GetButtonState(key) == KeyStatePressed
}

// IsButtonHold возвращает true если кнопка удерживается
func (in *Input) IsButtonHold(key ebiten.Key) bool {
	return in.GetButtonState(key) == KeyStateHold
}

// IsButtonReleased возвращает true если кнопка отпущена
func (in *Input) IsButtonReleased(key ebiten.Key) bool {
	return in.GetButtonState(key) == KeyStateReleased
}

// IsButtonNotPressed возвращает true если кнопка не нажата, не удерживается и не отпущена
func (in *Input) IsButtonNotPressed(key ebiten.Key) bool {
	return in.GetButtonState(key) == KeyStateNotPressed
}
