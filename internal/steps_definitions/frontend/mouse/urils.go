package mouse

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/browser/common"
	"time"
)

func getInteractableElement(page common.Page, label string) (common.Element, error) {
	element, err := browser.GetElementByLabel(page, label)
	if err != nil {
		return nil, err
	}

	if element.IsDisabled() {
		time.Sleep(500 * time.Millisecond)
		element, _ = browser.GetElementByLabel(page, label)
		return element, err
	}

	return element, err
}
