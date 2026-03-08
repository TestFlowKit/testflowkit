package browser

import (
	"testflowkit/pkg/browser/rod"
	"time"
)

// DriverType represents the type of browser driver.
type DriverType string

const (
	DriverRod DriverType = "rod"
	// DriverPlaywright DriverType = "playwright".
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

// InitEngine initializes a browser engine based on configuration
// Currently only supports Rod driver.
func InitEngine(_ Config) (Engine, error) {
	// Only Rod is supported for Engine pattern
	return rod.InitEngine()
}
