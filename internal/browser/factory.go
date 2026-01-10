package browser

import (
	"testflowkit/pkg/browser"
	"testflowkit/pkg/browser/playwright"
	"testflowkit/pkg/browser/rod"
	"testflowkit/pkg/logger"
	"time"
)

// DriverType represents the type of browser driver.
type DriverType string

const (
	DriverRod        DriverType = "rod"
	DriverPlaywright DriverType = "playwright"
)

// Config holds the configuration for creating a browser client instance.
type Config struct {
	DriverType    DriverType
	HeadlessMode  bool
	ThinkTime     time.Duration
	IncognitoMode bool
}

// CreateInstance creates a new browser client instance using the specified configuration.
// If DriverType is empty or unknown, it defaults to DriverRod.
func CreateInstance(cfg Config) browser.Client {
	if cfg.DriverType == "" {
		cfg.DriverType = DriverRod
	}

	switch cfg.DriverType {
	case DriverRod:
		logger.Info("Using ROD as browser driver")
		return rod.New(cfg.HeadlessMode, cfg.ThinkTime, cfg.IncognitoMode)
	case DriverPlaywright:
		logger.Info("Using Playwright as browser driver")
		return playwright.New(cfg.HeadlessMode, cfg.ThinkTime, cfg.IncognitoMode)
	default:
		panic("unknown driver type: " + string(cfg.DriverType))
	}
}
