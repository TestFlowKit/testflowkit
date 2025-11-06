package browser

import (
	"testflowkit/pkg/browser"
	"testflowkit/pkg/browser/rod"
	"time"
)

// DriverType represents the type of browser driver.
type DriverType string

const (
	// DriverRod represents the Rod browser driver.
	DriverRod DriverType = "rod"
	// Future: DriverPlaywright DriverType = "playwright".
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
		return rod.New(cfg.HeadlessMode, cfg.ThinkTime, cfg.IncognitoMode)
	// Future: case DriverPlaywright:
	//     return playwright.New(cfg.HeadlessMode, cfg.ThinkTime, cfg.IncognitoMode)
	default:
		// Default to Rod driver
		return rod.New(cfg.HeadlessMode, cfg.ThinkTime, cfg.IncognitoMode)
	}
}
