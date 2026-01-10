package playwright

import (
	"errors"
	"fmt"
	"strings"
	"testflowkit/pkg/browser"
	"time"

	pw "github.com/playwright-community/playwright-go"
)

type playwrightPage struct {
	page pw.Page
}

func (p *playwrightPage) GetOneBySelector(selector string) (browser.Element, error) {
	element := p.page.Locator(selector)
	count, err := element.Count()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New("element not found")
	}

	return newPlaywrightElement(element), nil
}

func (p *playwrightPage) GetAllBySelector(selector string) ([]browser.Element, error) {
	locators := p.page.Locator(selector)
	count, err := locators.Count()
	if err != nil {
		return nil, err
	}

	var elts []browser.Element
	for i := 0; i < count; i++ {
		locator := locators.Nth(i)
		elts = append(elts, newPlaywrightElement(locator))
	}

	return elts, nil
}

func (p *playwrightPage) GetInfo() browser.PageInfo {
	title, err := p.page.Title()
	if err != nil {
		title = ""
	}
	return browser.PageInfo{
		URL:   p.page.URL(),
		Title: title,
	}
}

func (p *playwrightPage) GetKeyboard() browser.Keyboard {
	return newPlaywrightKeyboard(p.page)
}

func (p *playwrightPage) HasSelector(selector string) bool {
	count, err := p.page.Locator(selector).Count()
	if err != nil {
		return false
	}
	return count > 0
}

func (p *playwrightPage) GetOneByXPath(xpath string) (browser.Element, error) {
	locator := p.page.Locator(xpath)
	count, err := locator.Count()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New("element not found")
	}

	return newPlaywrightElement(locator), nil
}

func (p *playwrightPage) GetOneByTextContent(text string) (browser.Element, error) {
	const searchTextLimit = 20

	isTextTooLong := len(text) > searchTextLimit
	limitedSearchText := text
	if isTextTooLong {
		limitedSearchText = text[:searchTextLimit]
	}

	notFoundError := errors.New("element not found")

	xpath := fmt.Sprintf(`//*[contains(text(),"%s")]`, limitedSearchText)
	locators := p.page.Locator(xpath)
	count, err := locators.Count()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, notFoundError
	}

	if !isTextTooLong {
		return newPlaywrightElement(locators), nil
	}

	for i := 0; i < count; i++ {
		locator := locators.Nth(i)
		textContent, err := locator.TextContent()
		if err != nil {
			continue
		}
		if strings.Contains(textContent, text) {
			return newPlaywrightElement(locator), nil
		}
	}

	return nil, notFoundError
}

func (p *playwrightPage) Focus() {
	// In Playwright, we can focus on the page itself if needed
	// This might not be directly equivalent to Rod's MustActivate
}

func (p *playwrightPage) WaitLoading() {
	err := p.page.WaitForLoadState(pw.PageWaitForLoadStateOptions{
		State: pw.LoadStateNetworkidle,
	})
	if err != nil {
		// Fallback to load if network idle fails
		p.page.WaitForLoadState(pw.PageWaitForLoadStateOptions{
			State: pw.LoadStateLoad,
		})
	}
}

func (p *playwrightPage) ExecuteJS(js string, args ...any) string {
	result, err := p.page.Evaluate(js, args)
	if err != nil {
		return ""
	}

	if result == nil {
		return ""
	}

	return fmt.Sprintf("%v", result)
}

func (p *playwrightPage) Back() {
	_, err := p.page.GoBack()
	if err != nil {
		panic(err)
	}
}

func (p *playwrightPage) Refresh() {
	_, err := p.page.Reload()
	if err != nil {
		panic(err)
	}
}

func (p *playwrightPage) Screenshot() ([]byte, error) {
	fullP := true
	screenshot, err := p.page.Screenshot(pw.PageScreenshotOptions{
		FullPage: &fullP,
	})
	if err != nil {
		return nil, err
	}

	return screenshot, nil
}

func (p *playwrightPage) SetTimeout(timeout time.Duration) {
	p.page.SetDefaultTimeout(float64(timeout.Milliseconds()))
}

func newPlaywrightPage(page pw.Page) browser.Page {
	return &playwrightPage{
		page: page,
	}
}
