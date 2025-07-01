package scenario

import (
	"testflowkit/internal/config"

	"testflowkit/internal/browser/common"

	"time"
)

type Context struct {
	frontend *frontend
	http     *RESTAPIContext
	config   *config.Config
}

func (c *Context) GetConfig() *config.Config {
	return c.config
}

func (c *Context) Done() {
	if c.frontend.browser != nil {
		c.frontend.browser.Close()
	}
}

func NewContext(cfg *config.Config) *Context {
	return &Context{
		frontend: &frontend{
			browser:      nil,
			page:         nil,
			timeout:      time.Duration(cfg.Settings.DefaultTimeout) * time.Millisecond,
			headlessMode: cfg.IsHeadlessModeEnabled(),
			thinkTime:    cfg.GetThinkTime(),
		},
		http: &RESTAPIContext{
			RequestHeaders: make(map[string]string),
		},
		config: cfg,
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
