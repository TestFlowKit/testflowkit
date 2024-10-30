package common

import (
	"cucumber/frontend/common/browser"
	"log"
	"time"
)

type Context struct {
	browser      browser.Browser
	page         browser.Page
	timeout      time.Duration
	headlessMode bool
}

func (fc *Context) InitBrowser() {
	fc.browser = browser.CreateInstance(fc.headlessMode)
}

func (fc *Context) OpenNewPage(url string) {
	fc.page = fc.browser.NewPage(url)
}

func (fc *Context) GetCurrentPage() browser.Page {
	return fc.page
}

func (fc *Context) GetCurrentPageKeyboard() browser.Keyboard {
	return fc.page.GetKeyboard()
}

func NewFrontendContext(timeout string, headlessMode bool) *Context {
	duration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Panicf("timeout is not correct (%s)", timeout)
	}

	return &Context{
		browser:      nil,
		page:         nil,
		timeout:      duration,
		headlessMode: headlessMode,
	}
}
