package scenario

import (
	internalbrowser "testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/pkg/browser"

	"time"
)

type Context struct {
	frontend  *frontend
	backend   *BackendContext
	config    *config.Config
	variables map[string]any
}

func (c *Context) GetConfig() *config.Config {
	return c.config
}

func (c *Context) GetBackendContext() *BackendContext {
	return c.backend
}

func (c *Context) GetHTMLElementByLabel(label string) (browser.Element, error) {
	page, pageName, errPage := c.GetCurrentPage()
	if errPage != nil {
		return nil, errPage
	}
	// we can't do a cancel timeout here because it create a new page instance
	page.SetTimeout(c.frontend.timeout)
	element, err := internalbrowser.GetElementByLabel(page, pageName, label)

	return element, err
}

func (c *Context) Done() {
	if c.frontend.browser != nil {
		c.frontend.browser.Close()
	}
}

func (c *Context) SetVariable(name string, value any) {
	c.variables[name] = value
}

func NewContext(cfg *config.Config) *Context {
	return &Context{
		frontend:  newFrontCtx(cfg),
		backend:   NewBackendContext(),
		config:    cfg,
		variables: make(map[string]any),
	}
}

func newFrontCtx(cfg *config.Config) *frontend {
	if !cfg.IsFrontendDefined() {
		return &frontend{}
	}

	var thinkTime = cfg.GetThinkTime()
	if cfg.IsHeadlessModeEnabled() {
		thinkTime = 0
	}

	return &frontend{
		browser:      nil,
		page:         nil,
		timeout:      cfg.GetFrontendTimeout(),
		headlessMode: cfg.IsHeadlessModeEnabled(),
		thinkTime:    thinkTime,
	}
}

type frontend struct {
	browser            browser.Client
	page               browser.Page
	timeout, thinkTime time.Duration
	currentPageName    string
	headlessMode       bool
}

type HTTPResponse struct {
	StatusCode int
	Body       []byte
}
