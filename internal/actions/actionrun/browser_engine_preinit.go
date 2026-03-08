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
var globalBrowserEngine internalbrowser.Engine

type cBrowserEngineParams struct {
	cfg      *config.Config
	features []*gherkinparser.Feature
}

func configureBrowserEngineForFrontend(p cBrowserEngineParams) internalbrowser.Engine {
	var engine internalbrowser.Engine
	if !shouldPreinitializeBrowserEngine(p.cfg, p.features) {
		return nil
	}

	browserInitOnce.Do(func() {
		engine = preinitializeBrowserEngine(p.cfg)
		globalBrowserEngine = engine
	})
	return globalBrowserEngine
}

func shouldPreinitializeBrowserEngine(cfg *config.Config, features []*gherkinparser.Feature) bool {
	isFrontendConfigDefined := cfg != nil && cfg.IsFrontendDefined()
	hasFrontendStep := hasFrontendStepInFeatures(features)

	return isFrontendConfigDefined && hasFrontendStep
}

func preinitializeBrowserEngine(cfg *config.Config) internalbrowser.Engine {
	thinkTime, err := time.ParseDuration(fmt.Sprintf("%dms", cfg.Frontend.ThinkTime))
	if err != nil {
		thinkTime = 0
	}

	engine, errInit := internalbrowser.InitEngine(internalbrowser.Config{
		DriverType:    internalbrowser.DriverRod,
		HeadlessMode:  cfg.IsHeadlessModeEnabled(),
		ThinkTime:     thinkTime,
		IncognitoMode: false,
		UserAgent:     cfg.Frontend.UserAgent,
		Locale:        cfg.Frontend.Locale,
		TimezoneID:    cfg.Frontend.TimezoneID,
	})

	if errInit != nil {
		logger.Error("Failed to initialize browser engine", []string{errInit.Error()}, nil)
		panic(fmt.Sprintf("failed to initialize browser engine: %v", errInit))
	}

	logger.Info("Browser engine pre-initialization completed")
	return engine
}
