package scenario

import (
	"testflowkit/internal/browser"
	"testflowkit/internal/config"

	"testflowkit/internal/browser/common"

	"time"
)

type Context struct {
	frontend  *frontend
	http      *RESTAPIContext
	config    *config.Config
	variables map[string]any
}

func (c *Context) GetConfig() *config.Config {
	return c.config
}

func (c *Context) GetHTMLElementByLabel(label string) (common.Element, error) {
	page, pageName, errPage := c.GetCurrentPage()
	if errPage != nil {
		return nil, errPage
	}
	// we can't do a cancel timeout here because it create a new page instance
	page.SetTimeout(c.frontend.timeout)
	element, err := browser.GetElementByLabel(page, pageName, label)

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
		frontend: newFrontCtx(cfg),
		http: &RESTAPIContext{
			RequestHeaders: make(map[string]string),
		},
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
	browser            common.Browser
	page               common.Page
	timeout, thinkTime time.Duration
	currentPageName    string
	headlessMode       bool
}

type HTTPResponse struct {
	StatusCode int
	Body       []byte
}
