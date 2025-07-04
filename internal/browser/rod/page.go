package rod

import (
	"errors"
	"fmt"
	"strings"
	"testflowkit/internal/browser/common"
	"time"

	"github.com/go-rod/rod"
)

type rodPage struct {
	page *rod.Page
}

func (p *rodPage) GetOneBySelector(selector string) (common.Element, error) {
	element, err := p.page.Element(selector)
	if err != nil {
		return nil, err
	}

	return newRodElement(element), nil
}

func (p *rodPage) GetAllBySelector(selector string) ([]common.Element, error) {
	rodElts, err := p.page.Elements(selector)
	if err != nil {
		return nil, err
	}

	var elts []common.Element
	for _, elt := range rodElts {
		elts = append(elts, newRodElement(elt))
	}

	return elts, nil
}

func (p *rodPage) GetInfo() common.PageInfo {
	return common.PageInfo{
		URL: p.page.MustInfo().URL,
	}
}

func (p *rodPage) GetKeyboard() common.Keyboard {
	return newRodKeyboard(p.page.Keyboard)
}

func (p *rodPage) HasSelector(selector string) bool {
	has, _, err := p.page.Has(selector)
	if err != nil {
		return false
	}
	return has
}

func (p *rodPage) GetOneByXPath(xpath string) (common.Element, error) {
	exists, element, err := p.page.HasX(xpath)

	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("element not found")
	}

	return newRodElement(element), nil
}

func (p *rodPage) GetOneByTextContent(text string) (common.Element, error) {
	const searchTextLimit = 20

	isTextTooLong := len(text) > searchTextLimit
	limitedSearchText := text
	if isTextTooLong {
		limitedSearchText = text[:searchTextLimit]
	}

	notFoundError := errors.New("element not found")

	elements, err := p.page.ElementsX(fmt.Sprintf(`//*[contains(text(),"%s")]`, limitedSearchText))
	if err != nil {
		return nil, err
	}

	if len(elements) == 0 {
		return nil, notFoundError
	}

	if !isTextTooLong {
		return newRodElement(elements[0]), nil
	}

	var expectedElement common.Element
	for _, element := range elements {
		if strings.Contains(element.MustText(), text) {
			expectedElement = newRodElement(element)
			break
		}
	}

	if expectedElement == nil {
		return nil, notFoundError
	}

	return expectedElement, nil
}

func (p *rodPage) Focus() {
	p.page = p.page.MustActivate()
}

// TODO: be sure its work on SPA
func (p *rodPage) WaitLoading() {
	p.page.MustWaitNavigation()
	p.page = p.page.MustWaitDOMStable()
	p.page = p.page.MustWaitIdle()
}

func (p *rodPage) ExecuteJS(js string, args ...any) string {
	return p.page.MustEval(js, args...).String()
}

func (p *rodPage) Back() {
	p.page = p.page.MustNavigateBack()
}

func (p *rodPage) Refresh() {
	p.page = p.page.MustReload()
}

func (p *rodPage) Screenshot() ([]byte, error) {
	screenshot, err := p.page.Screenshot(true, nil)
	if err != nil {
		return nil, err
	}

	return screenshot, nil
}

func (p *rodPage) SetTimeout(timeout time.Duration) {
	p.page = p.page.Timeout(timeout)
}

func newRodPage(page *rod.Page) common.Page {
	return &rodPage{
		page: page,
	}
}
