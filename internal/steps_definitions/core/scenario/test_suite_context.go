package scenario

import (
	"log"
	"testflowkit/internal/browser"
	"testflowkit/internal/browser/common"
	"time"
)

type Context struct {
	browser             common.Browser
	page                common.Page
	timeout, slowMotion time.Duration
	headlessMode        bool
}

func (fc *Context) InitBrowser(incognitoMode bool) {
	fc.browser = browser.CreateInstance(fc.headlessMode, fc.timeout, fc.slowMotion, incognitoMode)
}

func (fc *Context) OpenNewPage(url string) {
	fc.page = fc.browser.NewPage(url)
}

func (fc *Context) GetPages() []common.Page {
	return fc.browser.GetPages()
}

func (fc *Context) GetCurrentPage() common.Page {
	return fc.page
}

func (fc *Context) SetCurrentPage(page common.Page) {
	page.Focus()
	page.WaitLoading()
	fc.page = page
}

func (fc *Context) GetCurrentPageKeyboard() common.Keyboard {
	return fc.page.GetKeyboard()
}

func NewContext(timeout string, headlessMode bool, slowMotion time.Duration) *Context {
	duration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Panicf("timeout is not correct (%s)", timeout)
	}

	return &Context{
		browser:      nil,
		page:         nil,
		timeout:      duration,
		headlessMode: headlessMode,
		slowMotion:   slowMotion,
	}
}
