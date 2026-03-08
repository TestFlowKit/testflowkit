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

var installOnce sync.Once
var pwInstance *pw.Playwright

const playwrightModulePath = "github.com/playwright-community/playwright-go"

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

func (pb *playwrightBrowser) InitEngine() {
	initPlaywright()
}

// New creates a new Playwright browser client instance with Chromium.
func New(args browser.CreationArgs) browser.Client {
	browserInstance, err := pwInstance.Chromium.Launch(buildLaunchOptions(args))
	if err != nil {
		panic(fmt.Errorf("failed to launch browser: %w", err))
	}

	return &playwrightBrowser{
		browser:    browserInstance,
		userAgent:  args.UserAgent,
		locale:     args.Locale,
		timezoneID: args.TimezoneID,
	}
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
