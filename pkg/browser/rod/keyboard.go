package rod

import (
	"testflowkit/pkg/browser"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

// keyMap maps abstract browser.Key to rod input.Key.
var keyMap = map[browser.Key]input.Key{
	browser.KeyEnter:      input.Enter,
	browser.KeyTab:        input.Tab,
	browser.KeyDelete:     input.Delete,
	browser.KeyEscape:     input.Escape,
	browser.KeySpace:      input.Space,
	browser.KeyArrowUp:    input.ArrowUp,
	browser.KeyArrowRight: input.ArrowRight,
	browser.KeyArrowDown:  input.ArrowDown,
	browser.KeyArrowLeft:  input.ArrowLeft,
}

type rodKeyboard struct {
	keyboard *rod.Keyboard
}

func (k *rodKeyboard) Press(key browser.Key) error {
	rodKey, ok := keyMap[key]
	if !ok {
		// If key not found in map, try to use it as a rune (for single character keys)
		if len(key) == 1 {
			rodKey = input.Key(key[0])
		} else {
			return &UnknownKeyError{Key: key}
		}
	}
	return k.keyboard.Press(rodKey)
}

func newRodKeyboard(keyboard *rod.Keyboard) browser.Keyboard {
	return &rodKeyboard{keyboard: keyboard}
}

// UnknownKeyError represents an error when an unknown key is pressed.
type UnknownKeyError struct {
	Key browser.Key
}

func (e *UnknownKeyError) Error() string {
	return "unknown key: " + string(e.Key)
}
