package playwright

import (
	"testflowkit/pkg/browser"
	"testflowkit/pkg/logger"
	"time"

	"github.com/playwright-community/playwright-go"
	pw "github.com/playwright-community/playwright-go"
)

type playwrightBrowser struct {
	browser pw.Browser
}

func (pb *playwrightBrowser) NewPage(url string) browser.Page {
	logger.InfoFf("Opening URL %s in Playwright:", url)

	context, err := pb.browser.NewContext()
	if err != nil {
		panic(err)
	}

	page, err := context.NewPage()
	if err != nil {
		panic(err)
	}

	waitUntil := pw.WaitUntilStateLoad
	_, err = page.Goto(url, pw.PageGotoOptions{
		WaitUntil: waitUntil,
	})
	if err != nil {
		panic(err)
	}

	return newPlaywrightPage(page)
}

func (pb *playwrightBrowser) GetPages() []browser.Page {
	contexts := pb.browser.Contexts()
	var pages []browser.Page

	for _, context := range contexts {
		contextPages := context.Pages()
		for _, page := range contextPages {
			pages = append(pages, newPlaywrightPage(page))
		}
	}

	return pages
}

func (pb *playwrightBrowser) Close() {
	err := pb.browser.Close()
	if err != nil {
		panic(err)
	}
}

// New creates a new Playwright browser client instance with Chromium.
func New(headlessMode bool, thinkTime time.Duration, incognitoMode bool) browser.Client {
	// move in order ro run once
	errInstall := pw.Install()
	if errInstall != nil {
		panic(errInstall)
	}

	pw, err := pw.Run()
	if err != nil {
		panic(err)
	}

	chromium := pw.Chromium
	opts := playwright.BrowserTypeLaunchOptions{
		Headless: &headlessMode,
	}

	if thinkTime > 0 {
		slowMo := float64(thinkTime.Milliseconds())
		opts.SlowMo = &slowMo
	}

	browser, err := chromium.Launch(opts)
	if err != nil {
		panic(err)
	}

	return &playwrightBrowser{
		browser: browser,
	}
}
