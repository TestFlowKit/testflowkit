package browser

import (
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
}

func InitEngine(cfg Config) (Engine, error) {
	switch cfg.DriverType {
	case DriverRod:
		return rod.InitEngine()
	case DriverPlaywright:
		// TODO: replace with actual Playwright engine initialization when implemented
		return rod.InitEngine()
	default:
		panic("unknown driver type: " + string(cfg.DriverType))
	}
}
