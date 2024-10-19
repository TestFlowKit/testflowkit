package browser

import (
	"os"
	"time"

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
	headlessMode := os.Getenv("DEBUG") != "true"
	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).
		Headless(headlessMode).
		MustLaunch()

	const seconds = 10
	return rod.New().ControlURL(u).MustConnect().Timeout(seconds * time.Second)
}
