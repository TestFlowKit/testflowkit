package keyboard

import (
	"etoolse/frontend/common"
	"fmt"
	"strings"

	"github.com/go-rod/rod/lib/input"
)

func (k keyboardSteps) iPressButton() common.TestStep {
	dic := map[string]input.Key{
		"enter":       input.Enter,
		"tab":         input.Tab,
		"delete":      input.Delete,
		"escape":      input.Escape,
		"space":       input.Space,
		"arrow up":    input.ArrowUp,
		"arrow right": input.ArrowRight,
		"arrow down":  input.ArrowDown,
		"arrow left":  input.ArrowLeft,
	}

	var supportedKeys []string
	for key := range dic {
		supportedKeys = append(supportedKeys, key)
	}

	return common.NewStepWithOneVariable(
		[]string{fmt.Sprintf(`^I press the "(%s)" button$`, strings.Join(supportedKeys, "|"))},
		func(ctx *common.TestSuiteContext) func(string) error {
			return func(button string) error {
				inputKey := dic[button]
				if inputKey == '0' {
					return fmt.Errorf("%s button not recognized", button)
				}

				return ctx.GetCurrentPageKeyboard().Press(inputKey)
			}
		},
		nil,
	)
}
