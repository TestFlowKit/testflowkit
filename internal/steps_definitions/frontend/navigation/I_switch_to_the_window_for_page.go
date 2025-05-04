package navigation

import (
	"fmt"
	"slices"
	"strings"
	"testflowkit/internal/browser/common"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
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
				windowIdx := slices.IndexFunc(pages, func(page common.Page) bool {
					return strings.HasPrefix(page.GetInfo().URL, url)
				})

				if windowIdx == -1 {
					return fmt.Errorf("window for page %s not found", pageName)
				}

				pages[windowIdx].Focus()

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
			Description: "switches to the window for the specified page.",
			Variables: []shared.StepVariable{
				{Name: "pageName", Description: "The name of the page to switch to.", Type: shared.DocVarTypeString},
			},
			Example:  "When I switch to the window for the \"Login\" page",
			Category: shared.Navigation,
		},
	)
}
