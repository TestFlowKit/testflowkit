package visual

import (
	"etoolse/internal/browser"
	"etoolse/internal/config/testsconfig"
	"etoolse/internal/steps_definitions/core"
	"etoolse/shared"
	"fmt"
	"strings"
)

func (s steps) iShouldSeeAndHandleAlert() core.TestStep {
	return core.NewStepWithTwoVariables(
		[]string{`^I should see "{string}" in the alert and {string} it$`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(expectedText, action string) error {
				page := ctx.GetCurrentPage()
				
				// Check if alert is visible
				if !page.IsAlertVisible() {
					return fmt.Errorf("no alert is visible")
				}

				// Get alert text and verify it contains expected text
				alertText := page.ExecuteJS("window.alert")
				if !strings.Contains(alertText, expectedText) {
					return fmt.Errorf("alert text '%s' does not contain expected text '%s'", alertText, expectedText)
				}

				// Handle the alert based on action
				var alertAction common.AlertAction
				switch action {
				case "accept":
					alertAction = common.AlertAccept
				case "dismiss":
					alertAction = common.AlertDismiss
				default:
					return fmt.Errorf("unsupported alert action: %s", action)
				}

				return page.HandleAlert(alertAction)
			}
		},
		func(expectedText, action string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			if !testsconfig.IsElementDefined(expectedText) {
				vc.AddMissingElement(expectedText)
			}
			return vc
		},
		core.StepDefDocParams{
			Description: "checks if an alert is visible with specific text and handles it.",
			Variables: []shared.StepVariable{
				{Name: "expectedText", Description: "The text that should be visible in the alert.", Type: shared.DocVarTypeString},
				{Name: "action", Description: "The action to take on the alert (accept or dismiss).", Type: shared.DocVarTypeEnum("accept", "dismiss")},
			},
			Example:  `When I should see "Are you sure?" in the alert and "accept" it`,
			Category: shared.Visual,
		},
	)
} 