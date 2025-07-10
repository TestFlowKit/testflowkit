package scenario

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core"

	"testflowkit/internal/browser/common"

	"time"
)

type Context struct {
	frontend  *frontend
	http      *RESTAPIContext
	config    *config.Config
	variables map[string]any
}

func (c *Context) ReplaceVariableOccurence(sentence string) string {
	re := regexp.MustCompile(core.VariablePattern)
	matches := re.FindAllStringSubmatch(sentence, -1)
	replacedSentence := sentence

	const correctMatchLength = 2
	for _, v := range matches {
		if len(v) < correctMatchLength {
			log.Printf("Invalid variable match: %s", v)
			continue
		}

		varDef, varName := v[0], strings.TrimSpace(v[1])
		if value, exists := c.variables[varName]; exists {
			replacedSentence = strings.ReplaceAll(replacedSentence, varDef, fmt.Sprintf("%v", value))
		} else {
			log.Printf("Variable '%s' not found in context", varName)
		}
	}

	return replacedSentence
}

func (c *Context) GetConfig() *config.Config {
	return c.config
}

func (c *Context) GetHTMLElementByLabel(label string) (common.Element, error) {
	// we can't do a cancel timeout here because it create a new page instance
	c.frontend.page.SetTimeout(c.frontend.timeout)
	element, err := browser.GetElementByLabel(c.frontend.page, c.frontend.currentPageName, label)

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

func (c *Context) GetVariable(name string) (any, bool) {
	value, exists := c.variables[name]
	return value, exists
}

func (c *Context) HasVariable(name string) bool {
	_, exists := c.variables[name]
	return exists
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
		config:    cfg,
		variables: make(map[string]any),
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
