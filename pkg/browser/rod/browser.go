package rod

import (
	"testflowkit/pkg/browser"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type rodBrowser struct {
	browser *rod.Browser
}

func (rb *rodBrowser) NewPage(url string) browser.Page {
	page := rb.browser.MustPage(url)

	page.MustWaitNavigation()
	page = page.MustWaitIdle()
	return newRodPage(page)
}

func (rb *rodBrowser) GetPages() []browser.Page {
	rodPages := rb.browser.MustPages()
	var pages []browser.Page
	for _, rodPage := range rodPages {
		pages = append(pages, newRodPage(rodPage))
	}

	return pages
}

func (rb *rodBrowser) Close() {
	rb.browser.MustClose()
}

// New creates a new rod browser client instance.
func New(headlessMode bool, thinkTime time.Duration, incognitoMode bool) browser.Client {
	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).
		Headless(headlessMode).
		MustLaunch()

	newBrowser := rod.New().ControlURL(u).SlowMotion(thinkTime).MustConnect()
	if incognitoMode {
		newBrowser = newBrowser.MustIncognito()
	}

	return &rodBrowser{
		browser: newBrowser,
	}
}
