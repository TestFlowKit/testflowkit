package mouse

import "testflowkit/internal/step_definitions/core/stepbuilder"

func elementWhichContainsDocParams(elementDesc, description, example string) stepbuilder.DocParams {
	return stepbuilder.DocParams{
		Description: description,
		Variables: []stepbuilder.DocVariable{
			{
				Name:        stepbuilder.DocVarName,
				Description: elementDesc,
				Type:        stepbuilder.VarTypeString,
			},
			{
				Name:        stepbuilder.DocVarText,
				Description: stepbuilder.DocDescTextContains,
				Type:        stepbuilder.VarTypeString,
			},
		},
		Example:    example,
		Categories: []stepbuilder.StepCategory{stepbuilder.Mouse, stepbuilder.Frontend},
	}
}
