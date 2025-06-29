package scenario

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	"testflowkit/internal/browser"
	"testflowkit/internal/browser/common"
)

func (c *Context) InitBrowser(incognitoMode bool) {
	frontCtx := c.frontend
	frontCtx.browser = browser.CreateInstance(frontCtx.headlessMode, frontCtx.timeout, frontCtx.slowMotion, incognitoMode)
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

func (c *Context) GetCurrentPageOnly() common.Page {
	return c.frontend.page
}

func (c *Context) getPageNameByURL(completeURL string) (string, error) {
	completeURLParsed, err := url.Parse(completeURL)
	if err != nil {
		return "", errors.New("cannot parse complete URL: " + err.Error())
	}

	for pageName, pageURL := range c.config.Frontend.Pages {
		parsedPageURL, parseErr := url.Parse(pageURL)
		if parseErr != nil {
			continue
		}

		isPageURLComplete := parsedPageURL.Scheme != ""
		if !isPageURLComplete {
			pageURL, err = url.JoinPath(c.config.GetFrontendBaseURL(), pageURL)
			if err != nil {
				continue
			}

			parsedPageURL, err = url.Parse(pageURL)
			if err != nil {
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

func (c *Context) SetCurrentPage(page common.Page) error {
	if page == nil {
		return errors.New("cannot set current page: page is nil")
	}

	page.Focus()
	page.WaitLoading()

	c.frontend.page = page

	return nil
}

func (c *Context) GetCurrentPageKeyboard() common.Keyboard {
	return c.frontend.page.GetKeyboard()
}
