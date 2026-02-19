package playwright

import (
	"fmt"
	"os/exec"
	"regexp"
	"sync"
	"testflowkit/pkg/browser"

	"github.com/playwright-community/playwright-go"
	pw "github.com/playwright-community/playwright-go"
)

type playwrightBrowser struct {
	browser    pw.Browser
	userAgent  string
	locale     string
	timezoneID string
}

func (pb *playwrightBrowser) NewPage(url string) browser.Page {
	contextOpts := pw.BrowserNewContextOptions{}

	if pb.timezoneID != "" {
		contextOpts.TimezoneId = &pb.timezoneID
	}

	if pb.userAgent != "" {
		contextOpts.UserAgent = &pb.userAgent
	}

	if pb.locale != "" {
		contextOpts.Locale = &pb.locale
	}

	context, err := pb.browser.NewContext(contextOpts)
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
func New(args browser.CreationArgs) browser.Client {
	initPlaywright()

	var browserInstance pw.Browser
	if args.IncognitoMode {
		// For incognito mode, launch browser without persistent context
		chromium := pwInstance.Chromium
		opts := pw.BrowserTypeLaunchOptions{
			Headless: &args.HeadlessMode,
		}

		if args.ThinkTime > 0 {
			slowMo := float64(args.ThinkTime.Milliseconds())
			opts.SlowMo = &slowMo
		}

		var err error
		browserInstance, err = chromium.Launch(opts)
		if err != nil {
			panic(err)
		}
	} else {
		// Normal mode
		chromium := pwInstance.Chromium
		opts := playwright.BrowserTypeLaunchOptions{
			Headless: &args.HeadlessMode,
		}

		if args.ThinkTime > 0 {
			slowMo := float64(args.ThinkTime.Milliseconds())
			opts.SlowMo = &slowMo
		}

		var err error
		browserInstance, err = chromium.Launch(opts)
		if err != nil {
			panic(err)
		}
	}

	return &playwrightBrowser{
		browser:    browserInstance,
		userAgent:  args.UserAgent,
		locale:     args.Locale,
		timezoneID: args.TimezoneID,
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

// Install installs Playwright browser driver using go run
// Requires Go to be installed and internet connection
func Install() error {
	// Check if Go is available
	goVersionCmd := exec.Command("go", "version")
	if err := goVersionCmd.Run(); err != nil {
		return fmt.Errorf("go command not found. Please install Go from https://golang.org/dl/")
	}

	installCmd := exec.Command(
		"go",
		"run",
		"github.com/playwright-community/playwright-go/cmd/playwright@latest",
		"install",
		"--with-deps",
		"chromium",
	)

	fmt.Println("Installing Playwright browsers. This may take a few minutes...")
	output, err := installCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("playwright installation failed: %w\nOutput: %s", err, string(output))
	}

	re := regexp.MustCompile(`playwright build v\w+ downloaded`)
	isInstalled := re.Match(output)
	if isInstalled {
		fmt.Println("Playwright browsers installed successfully.")
	} else {
		return fmt.Errorf("playwright installation output did not indicate success: %s", string(output))
	}

	return nil
}
