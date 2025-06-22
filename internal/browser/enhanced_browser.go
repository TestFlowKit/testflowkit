package browser

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

// ElementFinder provides methods for finding elements with enhanced configuration support.
// This interface allows for easier testing and potential future implementations.
type ElementFinder interface {
	FindElement(page *rod.Page, pageName, elementName string) (*rod.Element, error)
	WaitForElement(page *rod.Page, pageName, elementName string) (*rod.Element, error)
	FindElementWithTimeout(page *rod.Page, pageName, elementName string, timeout time.Duration) (*rod.Element, error)
}

// EnhancedElementFinder implements ElementFinder using the enhanced configuration.
type EnhancedElementFinder struct {
	config *config.Config
}

// NewEnhancedElementFinder creates a new element finder with enhanced configuration.
// The configuration must be loaded before creating the element finder.
//
// Returns:
//   - *EnhancedElementFinder: New instance with configuration
//   - error: Error if configuration is not available
func NewEnhancedElementFinder() (*EnhancedElementFinder, error) {
	cfg, err := config.GetEnhancedConfig()
	if err != nil {
		return nil, fmt.Errorf("enhanced configuration required for element finder: %w", err)
	}
	
	return &EnhancedElementFinder{
		config: cfg,
	}, nil
}

// FindElement uses the enhanced configuration to find elements with fallback selectors.
// This method tries selectors in priority order: page-specific first, then common selectors.
//
// The selection strategy follows these principles:
// 1. Try page-specific selectors first (most specific)
// 2. Fall back to common selectors (shared across pages)
// 3. Log which selector was successful for debugging
// 4. Provide detailed error messages for troubleshooting
//
// Parameters:
//   - page: The Rod page instance to search within
//   - pageName: The logical page name (e.g., "login_page", "dashboard_page")
//   - elementName: The element name (e.g., "email_field", "submit_button")
//
// Returns:
//   - *rod.Element: The found element
//   - error: Error if element not found with any selector
func (ef *EnhancedElementFinder) FindElement(page *rod.Page, pageName, elementName string) (*rod.Element, error) {
	selectors := ef.config.GetElementSelectors(pageName, elementName)
	if len(selectors) == 0 {
		return nil, fmt.Errorf("no selectors configured for element '%s' on page '%s'", elementName, pageName)
	}

	var lastErr error
	for i, selector := range selectors {
		element, err := page.Element(selector)
		if err == nil {
			// Log successful selector for debugging (only if not the first one)
			if i > 0 {
				logger.InfoFf("Element '%s' found using fallback selector #%d: %s", elementName, i+1, selector)
			}
			return element, nil
		}
		lastErr = err
	}

	return nil, fmt.Errorf("element '%s' not found on page '%s' with any selector (tried %d selectors): %w", 
		elementName, pageName, len(selectors), lastErr)
}

// WaitForElement waits for an element to appear using the configured default timeout.
// This method uses the enhanced configuration's default timeout setting.
//
// Parameters:
//   - page: The Rod page instance to search within
//   - pageName: The logical page name
//   - elementName: The element name
//
// Returns:
//   - *rod.Element: The found element
//   - error: Error if element not found within timeout
func (ef *EnhancedElementFinder) WaitForElement(page *rod.Page, pageName, elementName string) (*rod.Element, error) {
	return ef.FindElementWithTimeout(page, pageName, elementName, ef.config.GetTimeout())
}

// FindElementWithTimeout waits for an element to appear with a custom timeout.
// This method tries each selector with the specified timeout before moving to the next.
//
// Parameters:
//   - page: The Rod page instance to search within
//   - pageName: The logical page name
//   - elementName: The element name
//   - timeout: Custom timeout duration
//
// Returns:
//   - *rod.Element: The found element
//   - error: Error if element not found within timeout
func (ef *EnhancedElementFinder) FindElementWithTimeout(page *rod.Page, pageName, elementName string, timeout time.Duration) (*rod.Element, error) {
	selectors := ef.config.GetElementSelectors(pageName, elementName)
	if len(selectors) == 0 {
		return nil, fmt.Errorf("no selectors configured for element '%s' on page '%s'", elementName, pageName)
	}

	var lastErr error
	for i, selector := range selectors {
		element, err := page.Timeout(timeout).Element(selector)
		if err == nil {
			// Log successful selector for debugging (only if not the first one)
			if i > 0 {
				logger.InfoFf("Element '%s' found using fallback selector #%d: %s (timeout: %v)", 
					elementName, i+1, selector, timeout)
			}
			return element, nil
		}
		lastErr = err
	}

	return nil, fmt.Errorf("element '%s' not found on page '%s' within timeout %v (tried %d selectors): %w", 
		elementName, pageName, timeout, len(selectors), lastErr)
}

// Backward compatibility functions for existing code
// These functions maintain the existing API while using the enhanced configuration

// FindElementWithFallback uses the enhanced configuration to find elements with fallback selectors.
// This is a convenience function that maintains backward compatibility.
//
// Deprecated: Use EnhancedElementFinder.FindElement instead for better error handling and configuration management.
func FindElementWithFallback(page *rod.Page, pageName, elementName string) (*rod.Element, error) {
	finder, err := NewEnhancedElementFinder()
	if err != nil {
		return nil, err
	}
	return finder.FindElement(page, pageName, elementName)
}

// WaitForElementWithFallback waits for an element using fallback selectors.
// This is a convenience function that maintains backward compatibility.
//
// Deprecated: Use EnhancedElementFinder.WaitForElement instead for better configuration management.
func WaitForElementWithFallback(page *rod.Page, pageName, elementName string) (*rod.Element, error) {
	finder, err := NewEnhancedElementFinder()
	if err != nil {
		return nil, err
	}
	return finder.WaitForElement(page, pageName, elementName)
}

// Enhanced browser context utilities

// BrowserContext represents an enhanced browser context with configuration-aware settings.
type BrowserContext struct {
	Browser *rod.Browser
	Page    *rod.Page
	Config  *config.Config
	Finder  ElementFinder
}

// NewBrowserContext creates a new browser context with enhanced configuration.
// This sets up the browser with appropriate settings from the configuration.
//
// Returns:
//   - *BrowserContext: New browser context instance
//   - error: Error if configuration is not available or browser setup fails
func NewBrowserContext() (*BrowserContext, error) {
	cfg, err := config.GetEnhancedConfig()
	if err != nil {
		return nil, fmt.Errorf("enhanced configuration required for browser context: %w", err)
	}

	// Create browser launcher with configuration-specific settings
	launcher := rod.New()
	
	// Apply headless mode setting
	if cfg.IsHeadlessModeEnabled() {
		launcher = launcher.Headless(true)
	} else {
		launcher = launcher.Headless(false)
	}

	// Set slow motion if configured and not in headless mode
	if slowMotion := cfg.GetSlowMotion(); slowMotion > 0 {
		launcher = launcher.SlowMotion(slowMotion)
	}

	// Connect to browser
	browser := launcher.MustConnect()

	// Create new page
	page := browser.MustPage()

	// Set page load timeout
	page = page.Timeout(cfg.GetPageLoadTimeout())

	// Create element finder
	finder, err := NewEnhancedElementFinder()
	if err != nil {
		browser.MustClose()
		return nil, fmt.Errorf("failed to create element finder: %w", err)
	}

	logger.InfoFf("Browser context created - Headless: %t, SlowMotion: %v, PageLoadTimeout: %v", 
		cfg.IsHeadlessModeEnabled(), cfg.GetSlowMotion(), cfg.GetPageLoadTimeout())

	return &BrowserContext{
		Browser: browser,
		Page:    page,
		Config:  cfg,
		Finder:  finder,
	}, nil
}

// NavigateToPage navigates to a configured page using the enhanced configuration.
// This method resolves page names to URLs using the current environment configuration.
//
// Parameters:
//   - pageName: The logical page name as defined in the configuration
//
// Returns:
//   - error: Error if navigation fails or page not configured
func (bc *BrowserContext) NavigateToPage(pageName string) error {
	url, err := bc.Config.GetFrontendURL(pageName)
	if err != nil {
		return fmt.Errorf("failed to get URL for page '%s': %w", pageName, err)
	}

	logger.InfoFf("Navigating to page '%s': %s", pageName, url)
	
	return bc.Page.Navigate(url)
}

// FindElement finds an element using the enhanced element finder.
//
// Parameters:
//   - pageName: The logical page name
//   - elementName: The element name
//
// Returns:
//   - *rod.Element: The found element
//   - error: Error if element not found
func (bc *BrowserContext) FindElement(pageName, elementName string) (*rod.Element, error) {
	return bc.Finder.FindElement(bc.Page, pageName, elementName)
}

// WaitForElement waits for an element to appear using the enhanced element finder.
//
// Parameters:
//   - pageName: The logical page name
//   - elementName: The element name
//
// Returns:
//   - *rod.Element: The found element
//   - error: Error if element not found within timeout
func (bc *BrowserContext) WaitForElement(pageName, elementName string) (*rod.Element, error) {
	return bc.Finder.WaitForElement(bc.Page, pageName, elementName)
}

// Close closes the browser context and cleans up resources.
func (bc *BrowserContext) Close() error {
	if bc.Browser != nil {
		return bc.Browser.Close()
	}
	return nil
}

// TakeScreenshot takes a screenshot and returns the image data.
// This is useful for debugging and test failure analysis.
//
// Returns:
//   - []byte: Screenshot image data
//   - error: Error if screenshot fails
func (bc *BrowserContext) TakeScreenshot() ([]byte, error) {
	if bc.Page == nil {
		return nil, fmt.Errorf("page not available for screenshot")
	}
	
	return bc.Page.Screenshot(true, nil)
}

// GetPageTitle returns the current page title.
//
// Returns:
//   - string: Page title
//   - error: Error if title cannot be retrieved
func (bc *BrowserContext) GetPageTitle() (string, error) {
	if bc.Page == nil {
		return "", fmt.Errorf("page not available for title retrieval")
	}
	
	return bc.Page.MustInfo().Title, nil
}

// GetCurrentURL returns the current page URL.
//
// Returns:
//   - string: Current URL
//   - error: Error if URL cannot be retrieved
func (bc *BrowserContext) GetCurrentURL() (string, error) {
	if bc.Page == nil {
		return "", fmt.Errorf("page not available for URL retrieval")
	}
	
	return bc.Page.MustInfo().URL, nil
} 