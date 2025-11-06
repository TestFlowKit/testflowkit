package keyboard

import (
	"context"
	"fmt"
	"strings"
	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/pkg/browser"
)

func (k keyboardSteps) userPressButton() stepbuilder.Step {
	dic := map[string]browser.Key{
		"Enter":       browser.KeyEnter,
		"Tab":         browser.KeyTab,
		"Delete":      browser.KeyDelete,
		"Escape":      browser.KeyEscape,
		"Space":       browser.KeySpace,
		"Arrow Up":    browser.KeyArrowUp,
		"Arrow Right": browser.KeyArrowRight,
		"Arrow Down":  browser.KeyArrowDown,
		"Arrow Left":  browser.KeyArrowLeft,
	}

	var supportedKeys []string
	for key := range dic {
		supportedKeys = append(supportedKeys, key)
	}

	return stepbuilder.NewWithOneVariable(
		[]string{fmt.Sprintf(`the user presses the "(%s)" key`, strings.Join(supportedKeys, "|"))},
		func(ctx context.Context, key string) (context.Context, error) {
			scenarioCtx := scenario.MustFromContext(ctx)
			inputKey, ok := dic[key]
			if !ok {
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
