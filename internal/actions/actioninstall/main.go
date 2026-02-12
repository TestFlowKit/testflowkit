package actioninstall

import (
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

	if driver == string(browser.DriverRod) {
		rod.Install()
	}

	mapping := map[string]func() error{
		string(browser.DriverPlaywright): playwright.Install,
		string(browser.DriverRod):        rod.Install,
	}

	if installFunc, exists := mapping[driver]; exists {
		if err := installFunc(); err != nil {
			logger.Fatal("INSTALL", err)
		}
	} else {
		logger.InfoFf("No installation needed, please specify a valid driver in your config.yml (e.g., 'playwright' or 'rod')", nil, nil)
	}

}
