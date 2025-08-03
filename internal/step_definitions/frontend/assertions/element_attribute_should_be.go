package assertions

import (
	"context"
	"fmt"
	"reflect"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (steps) elementAttributeShouldBe() stepbuilder.Step {
	return stepbuilder.NewWithThreeVariables(
		[]string{`the {string} attribute of the {string} element should be {string}`},
		func(ctx context.Context, attributeName, elementName, expectedValue string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			element, err := scenarioCtx.GetHTMLElementByLabel(elementName)
			if err != nil {
				return ctx, err
			}

			if !element.IsVisible() {
				return ctx, fmt.Errorf("%s is not visible", elementName)
			}

			actualValue, isString := element.GetAttributeValue(attributeName, reflect.String).(string)
			if !isString {
				msg := "attribute '%s' is not a string or does not exist on element '%s'"
				return ctx, fmt.Errorf(msg, attributeName, elementName)
			}

			if actualValue != expectedValue {
				msg := "attribute '%s' of element '%s' is '%s', expected '%s'"
				return ctx, fmt.Errorf(msg, attributeName, elementName, actualValue, expectedValue)
			}

			return ctx, nil
		},
		func(_, elementName, _ string) stepbuilder.ValidationErrors {
			vc := stepbuilder.ValidationErrors{}
			if !config.IsElementDefined(elementName) {
				vc.AddMissingElement(elementName)
			}

			return vc
		},
		stepbuilder.DocParams{
			Description: "This assertion checks if the specified attribute of an element matches the expected value exactly.",
			Variables: []stepbuilder.DocVariable{
				{Name: "attributeName", Description: "The name of the HTML attribute to check.", Type: stepbuilder.VarTypeString},
				{Name: "elementName", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
				{Name: "expectedValue", Description: "The expected value of the attribute.", Type: stepbuilder.VarTypeString},
			},
			Example:  `Then the "href" attribute of the "login_link" element should be "/login"`,
			Category: stepbuilder.Visual,
		},
	)
}
