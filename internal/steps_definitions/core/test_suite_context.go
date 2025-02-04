package core

import (
	"etoolse/internal/browser"
	"etoolse/internal/browser/common"
	"log"
	"time"
)

type TestSuiteContext struct {
	browser             common.Browser
	page                common.Page
	timeout, slowMotion time.Duration
	headlessMode        bool
}

func (fc *TestSuiteContext) InitBrowser(incognitoMode bool) {
	fc.browser = browser.CreateInstance(fc.headlessMode, fc.timeout, fc.slowMotion, incognitoMode)
}

func (fc *TestSuiteContext) OpenNewPage(url string) {
	fc.page = fc.browser.NewPage(url)
}

func (fc *TestSuiteContext) GetCurrentPage() common.Page {
	return fc.page
}

func (fc *TestSuiteContext) GetCurrentPageKeyboard() common.Keyboard {
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
