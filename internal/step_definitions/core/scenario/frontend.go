package scenario

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	internalbrowser "testflowkit/internal/browser"
	"testflowkit/pkg/browser"
	"testflowkit/pkg/logger"
)

var errNoCurrentPageAvailable = errors.New("no current page available")

func (c *Context) InitBrowser(incognitoMode bool) {
	frontCtx := c.frontend
	frontCtx.browser = internalbrowser.CreateInstance(internalbrowser.Config{
		DriverType:    internalbrowser.DriverRod,
		HeadlessMode:  frontCtx.headlessMode,
		ThinkTime:     frontCtx.thinkTime,
		IncognitoMode: incognitoMode,
		UserAgent:     frontCtx.userAgent,
		Locale:        frontCtx.locale,
		TimezoneID:    frontCtx.timezoneID,
	})
}

func (c *Context) OpenNewPage(url string) {
	c.EnsureBrowserInitialized()
	c.frontend.page = c.frontend.browser.NewPage(url)
	c.frontend.page.WaitLoading()
}

func (c *Context) EnsureBrowserInitialized() {
	if c.frontend.browser == nil {
		logger.Info("Browser not initialized, automatically opening browser")
		c.InitBrowser(false)
	}
}

func (c *Context) GetPages() []browser.Page {
	return c.frontend.browser.GetPages()
}

func (c *Context) GetCurrentPage() (browser.Page, string, error) {
	if c.frontend.page == nil {
		return nil, "", errNoCurrentPageAvailable
	}
	return c.frontend.page, c.frontend.currentPageName, nil
}

func (c *Context) GetCurrentPageOnly() (browser.Page, error) {
	if c.frontend.page == nil {
		return nil, errNoCurrentPageAvailable
	}
	return c.frontend.page, nil
}

func (c *Context) UpdatePageNameIfNeeded() {
	if c.frontend.page == nil {
		return
	}

	currentPageURL := c.frontend.page.GetInfo().URL
	if c.frontend.currentPageName == currentPageURL {
		return
	}

	savedPageName, err := c.getPageNameByURL(currentPageURL)
	if err != nil {
		c.frontend.currentPageName = ""
		return
	}

	c.frontend.currentPageName = savedPageName
}

func (c *Context) getPageNameByURL(completeURL string) (string, error) {
	completeURLParsed, err := url.Parse(completeURL)
	if err != nil {
		return "", errors.New("cannot parse complete URL: " + err.Error())
	}

	for pageName, pageURL := range c.config.GetFrontendPages() {
		parsedPageURL, parseErr := url.Parse(pageURL)
		if parseErr != nil {
			continue
		}

		isPageURLComplete := parsedPageURL.Scheme != ""
		if !isPageURLComplete {
			completePageURL, joinErr := url.JoinPath(c.config.GetFrontendBaseURL(), pageURL)
			if joinErr != nil {
				continue
			}
			parsedPageURL, parseErr = url.Parse(completePageURL)
			if parseErr != nil {
				continue
			}
		}

		if completeURL == parsedPageURL.String() {
			return pageName, nil
		}

		s, err1 := c.getPageNameForVariablizedURL(parsedPageURL, completeURLParsed, pageName)
		if err1 == nil {
			return s, nil
		}
	}

	return "", errors.New("page name not found")
}

func (*Context) getPageNameForVariablizedURL(varURL, currentURL *url.URL, pageName string) (string, error) {
	variableRegex := regexp.MustCompile(`:(\w+)`)
	matches := variableRegex.FindAllStringSubmatch(varURL.Path, -1)

	isURLContainsVariable := len(matches) > 0
	if !isURLContainsVariable {
		return "", errors.New("URL does not contain variable")
	}

	pathRegexStr := regexp.QuoteMeta(varURL.Path)
	for _, match := range matches {
		pathRegexStr = strings.Replace(pathRegexStr, regexp.QuoteMeta(match[0]), `(\w+)`, 1)
	}

	urlRegexStrPart := varURL.Scheme + "://" + varURL.Host
	urlRegex, _ := regexp.Compile(urlRegexStrPart + pathRegexStr)
	if urlRegex.MatchString(currentURL.String()) {
		return pageName, nil
	}
	return "", errors.New("URL does not match")
}

func (c *Context) SetCurrentPage(page browser.Page) error {
	if page == nil {
		return errors.New("cannot set current page: page is nil")
	}

	page.Focus()
	page.WaitLoading()

	c.frontend.page = page

	return nil
}

func (c *Context) GetCurrentPageKeyboard() browser.Keyboard {
	return c.frontend.page.GetKeyboard()
}
