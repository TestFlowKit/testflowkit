package browser

import (
	pkgbrowser "testflowkit/pkg/browser"
	"testflowkit/pkg/browser/playwright"
	"testflowkit/pkg/browser/rod"
	"time"
)

// DriverType represents the type of browser driver.
type DriverType string

const (
	DriverRod        DriverType = "rod"
	DriverPlaywright DriverType = "playwright"
)

type Config struct {
	DriverType    DriverType
	HeadlessMode  bool
	ThinkTime     time.Duration
	IncognitoMode bool
	UserAgent     string
	Locale        string
	TimezoneID    string
	WarmUpCount   int
}

func InitEngine(cfg Config) (Engine, error) {
	switch cfg.DriverType {
	case DriverRod:
		return rod.InitEngine()
	case DriverPlaywright:
		warmUpArgs := pkgbrowser.CreationArgs{
			HeadlessMode: cfg.HeadlessMode,
			ThinkTime:    cfg.ThinkTime,
			UserAgent:    cfg.UserAgent,
			Locale:       cfg.Locale,
			TimezoneID:   cfg.TimezoneID,
		}
		return playwright.InitEngine(cfg.WarmUpCount, warmUpArgs)
	default:
		panic("unknown driver type: " + string(cfg.DriverType))
	}
}
