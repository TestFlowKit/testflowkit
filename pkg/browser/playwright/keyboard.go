package playwright

import (
	"testflowkit/pkg/browser"

	pw "github.com/playwright-community/playwright-go"
)

var keyMap = map[browser.Key]string{
	browser.KeyEnter:      "Enter",
	browser.KeyTab:        "Tab",
	browser.KeyDelete:     "Delete",
	browser.KeyEscape:     "Escape",
	browser.KeySpace:      " ",
	browser.KeyArrowUp:    "ArrowUp",
	browser.KeyArrowRight: "ArrowRight",
	browser.KeyArrowDown:  "ArrowDown",
	browser.KeyArrowLeft:  "ArrowLeft",
}

type playwrightKeyboard struct {
	page pw.Page
}

func (k *playwrightKeyboard) Press(key browser.Key) error {
	playwrightKey, ok := keyMap[key]
	if ok {
		return k.page.Keyboard().Press(playwrightKey)
	}

	if len(key) == 1 {
		return k.page.Keyboard().Press(string(key))
	}

	return &UnknownKeyError{Key: key}
}

func newPlaywrightKeyboard(page pw.Page) browser.Keyboard {
	return &playwrightKeyboard{page: page}
}

// UnknownKeyError represents an error when an unknown key is pressed.
type UnknownKeyError struct {
	Key browser.Key
}

func (e *UnknownKeyError) Error() string {
	return "unknown key: " + string(e.Key)
}
