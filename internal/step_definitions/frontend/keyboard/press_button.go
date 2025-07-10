package keyboard

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"

	"github.com/go-rod/rod/lib/input"
)

func (k keyboardSteps) userPressButton() stepbuilder.Step {
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

	return stepbuilder.NewWithOneVariable(
		[]string{fmt.Sprintf(`the user presses the "(%s)" key`, strings.Join(supportedKeys, "|"))},
		func(ctx context.Context, key string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			inputKey := dic[key]
			if inputKey == '0' {
				return ctx, fmt.Errorf("%s key not recognized", key)
			}

			return ctx, scenarioCtx.GetCurrentPageKeyboard().Press(inputKey)
		},
		nil,
		stepbuilder.DocParams{
			Description: "Simulates pressing a specific keyboard key (e.g., \"Enter\", \"Tab\", \"Escape\").",
			Variables: []stepbuilder.DocVariable{
				{Name: "key", Description: "The button to press.", Type: stepbuilder.VarTypeEnum(supportedKeys...)},
			},
			Example:  "When the user presses the \"Enter\" key",
			Category: stepbuilder.Keyboard,
		},
	)
}
