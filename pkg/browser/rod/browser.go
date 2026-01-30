package rod

import (
	"context"
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
func New(headlessMode bool, thinkTime time.Duration, incognitoMode bool, userAgent, locale, timezoneId string) browser.Client {
	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).
		Headless(headlessMode).
		MustLaunch()

	newBrowser := rod.New().ControlURL(u).SlowMotion(thinkTime).MustConnect()
	if incognitoMode {
		newBrowser = newBrowser.MustIncognito()
	}

	hasSpecificBrowserSetup := userAgent != "" || locale != "" || timezoneId != ""
	if !hasSpecificBrowserSetup {
		return &rodBrowser{
			browser: newBrowser,
		}
	}

	// Set user agent and locale using DevTools Protocol
	if userAgent != "" || locale != "" {
		emulationParams := map[string]any{
			"userAgent": userAgent,
		}
		if locale != "" {
			emulationParams["acceptLanguage"] = locale
		}
		_, _ = newBrowser.Call(context.TODO(), "", "Emulation.setUserAgentOverride", emulationParams)
	}

	// Set timezone if provided
	if timezoneId != "" {
		_, _ = newBrowser.Call(context.TODO(), "", "Emulation.setTimezoneOverride", map[string]string{
			"timezoneId": timezoneId,
		})
	}

	return &rodBrowser{
		browser: newBrowser,
	}
}
