package common

import (
	"cucumber/frontend/common/browser"
	"log"
	"time"
)

type TestSuiteContext struct {
	browser             browser.Browser
	page                browser.Page
	timeout, slowMotion time.Duration
	headlessMode        bool
}

func (fc *TestSuiteContext) InitBrowser() {
	fc.browser = browser.CreateInstance(fc.headlessMode, fc.slowMotion)
}

func (fc *TestSuiteContext) OpenNewPage(url string) {
	fc.page = fc.browser.NewPage(url)
}

func (fc *TestSuiteContext) GetCurrentPage() browser.Page {
	return fc.page
}

func (fc *TestSuiteContext) GetCurrentPageKeyboard() browser.Keyboard {
	return fc.page.GetKeyboard()
}

func NewFrontendContext(timeout string, headlessMode bool, slowMotion time.Duration) *TestSuiteContext {
	duration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Panicf("timeout is not correct (%s)", timeout)
	}

	return &TestSuiteContext{
		browser:      nil,
		page:         nil,
		timeout:      duration,
		headlessMode: headlessMode,
		slowMotion:   slowMotion,
	}
}
