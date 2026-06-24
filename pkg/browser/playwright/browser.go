package playwright

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"runtime/debug"
	"strings"
	"sync"
	"testflowkit/pkg/browser"

	pw "github.com/playwright-community/playwright-go"
)

const playwrightModulePath = "github.com/playwright-community/playwright-go"

type Engine struct {
	mu             sync.Mutex
	initialized    bool
	instance       *pw.Playwright
	warmedBrowsers chan pw.Browser
}

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

func InitEngine(warmUpCount int, warmUpArgs browser.CreationArgs) (*Engine, error) {
	engine := &Engine{}
	engine.mu.Lock()
	defer engine.mu.Unlock()

	if !engine.initialized {
		engine.init()
		engine.initialized = true
	}

	if warmUpCount > 0 {
		engine.warmUp(warmUpCount, warmUpArgs)
	}

	return engine, nil
}

// warmUp pre-launches warmUpCount Chromium processes in parallel so that
// subsequent NewBrowser calls can skip the cold-start cost.
func (e *Engine) warmUp(count int, args browser.CreationArgs) {
	e.warmedBrowsers = make(chan pw.Browser, count)
	launchOpts := buildLaunchOptions(args)

	var wg sync.WaitGroup
	for range count {
		wg.Go(func() {
			b, err := e.instance.Chromium.Launch(launchOpts)
			if err != nil {
				log.Printf("browser warm-up: failed to pre-launch chromium: %v", err)
				return
			}
			e.warmedBrowsers <- b
		})
	}
	wg.Wait()
	log.Printf("browser warm-up: %d/%d chromium instance(s) ready", len(e.warmedBrowsers), count)
}

func (e *Engine) NewBrowser(args browser.CreationArgs) browser.Client {
	var browserInstance pw.Browser

	select {
	case warmed, ok := <-e.warmedBrowsers:
		if ok {
			browserInstance = warmed
		}
	default:
	}

	if browserInstance == nil {
		var err error
		browserInstance, err = e.instance.Chromium.Launch(buildLaunchOptions(args))
		if err != nil {
			panic(fmt.Errorf("failed to launch browser: %w", err))
		}
	}

	return &playwrightBrowser{
		browser:    browserInstance,
		userAgent:  args.UserAgent,
		locale:     args.Locale,
		timezoneID: args.TimezoneID,
	}
}

// Close closes the engine and all resources.
func (e *Engine) Close() {
	if e.warmedBrowsers != nil {
		close(e.warmedBrowsers)
		for b := range e.warmedBrowsers {
			if err := b.Close(); err != nil {
				log.Printf("browser warm-up: failed to close unused pre-warmed browser: %v", err)
			}
		}
	}

	err := e.instance.Stop()
	if err != nil {
		log.Println(err)
	}
}

// init initializes the Playwright instance with lazy initialization pattern.
func (e *Engine) init() {
	sucessMsg := "Playwright engine initialized successfully"
	runOpts := &pw.RunOptions{
		SkipInstallBrowsers: true,
	}
	inst, errFirstRun := pw.Run(runOpts)
	if errFirstRun == nil {
		e.instance = inst
		log.Println(sucessMsg)

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
	e.instance = instance
	log.Println(sucessMsg)
}

func buildLaunchOptions(args browser.CreationArgs) pw.BrowserTypeLaunchOptions {
	opts := pw.BrowserTypeLaunchOptions{
		Headless: &args.HeadlessMode,
	}

	if args.ThinkTime > 0 {
		slowMo := float64(args.ThinkTime.Milliseconds())
		opts.SlowMo = &slowMo
	}

	return opts
}

func getInstalledPlaywrightVersion(ctx context.Context) (string, error) {
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range buildInfo.Deps {
			if dep.Path != playwrightModulePath {
				continue
			}

			if dep.Replace != nil && dep.Replace.Version != "" {
				return dep.Replace.Version, nil
			}

			if dep.Version != "" && dep.Version != "(devel)" {
				return dep.Version, nil
			}
		}
	}

	listCmd := exec.CommandContext(
		ctx,
		"go",
		"list",
		"-m",
		"-f",
		"{{.Version}}",
		playwrightModulePath,
	)
	out, err := listCmd.CombinedOutput()
	if err != nil {
		errMsg := "failed to resolve playwright-go module version: %w (output: %s)"
		return "", fmt.Errorf(errMsg, err, strings.TrimSpace(string(out)))
	}

	version := strings.TrimSpace(string(out))
	if version == "" || version == "(devel)" {
		return "", fmt.Errorf("invalid playwright-go module version: %q", version)
	}

	return version, nil
}

// Install installs Playwright browser driver using go run
// Requires Go to be installed and internet connection.
func Install() error {
	// Check if Go is available
	ctx := context.Background()
	goVersionCmd := exec.CommandContext(ctx, "go", "version")
	if err := goVersionCmd.Run(); err != nil {
		return errors.New("go command not found. Please install Go from https://golang.org/dl/")
	}

	playwrightVersion, errVersion := getInstalledPlaywrightVersion(ctx)
	if errVersion != nil {
		return errVersion
	}

	playwrightCLI := fmt.Sprintf(
		"%s/cmd/playwright@%s",
		playwrightModulePath,
		playwrightVersion,
	)

	installCmd := exec.CommandContext(
		ctx,
		"go",
		"run",
		playwrightCLI,
		"install",
		"--with-deps",
		"chromium",
	)

	log.Println("Installing Playwright browsers. This may take a few minutes...")
	output, err := installCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("playwright installation failed: %w\nOutput: %s", err, string(output))
	}

	re := regexp.MustCompile(`playwright build v\w+ downloaded`)
	isInstalled := re.Match(output)
	if isInstalled {
		log.Println("Playwright browsers installed successfully.")
	} else {
		return fmt.Errorf("playwright installation output did not indicate success: %s", string(output))
	}

	return nil
}
