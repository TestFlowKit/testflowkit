package scenario

import (
	"errors"
	"testflowkit/internal/config"

	"testflowkit/internal/browser"
	"testflowkit/internal/browser/common"

	"time"
)

type Context struct {
	frontend *Frontend
	config   *config.Config
}

func (c *Context) InitBrowser(incognitoMode bool) {
	c.frontend.browser = browser.CreateInstance(c.frontend.headlessMode, c.frontend.timeout, c.frontend.slowMotion, incognitoMode)
}

func (c *Context) OpenNewPage(url string) {
	c.frontend.page = c.frontend.browser.NewPage(url)
}

func (c *Context) GetPages() []common.Page {
	return c.frontend.browser.GetPages()
}

func (c *Context) GetCurrentPage() (common.Page, string) {
	return c.frontend.page, c.frontend.currentPageName
}

func (c *Context) GetCurrentPageOnly() common.Page {
	return c.frontend.page
}

func (c *Context) SetCurrentPage(page common.Page, pageName string) error {
	if page == nil {
		return errors.New("cannot set current page: page is nil")
	}

	page.Focus()
	page.WaitLoading()

	c.frontend.page = page

	if pageName == "" {
		pageName = "unknown"
	}
	c.frontend.currentPageName = pageName

	return nil
}

func (c *Context) SetCurrentPageOnly(page common.Page) error {
	return c.SetCurrentPage(page, "")
}

func (c *Context) GetCurrentPageKeyboard() common.Keyboard {
	return c.frontend.page.GetKeyboard()
}

func (c *Context) GetConfig() *config.Config {
	return c.config
}

func NewContext(cfg *config.Config) *Context {
	return &Context{
		frontend: &Frontend{
			browser:      nil,
			page:         nil,
			timeout:      time.Duration(cfg.Settings.DefaultTimeout) * time.Millisecond,
			headlessMode: cfg.IsHeadlessModeEnabled(),
			slowMotion:   cfg.GetSlowMotion(),
		},
		config: cfg,
	}
}

type Frontend struct {
	browser             common.Browser
	page                common.Page
	timeout, slowMotion time.Duration
	currentPageName     string
	headlessMode        bool
}
