package actionrun

import (
	"fmt"
	"sync"
	"time"

	internalbrowser "testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/pkg/gherkinparser"
	"testflowkit/pkg/logger"
)

var browserInitOnce sync.Once

func initializeBrowserEngineIfFrontendStepExists(cfg *config.Config, features []*gherkinparser.Feature) {
	if !shouldPreinitializeBrowserEngine(cfg, features) {
		return
	}

	browserInitOnce.Do(func() {
		preinitializeBrowserEngine(cfg)
	})
}

func shouldPreinitializeBrowserEngine(cfg *config.Config, features []*gherkinparser.Feature) bool {
	isFrontendConfigDefined := cfg != nil && cfg.IsFrontendDefined()
	hasFrontendStep := hasFrontendStepInFeatures(features)

	return isFrontendConfigDefined && hasFrontendStep
}

func preinitializeBrowserEngine(cfg *config.Config) {
	thinkTime, err := time.ParseDuration(fmt.Sprintf("%dms", cfg.Frontend.ThinkTime))
	if err != nil {
		thinkTime = 0
	}

	browserInstance := internalbrowser.CreateInstance(internalbrowser.Config{
		DriverType:    internalbrowser.DriverRod,
		HeadlessMode:  cfg.IsHeadlessModeEnabled(),
		ThinkTime:     thinkTime,
		IncognitoMode: false,
		UserAgent:     cfg.Frontend.UserAgent,
		Locale:        cfg.Frontend.Locale,
		TimezoneID:    cfg.Frontend.TimezoneID,
	})

	browserInstance.InitEngine()
	logger.Info("Browser engine pre-initialization completed")
}
