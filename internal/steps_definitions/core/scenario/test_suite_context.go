package scenario

import (
	"errors"
	"testflowkit/internal/config"

	"testflowkit/internal/browser"
	"testflowkit/internal/browser/common"

	"time"
)

type Context struct {
	browser             common.Browser
	page                common.Page
	timeout, slowMotion time.Duration
	currentPageName     string
	headlessMode        bool
	config              *config.Config
}

func (c *Context) InitBrowser(incognitoMode bool) {
	c.browser = browser.CreateInstance(c.headlessMode, c.timeout, c.slowMotion, incognitoMode)
}

func (c *Context) OpenNewPage(url string) {
	c.page = c.browser.NewPage(url)
}

func (c *Context) GetPages() []common.Page {
	return c.browser.GetPages()
}

func (c *Context) GetCurrentPage() (common.Page, string) {
	return c.page, c.currentPageName
}

func (c *Context) GetCurrentPageOnly() common.Page {
	return c.page
}

func (c *Context) SetCurrentPage(page common.Page, pageName string) error {
	if page == nil {
		return errors.New("cannot set current page: page is nil")
	}

	page.Focus()
	page.WaitLoading()

	c.page = page

	if pageName == "" {
		pageName = "unknown"
	}
	c.currentPageName = pageName

	return nil
}

func (c *Context) SetCurrentPageOnly(page common.Page) error {
	return c.SetCurrentPage(page, "")
}

func (c *Context) GetCurrentPageKeyboard() common.Keyboard {
	return c.page.GetKeyboard()
}

func (c *Context) GetConfig() *config.Config {
	return c.config
}

func NewContext(cfg *config.Config) *Context {
	return &Context{
		browser:      nil,
		page:         nil,
		timeout:      time.Duration(cfg.Settings.DefaultTimeout) * time.Millisecond,
		headlessMode: cfg.IsHeadlessModeEnabled(),
		slowMotion:   cfg.GetSlowMotion(),
		config:       cfg,
	}
}
