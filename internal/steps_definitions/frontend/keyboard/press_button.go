package keyboard

import (
	"fmt"
	"strings"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/shared"

	"github.com/go-rod/rod/lib/input"
)

func (k keyboardSteps) userPressButton() stepbuilder.TestStep {
	dic := map[string]input.Key{
		"Enter":       input.Enter,
		"Tab":         input.Tab,
		"Delete":      input.Delete,
		"Escape":      input.Escape,
		"Space":       input.Space,
		"Arrow Up":    input.ArrowUp,
		"Arrow Right": input.ArrowRight,
		"Arrow Down":  input.ArrowDown,
		"Arrow Left":  input.ArrowLeft,
	}

	var supportedKeys []string
	for key := range dic {
		supportedKeys = append(supportedKeys, key)
	}

	return stepbuilder.NewStepWithOneVariable(
		[]string{fmt.Sprintf(`^the user presses the "(%s)" key$`, strings.Join(supportedKeys, "|"))},
		func(ctx *stepbuilder.TestSuiteContext) func(string) error {
			return func(key string) error {
				inputKey := dic[key]
				if inputKey == '0' {
					return fmt.Errorf("%s key not recognized", key)
				}

				return ctx.GetCurrentPageKeyboard().Press(inputKey)
			}
		},
		nil,
		stepbuilder.StepDefDocParams{
			Description: "Simulates pressing a specific keyboard key (e.g., “Enter”, “Tab”, “Escape”).",
			Variables: []shared.StepVariable{
				{Name: "key", Description: "The button to press.", Type: shared.DocVarTypeEnum(supportedKeys...)},
			},
			Example:  "When the user presses the “Enter” key",
			Category: shared.Keyboard,
		},
	)
}
