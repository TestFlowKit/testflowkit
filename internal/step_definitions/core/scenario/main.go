package scenario

import (
	"fmt"
	"maps"
	internalbrowser "testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/pkg/browser"
	"testflowkit/pkg/variables"

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

func NewContext(cfg *config.Config, initialVariables map[string]any) *Context {
	initVars := make(map[string]any)
	if initialVariables != nil {
		maps.Copy(initVars, initialVariables)
	}

	return &Context{
		frontend:  newFrontCtx(cfg),
		backend:   newBackendCtx(),
		config:    cfg,
		variables: initVars,
	}
}

func newFrontCtx(cfg *config.Config) *frontend {
	if !cfg.IsFrontendDefined() {
		return &frontend{}
	}

	thinkTime, err := time.ParseDuration(fmt.Sprintf("%dms", cfg.Frontend.ThinkTime))
	if err != nil {
		thinkTime = 0
	}

	driver := internalbrowser.DriverType(cfg.GetFrontendDriver())

	return &frontend{
		browser:      nil,
		page:         nil,
		timeout:      cfg.GetFrontendTimeout(),
		headlessMode: cfg.IsHeadlessModeEnabled(),
		thinkTime:    thinkTime,
		userAgent:    cfg.Frontend.UserAgent,
		locale:       cfg.Frontend.Locale,
		timezoneID:   cfg.Frontend.TimezoneID,
		driverType:   driver,
	}
}

func newBackendCtx() *BackendContext {
	bc := &BackendContext{
		Headers: make(map[string]string),
		GraphQL: GraphQLContext{
			Variables: make(map[string]any),
		},
	}
	bc.parser = variables.NewParser(bc)
	return bc
}

type frontend struct {
	browser            browser.Client
	page               browser.Page
	timeout, thinkTime time.Duration
	currentPageName    string
	headlessMode       bool
	driverType         internalbrowser.DriverType
	userAgent          string
	locale             string
	timezoneID         string
}

type HTTPResponse struct {
	StatusCode int
	Body       []byte
}
