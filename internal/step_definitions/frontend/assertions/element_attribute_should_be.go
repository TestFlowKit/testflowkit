package assertions

import (
	"context"
	"fmt"
	"reflect"
	"testflowkit/internal/config"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

func (s steps) elementAttributeShouldBe() stepbuilder.Step {
	return s.newElementAttributeStep(
		`the {string} attribute of the {string} element should be {string}`,
		true,
		stepbuilder.DocParams{
			Description: "This assertion checks if the specified attribute of an element matches the expected value exactly.",
			Variables: []stepbuilder.DocVariable{
				{Name: "attributeName", Description: "The name of the HTML attribute to check.", Type: stepbuilder.VarTypeString},
				{Name: "elementName", Description: "The logical name of the element to check.", Type: stepbuilder.VarTypeString},
				{Name: "expectedValue", Description: "The expected value of the attribute.", Type: stepbuilder.VarTypeString},
			},
			Example:    `Then the "href" attribute of the "login_link" element should be "/login"`,
			Categories: []stepbuilder.StepCategory{stepbuilder.Assertions},
		},
	)
}

func (s steps) elementAttributeShouldNotBe() stepbuilder.Step {
	return s.newElementAttributeStep(
		`the {string} attribute of the {string} element should not be {string}`,
		false,
		func() stepbuilder.DocParams {
			str := stepbuilder.VarTypeString
			vars := []stepbuilder.DocVariable{
				{Name: "attributeName", Description: "The name of the HTML attribute to check.", Type: str},
				{Name: "elementName", Description: "The logical name of the element to check.", Type: str},
				{Name: "unexpectedValue", Description: "The attribute value that should not be present.", Type: str},
			}
			return stepbuilder.DocParams{
				Description: "This assertion checks if the specified attribute of an element is different from a forbidden value.",
				Variables:   vars,
				Example:     `Then the "href" attribute of the "login_link" element should not be "/logout"`,
				Categories:  []stepbuilder.StepCategory{stepbuilder.Assertions},
			}
		}(),
	)
}

func (s steps) newElementAttributeStep(sentence string, shouldEqual bool, doc stepbuilder.DocParams) stepbuilder.Step {
	return stepbuilder.NewWithThreeVariables(
		[]string{sentence},
		func(ctx context.Context, attributeName, elementName, value string) (context.Context, error) {
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

			if shouldEqual && actualValue != value {
				msg := "attribute '%s' of element '%s' is '%s', expected '%s'"
				return ctx, fmt.Errorf(msg, attributeName, elementName, actualValue, value)
			}

			if !shouldEqual && actualValue == value {
				msg := "attribute '%s' of element '%s' should not be '%s'"
				return ctx, fmt.Errorf(msg, attributeName, elementName, value)
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
		doc,
	)
}
