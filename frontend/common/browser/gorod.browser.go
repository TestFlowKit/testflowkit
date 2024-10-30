package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type rodBrowser struct {
	browser *rod.Browser
}

func (rb *rodBrowser) NewPage(url string) Page {
	page := rb.browser.MustPage(url)

	page.MustWaitNavigation()
	page = page.MustWaitIdle()
	return newRodPage(page)
}

func newRodBrowser(headlessMode bool) Browser {
	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).
		Headless(headlessMode).
		MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()
	return &rodBrowser{
		browser: browser,
	}
}
