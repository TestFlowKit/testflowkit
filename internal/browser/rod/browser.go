package rod

import (
	"testflowkit/internal/browser/common"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type rodBrowser struct {
	browser *rod.Browser
}

func (rb *rodBrowser) NewPage(url string) common.Page {
	page := rb.browser.MustPage(url)

	page.MustWaitNavigation()
	page = page.MustWaitIdle()
	return newRodPage(page)
}

func (rb *rodBrowser) GetPages() []common.Page {
	rodPages := rb.browser.MustPages()
	var pages []common.Page
	for _, rodPage := range rodPages {
		pages = append(pages, newRodPage(rodPage))
	}

	return pages
}

func (rb *rodBrowser) Close() {
	rb.browser.MustClose()
}

func New(headlessMode bool, timeout, thinkTime time.Duration, incognitoMode bool) common.Browser {
	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).
		Headless(headlessMode).
		MustLaunch()

	newBrowser := rod.New().ControlURL(u).SlowMotion(thinkTime).Timeout(timeout).MustConnect()
	if incognitoMode {
		newBrowser = newBrowser.MustIncognito()
	}

	return &rodBrowser{
		browser: newBrowser,
	}
}
