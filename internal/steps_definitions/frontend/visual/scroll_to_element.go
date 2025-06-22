package visual

import (
	"fmt"
	"testflowkit/internal/browser"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

func (s steps) scrollToElement() stepbuilder.Step {
	return stepbuilder.NewWithOneVariable(
		[]string{`^the user scrolls to the {string} element$`},
		func(ctx *scenario.Context) func(string) error {
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
		stepbuilder.DocParams{
			Description: "Scrolls the page until the specified element is visible in the viewport.",
			Variables: []stepbuilder.DocVariable{
				{Name: "name", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
			},
			Example:  `When the user scrolls to the "Submit Button at the bottom" element`,
			Category: stepbuilder.Visual,
		},
	)
}
