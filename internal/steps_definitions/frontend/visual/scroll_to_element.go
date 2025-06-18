package visual

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"
)

func (s steps) scrollToElement() stepbuilder.TestStep {
	return stepbuilder.NewStepWithOneVariable(
		[]string{`^the user scrolls to the {string} element$`},
		func(ctx *stepbuilder.TestSuiteContext) func(string) error {
			return func(name string) error {
				element, err := browser.GetElementByLabel(ctx.GetCurrentPage(), fmt.Sprintf("%s_element", name))
				if err != nil {
					return err
				}

				scrollErr := element.ScrollIntoView()
				if scrollErr != nil {
					return scrollErr
				}

				return nil
			}
		},
		func(name string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			variable := fmt.Sprintf("%s_element", name)
			if !testsconfig.IsElementDefined(variable) {
				vc.AddMissingElement(variable)
			}

			return vc
		},
		stepbuilder.StepDefDocParams{
			Description: "Scrolls the page until the specified element is visible in the viewport.",
			Variables: []shared.StepVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: shared.DocVarTypeString},
			},
			Example:  `When the user scrolls to the "Submit Button at the bottom" element`,
			Category: shared.Visual,
		},
	)
}
