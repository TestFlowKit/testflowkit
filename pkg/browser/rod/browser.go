package rod

import (
	"testflowkit/pkg/browser"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// Engine manages the shared rod launcher.
type Engine struct {
	browserPath string
}

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

// ========== Engine Methods ==========

// InitEngine initializes the rod browser engine once.
func InitEngine() (*Engine, error) {
	path, _ := launcher.LookPath()

	return &Engine{
		browserPath: path,
	}, nil
}

func (e *Engine) NewBrowser(args browser.CreationArgs) browser.Client {
	launcherInstance := launcher.New().Bin(e.browserPath).Headless(args.HeadlessMode)
	u := launcherInstance.MustLaunch()

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

// Close closes the engine and all resources.
func (e *Engine) Close() {
	// Nothing to do here since each browser instance manages its own lifecycle.
}

func Install() error {
	return nil
}

func Install() error {
	return nil
}
