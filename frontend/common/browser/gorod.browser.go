package browser

import (
	"os"

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

func newRodBrowser() Browser {
	return &rodBrowser{
		browser: instantiateRodBrowser(),
	}
}

func instantiateRodBrowser() *rod.Browser {
	path, _ := launcher.LookPath()
	if os.Getenv("DEBUG") == "true" {
		u := launcher.New().Bin(path).
			Headless(false).
			MustLaunch()
		return rod.New().ControlURL(u).MustConnect()
	}
	return rod.New().ControlURL(launcher.New().Bin(path).MustLaunch()).MustConnect()
}
