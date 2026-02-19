package actioninstall

import (
	"errors"
	"testflowkit/internal/browser"
	"testflowkit/internal/config"

	"testflowkit/pkg/browser/playwright"
	"testflowkit/pkg/browser/rod"
	"testflowkit/pkg/logger"
)

func Execute(appConfig *config.Config, errCfg error) {
	if errCfg != nil {
		logger.Fatal("INSTALL", errCfg)
	}

	driver := appConfig.GetFrontendDriver()

	mapping := map[string]func() error{
		string(browser.DriverPlaywright): playwright.Install,
		string(browser.DriverRod):        rod.Install,
	}

	if installFunc, exists := mapping[driver]; exists {
		if err := installFunc(); err != nil {
			logger.Fatal("INSTALL", err)
		}
	} else {
		logger.Fatal("INSTALL", errors.New("unsupported frontend driver: "+driver))
	}

}
