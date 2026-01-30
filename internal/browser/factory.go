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
	UserAgent     string
	Locale        string
	TimezoneID    string
}

// CreateInstance creates a new browser client instance using the specified configuration.
// If DriverType is empty or unknown, it defaults to DriverRod.
func CreateInstance(cfg Config) browser.Client {
	if cfg.DriverType == "" {
		cfg.DriverType = DriverRod
	}

	args := browser.CreationArgs{
		HeadlessMode:  cfg.HeadlessMode,
		ThinkTime:     cfg.ThinkTime,
		IncognitoMode: cfg.IncognitoMode,
		UserAgent:     cfg.UserAgent,
		Locale:        cfg.Locale,
		TimezoneID:    cfg.TimezoneID,
	}

	switch cfg.DriverType {
	case DriverRod:
		return rod.New(args)
	// Future: case DriverPlaywright:
	//     return playwright.New(cfg.HeadlessMode, cfg.ThinkTime, cfg.IncognitoMode)
	default:
		// Default to Rod driver
		return rod.New(args)
	}
}
