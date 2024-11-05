package browser

import (
	"github.com/go-rod/rod"
)

type rodPage struct {
	page *rod.Page
}

func (p *rodPage) GetOneBySelector(selector string) (Element, error) {
	element, err := p.page.Element(selector)
	if err != nil {
		return nil, err
	}

	return newRodElement(element), nil
}

func (p *rodPage) GetAllBySelector(selector string) ([]Element, error) {
	rodElts, err := p.page.Elements(selector)
	if err != nil {
		return nil, err
	}

	var elts []Element
	for _, elt := range rodElts {
		elts = append(elts, newRodElement(elt))
	}

	return elts, nil
}

func (p *rodPage) GetInfo() PageInfo {
	return PageInfo{
		URL: p.page.MustInfo().URL,
	}
}

func (p *rodPage) GetKeyboard() Keyboard {
	return newRodKeyboard(p.page.Keyboard)
}

func (p *rodPage) HasSelector(selector string) bool {
	has, _, err := p.page.Has(selector)
	if err != nil {
		return false
	}
	return has
}

func (p *rodPage) GetOneByXPath(xpath string) (Element, error) {
	element, err := p.page.ElementX(xpath)
	if err != nil {
		return nil, err
	}
	return newRodElement(element), nil
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

func newRodPage(page *rod.Page) Page {
	return &rodPage{
		page: page,
	}
}
