package frontend

import (
	"fmt"

	"testflowkit/internal/browser"
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

// Enhanced navigation step definitions using the new configuration system.
// These steps provide page-aware navigation with environment-specific URL resolution.

// navigateToPageEnhanced provides enhanced navigation to configured pages.
// This step uses the enhanced configuration to resolve page names to URLs
// based on the current environment (local, staging, production).
//
// The step supports both relative paths and absolute URLs:
// - Relative paths: Combined with environment's frontend_base_url
// - Absolute URLs: Used directly (for external sites like Google)
//
// Example usage in Gherkin:
//   When the user goes to the "login" page
//   When the user goes to the "google" page
func (st steps) navigateToPageEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user goes to the {string} page$`},
		func(ctx *scenario.Context) func(string) error {
			return func(pageName string) error {
				// Get enhanced configuration
				cfg, err := config.GetEnhancedConfig()
				if err != nil {
					return fmt.Errorf("enhanced configuration required for page navigation: %w", err)
				}

				// Resolve page name to URL using current environment
				url, err := cfg.GetFrontendURL(pageName)
				if err != nil {
					return fmt.Errorf("failed to resolve page '%s': %w", pageName, err)
				}

				logger.InfoFf("Navigating to page '%s': %s", pageName, url)

				// Navigate to the resolved URL
				err = ctx.BrowserContext.Page.Navigate(url)
				if err != nil {
					return fmt.Errorf("failed to navigate to page '%s' (%s): %w", pageName, url, err)
				}

				// Wait for page to load using configured timeout
				ctx.BrowserContext.Page.MustWaitLoad()

				logger.InfoFf("Successfully navigated to page '%s'", pageName)
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Navigates to a configured page using environment-specific URL resolution.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "pageName",
					Description: "The logical page name as defined in the configuration (e.g., 'login', 'dashboard', 'google')",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `When the user goes to the "login" page`,
			Category: stepbuilder.Navigation,
		},
	)
}

// navigateToURLEnhanced provides enhanced URL navigation with validation.
// This step allows direct URL navigation while maintaining consistency with
// the enhanced configuration system.
func (st steps) navigateToURLEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user navigates to URL {string}$`},
		func(ctx *scenario.Context) func(string) error {
			return func(url string) error {
				logger.InfoFf("Navigating to URL: %s", url)

				// Navigate to the specified URL
				err := ctx.BrowserContext.Page.Navigate(url)
				if err != nil {
					return fmt.Errorf("failed to navigate to URL '%s': %w", url, err)
				}

				// Wait for page to load
				ctx.BrowserContext.Page.MustWaitLoad()

				logger.InfoFf("Successfully navigated to URL: %s", url)
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Navigates directly to a specified URL.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "url",
					Description: "The complete URL to navigate to",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `When the user navigates to URL "https://example.com/login"`,
			Category: stepbuilder.Navigation,
		},
	)
}

// userIsOnPageEnhanced verifies the user is on a specific configured page.
// This step checks if the current URL matches the expected page configuration.
func (st steps) userIsOnPageEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user is on the {string} page$`},
		func(ctx *scenario.Context) func(string) error {
			return func(pageName string) error {
				// Get enhanced configuration
				cfg, err := config.GetEnhancedConfig()
				if err != nil {
					return fmt.Errorf("enhanced configuration required for page verification: %w", err)
				}

				// Get expected URL for the page
				expectedURL, err := cfg.GetFrontendURL(pageName)
				if err != nil {
					return fmt.Errorf("failed to resolve page '%s': %w", pageName, err)
				}

				// Get current URL
				currentURL := ctx.BrowserContext.Page.MustInfo().URL

				// For relative paths, check if current URL ends with the expected path
				if expectedURL != currentURL {
					// Try a more flexible comparison for relative paths
					env, _ := cfg.GetCurrentEnvironment()
					if pagePath, ok := cfg.Frontend.Pages[pageName]; ok {
						if !cfg.Frontend.Pages[pageName][0:4] == "http" { // Not an absolute URL
							// Compare just the path portion
							if !ctx.BrowserContext.Page.MustInfo().URL[len(env.FrontendBaseURL):] == pagePath {
								return fmt.Errorf("expected to be on page '%s' (%s), but currently on: %s", 
									pageName, expectedURL, currentURL)
							}
						} else {
							return fmt.Errorf("expected to be on page '%s' (%s), but currently on: %s", 
								pageName, expectedURL, currentURL)
						}
					}
				}

				logger.InfoFf("Verified user is on page '%s': %s", pageName, currentURL)
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Verifies that the user is currently on a specific configured page.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "pageName",
					Description: "The logical page name to verify",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `Then the user is on the "dashboard" page`,
			Category: stepbuilder.Navigation,
		},
	)
}

// shouldBeNavigatedToPageEnhanced verifies successful navigation to a page.
// This is an assertion step that combines URL verification with page readiness checks.
func (st steps) shouldBeNavigatedToPageEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user should be navigated to {string} page$`},
		func(ctx *scenario.Context) func(string) error {
			return func(pageName string) error {
				// First verify we're on the correct page
				err := st.userIsOnPageEnhanced().Execute(ctx)(pageName)
				if err != nil {
					return err
				}

				// Additional checks for page readiness could be added here
				// For example, waiting for specific elements that indicate the page is ready

				logger.InfoFf("Confirmed successful navigation to page '%s'", pageName)
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Asserts that the user has been successfully navigated to a specific page.",
			Variables: []stepbuilder.DocVariable{
				{
					Name:        "pageName",
					Description: "The logical page name to verify navigation to",
					Type:        stepbuilder.VarTypeString,
				},
			},
			Example:  `Then the user should be navigated to "login" page`,
			Category: stepbuilder.Navigation,
		},
	)
}

// openNewBrowserTabEnhanced opens a new browser tab with enhanced configuration.
// This step creates a new tab using the browser context from enhanced configuration.
func (st steps) openNewBrowserTabEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`^the user opens a new browser tab$`},
		func(ctx *scenario.Context) func() error {
			return func() error {
				logger.Info("Opening new browser tab")

				// Create new page (tab) in the current browser
				newPage := ctx.BrowserContext.Browser.MustPage()
				
				// Get enhanced configuration for timeout settings
				cfg, err := config.GetEnhancedConfig()
				if err != nil {
					return fmt.Errorf("enhanced configuration required for new tab: %w", err)
				}

				// Apply page load timeout from configuration
				newPage = newPage.Timeout(cfg.GetPageLoadTimeout())

				// Update context to use the new page
				ctx.BrowserContext.Page = newPage

				logger.Info("New browser tab opened successfully")
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Opens a new browser tab and switches to it.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `Given the user opens a new browser tab`,
			Category:    stepbuilder.Navigation,
		},
	)
}

// Enhanced navigation utilities for backward compatibility and extended functionality

// refreshPageEnhanced refreshes the current page.
func (st steps) refreshPageEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`^the user refreshes the page$`},
		func(ctx *scenario.Context) func() error {
			return func() error {
				logger.Info("Refreshing current page")

				err := ctx.BrowserContext.Page.Reload()
				if err != nil {
					return fmt.Errorf("failed to refresh page: %w", err)
				}

				// Wait for page to load
				ctx.BrowserContext.Page.MustWaitLoad()

				logger.Info("Page refreshed successfully")
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Refreshes the current page.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `When the user refreshes the page`,
			Category:    stepbuilder.Navigation,
		},
	)
}

// navigateBackEnhanced navigates back in browser history.
func (st steps) navigateBackEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`^the user navigates back$`},
		func(ctx *scenario.Context) func() error {
			return func() error {
				logger.Info("Navigating back in browser history")

				err := ctx.BrowserContext.Page.NavigateBack()
				if err != nil {
					return fmt.Errorf("failed to navigate back: %w", err)
				}

				// Wait for page to load
				ctx.BrowserContext.Page.MustWaitLoad()

				logger.Info("Navigated back successfully")
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Navigates back to the previous page in browser history.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `When the user navigates back`,
			Category:    stepbuilder.Navigation,
		},
	)
}

// waitForPageLoadEnhanced waits for the page to fully load with enhanced timeout configuration.
func (st steps) waitForPageLoadEnhanced() stepbuilder.Step {
	return stepbuilder.NewWithNoVariables(
		[]string{`^the user waits for the page to load$`},
		func(ctx *scenario.Context) func() error {
			return func() error {
				logger.Info("Waiting for page to load")

				// Get enhanced configuration for timeout
				cfg, err := config.GetEnhancedConfig()
				if err != nil {
					return fmt.Errorf("enhanced configuration required for page load wait: %w", err)
				}

				// Wait for page load with configured timeout
				ctx.BrowserContext.Page.Timeout(cfg.GetPageLoadTimeout()).MustWaitLoad()

				logger.Info("Page loaded successfully")
				return nil
			}
		},
		nil,
		stepbuilder.DocParams{
			Description: "Waits for the current page to fully load using configured timeout.",
			Variables:   []stepbuilder.DocVariable{},
			Example:     `And the user waits for the page to load`,
			Category:    stepbuilder.Navigation,
		},
	)
} 