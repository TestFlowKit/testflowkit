package assertions

import (
	"fmt"
	"strings"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/shared"
)

func (s steps) elementShouldNotContainText() core.TestStep {
	return core.NewStepWithTwoVariables(
		[]string{`^the {string} should not contain the text {string}$`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(name, text string) error {
				element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), name)
				if err != nil {
					return err
				}

				if !element.IsVisible() {
					return fmt.Errorf("%s is not visible", name)
				}

				if strings.Contains(element.TextContent(), text) {
					return fmt.Errorf("%s unexpectedly contains text '%s'", name, text)
				}

				return nil
			}
		},
		func(name, _ string) core.ValidationErrors {
			vc := core.ValidationErrors{}
			if !testsconfig.IsElementDefined(name) {
				vc.AddMissingElement(name)
			}

			return vc
		},
		core.StepDefDocParams{
			Description: "This assertion checks if the element’s visible text does not include the specified substring.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The name of the element to check.", Type: shared.DocVarTypeString},
				{Name: "text", Description: "The text to check.", Type: shared.DocVarTypeString},
			},
			Example:  `Then the "Welcome Message" element should not contain the text "Hello John"`,
			Category: shared.Visual,
		},
	)
}
