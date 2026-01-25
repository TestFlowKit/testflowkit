package playwright

import (
	"sync"
	"testflowkit/pkg/browser"
	"time"

	pw "github.com/playwright-community/playwright-go"
)

type playwrightBrowser struct {
	browser pw.Browser
}

func (pb *playwrightBrowser) NewPage(url string) browser.Page {
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

var installOnce sync.Once
var pwInstance *pw.Playwright

// New creates a new Playwright browser client instance with Chromium.
func New(headlessMode bool, thinkTime time.Duration, incognitoMode bool) browser.Client {
	initPlaywright()

	chromium := pwInstance.Chromium
	opts := pw.BrowserTypeLaunchOptions{
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

func initPlaywright() {
	installOnce.Do(func() {
		runOpts := &pw.RunOptions{
			SkipInstallBrowsers: true,
		}
		inst, errFirstRun := pw.Run(runOpts)
		if errFirstRun == nil {
			pwInstance = inst
			return
		}

		errInstall := pw.Install(&pw.RunOptions{
			Browsers: []string{
				"chromium",
			},
		})
		if errInstall != nil {
			panic(errInstall)
		}

		instance, errRun := pw.Run(runOpts)
		if errRun != nil {
			panic(errRun)
		}
		pwInstance = instance
	})
}
