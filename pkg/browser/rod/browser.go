package rod

import (
	"testflowkit/pkg/browser"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type rodBrowser struct {
	browser    *rod.Browser
	userAgent  string
	locale     string
	timezoneID string
}

func (rb *rodBrowser) NewPage(url string) browser.Page {
	page := rb.browser.MustPage()

	if rb.timezoneID != "" {
		_ = proto.EmulationSetTimezoneOverride{TimezoneID: rb.timezoneID}.Call(page)
	}

	if rb.userAgent != "" || rb.locale != "" {
		_ = proto.EmulationSetUserAgentOverride{
			UserAgent:      rb.userAgent,
			AcceptLanguage: rb.locale,
		}.Call(page)
	}

	page = page.MustNavigate(url)

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
func New(args browser.CreationArgs) browser.Client {
	path, _ := launcher.LookPath()
	launcher := launcher.New().Bin(path).Headless(args.HeadlessMode)

	u := launcher.MustLaunch()

	newBrowser := rod.New().ControlURL(u).SlowMotion(args.ThinkTime).MustConnect()
	if args.IncognitoMode {
		newBrowser = newBrowser.MustIncognito()
	}

	return &rodBrowser{
		browser:    newBrowser,
		userAgent:  args.UserAgent,
		locale:     args.Locale,
		timezoneID: args.TimezoneID,
	}
}
