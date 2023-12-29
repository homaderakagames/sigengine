package input

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

func TestInput_GetButtonState(t *testing.T) {
	testAdapter := &testKeyboardAdapter{}
	input := NewInput(testAdapter)

	var _, cur KeyState
	t.Run("not pressed button", func(t *testing.T) {
		cur = input.GetButtonState(ebiten.KeyW)
		if cur != KeyStateNotPressed {
			t.Fatal(fmt.Sprintf("key state must be not pressed"))
		}
	})

	t.Run("1st time pressed button", func(t *testing.T) {
		testAdapter.pushKey(ebiten.KeyW)

		cur = input.GetButtonState(ebiten.KeyW)
		if cur != KeyStatePressed {
			t.Fatal(fmt.Sprintf("key state must be pressed"))
		}
	})

	t.Run("hold button", func(t *testing.T) {
		cur = input.GetButtonState(ebiten.KeyW)
		if cur != KeyStateHold {
			t.Fatal(fmt.Sprintf("key state must be hold"))
		}

	})

	t.Run("still hold button", func(t *testing.T) {
		cur = input.GetButtonState(ebiten.KeyW)
		if cur != KeyStateHold {
			t.Fatal(fmt.Sprintf("key state must be hold"))
		}
	})

	t.Run("released button", func(t *testing.T) {
		testAdapter.releaseKey()

		cur = input.GetButtonState(ebiten.KeyW)
		if cur != KeyStateReleased {
			t.Fatal(fmt.Sprintf("key state must be released, not: %d", cur))
		}
	})

	t.Run("not pressed", func(t *testing.T) {
		cur = input.GetButtonState(ebiten.KeyW)
		if cur != KeyStateNotPressed {
			t.Fatal(fmt.Sprintf("key state must be not pressed, not: %d", cur))
		}
	})
}
