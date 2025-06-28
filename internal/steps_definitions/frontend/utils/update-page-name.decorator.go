package utils

import (
	"context"
	"fmt"
	"reflect"
	"testflowkit/internal/steps_definitions/core/scenario"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
)

type updatePageNameDecorator struct {
	step stepbuilder.Step
}

func (d *updatePageNameDecorator) Validate(context *stepbuilder.ValidatorContext) any {
	return d.step.Validate(context)
}

func (d *updatePageNameDecorator) GetDocumentation() stepbuilder.Documentation {
	return d.step.GetDocumentation()
}

func (d *updatePageNameDecorator) GetSentences() []string {
	return d.step.GetSentences()
}

func (d *updatePageNameDecorator) GetDefinition() any {
	originalFunc := d.step.GetDefinition()
	originalFuncValue := reflect.ValueOf(originalFunc)

	if originalFuncValue.Kind() != reflect.Func {
		return func() error {
			return fmt.Errorf("expected a function, but got %T", originalFunc)
		}
	}

	wrapperFunc := func(args []reflect.Value) []reflect.Value {
		// Extract context from the first argument
		if len(args) > 0 {
			ctxValue := args[0]
			if ctxValue.Type() == reflect.TypeOf((*context.Context)(nil)).Elem() {
				if ctx, ok := ctxValue.Interface().(context.Context); ok {
					scenarioCtx := scenario.FromContext(ctx)
					if scenarioCtx != nil {
						scenarioCtx.UpdatePageNameIfNeeded()
					}
				}
			}
		}
		return originalFuncValue.Call(args)
	}

	return reflect.MakeFunc(originalFuncValue.Type(), wrapperFunc).Interface()
}

func NewUpdatePageNameDecorator(step stepbuilder.Step) stepbuilder.Step {
	return &updatePageNameDecorator{step: step}
}
