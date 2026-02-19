package playwright

import (
	"errors"
	"fmt"
	"testflowkit/pkg/browser"
	"testflowkit/pkg/logger"
	"time"

	pw "github.com/playwright-community/playwright-go"
)

type playwrightPage struct {
	page pw.Page
}

func (p *playwrightPage) GetOneBySelector(selector string) (browser.Element, error) {
	element := p.page.Locator(selector)
	warningMsg := fmt.Sprintf("multiple elements found with selector: %s", selector)
	return p.handleLocator(element, warningMsg)
}

func (p *playwrightPage) GetAllBySelector(selector string) ([]browser.Element, error) {
	locators := p.page.Locator(selector)
	count, err := locators.Count()
	if err != nil {
		return nil, err
	}

	var elts []browser.Element
	for i := range count {
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
	warningMsg := fmt.Sprintf("multiple elements found with xpath selector: %s", xpath)
	return p.handleLocator(locator, warningMsg)
}

func (p *playwrightPage) GetOneByTextContent(text string) (browser.Element, error) {
	el := p.page.GetByText(text)
	warningMsg := fmt.Sprintf("multiple elements found with text content selector: %s", text)
	return p.handleLocator(el, warningMsg)
}

func (p *playwrightPage) handleLocator(el pw.Locator, warningMsg string) (browser.Element, error) {
	notFoundError := errors.New("element not found")

	count, err := el.Count()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, notFoundError
	}

	warnTips := []string{
		"Please ensure that the elements are unique on the page.",
		"retrieving the first matching element.",
	}
	if count > 1 {
		logger.Warn(warningMsg, warnTips)
	}

	return newPlaywrightElement(el.First()), nil
}

func (p *playwrightPage) Focus() {
	// In Playwright, we can focus on the page itself if needed
	// This might not be directly equivalent to Rod's MustActivate
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
