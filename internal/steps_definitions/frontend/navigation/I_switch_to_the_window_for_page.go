package navigation

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"testflowkit/internal/browser/common"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/pkg/logger"
	"testflowkit/shared"
)

func (n navigation) iSwitchToTheWindowForPage() core.TestStep {
	return core.NewStepWithOneVariable(
		[]string{"^I switch to the window for the {string} page$"},
		func(ctx *core.TestSuiteContext) func(string) error {
			return func(pageName string) error {
				url, err := testsconfig.GetPageURL(pageName)
				if err != nil {
					return err
				}

				pages := ctx.GetPages()
				if len(pages) == 0 {
					return errors.New("no open browser windows found")
				}

				// Find the window with a matching URL
				windowIdx := slices.IndexFunc(pages, func(page common.Page) bool {
					pageURL := page.GetInfo().URL
					logger.Info(fmt.Sprintf("Checking page URL: %s against target: %s", pageURL, url))
					return strings.HasPrefix(pageURL, url) || strings.Contains(pageURL, url)
				})

				if windowIdx == -1 {
					// If exact match not found, try checking for popup windows
					logger.Warn(fmt.Sprintf("Window for page %s not found by URL, checking for any popup windows", pageName), nil)

					// If we have more than one page and we're looking for a popup
					if len(pages) > 1 {
						// Most recent window is usually the popup (typically last in the list)
						windowIdx = len(pages) - 1
						logger.Info(fmt.Sprintf("Using latest window (index %d) as popup", windowIdx))
					} else {
						return fmt.Errorf("window for page %s not found", pageName)
					}
				}

				// Focus on the window
				targetPage := pages[windowIdx]
				targetPage.Focus()

				// Wait for the page to load completely
				targetPage.WaitLoading()

				// CRITICAL FIX: Update the current page in the test context
				ctx.SetCurrentPage(targetPage)

				logger.Info(fmt.Sprintf("Switched to window with URL: %s", targetPage.GetInfo().URL))

				return nil
			}
		},
		func(pageName string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			if !testsconfig.IsPageDefined(pageName) {
				vc.AddMissingPage(pageName)
			}

			return vc
		},

		core.StepDefDocParams{
			Description: "switches to the window for the specified page or to a popup window.",
			Variables: []shared.StepVariable{
				{Name: "pageName", Description: "The name of the page to switch to.", Type: shared.DocVarTypeString},
			},
			Example:  "When I switch to the window for the \"Login\" page",
			Category: shared.Navigation,
		},
	)
}
