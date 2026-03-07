package actionrun

import (
	"fmt"
	"sync"
	"time"

	internalbrowser "testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

var browserInitOnce sync.Once

func initializeBrowserEngineIfFrontendStepExists(cfg *config.Config, stepText string) {
	if !shouldPreinitializeBrowserEngine(cfg, stepText) {
		return
	}

	browserInitOnce.Do(func() {
		preinitializeBrowserEngine(cfg, stepText)
	})
}

func shouldPreinitializeBrowserEngine(cfg *config.Config, stepText string) bool {
	if cfg == nil || !cfg.IsFrontendDefined() {
		return false
	}

	if stepText == "" {
		return false
	}

	if !IsFrontendStepTextMatch(stepText) {
		return false
	}

	return true
}

func preinitializeBrowserEngine(cfg *config.Config, stepText string) {
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
	logger.InfoFf("Browser engine pre-initialization completed (trigger: %s)", stepText)
}
