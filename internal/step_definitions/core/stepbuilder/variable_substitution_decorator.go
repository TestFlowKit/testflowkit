package stepbuilder

import (
	"context"
	"fmt"
	"reflect"
	"testflowkit/internal/step_definitions/core/scenario"
)

type VariableSubstitutionDecorator struct {
	step Step
}

func NewVariableSubstitutionDecorator(step Step) Step {
	return &VariableSubstitutionDecorator{step: step}
}

func (d *VariableSubstitutionDecorator) GetSentences() []string {
	return d.step.GetSentences()
}

func (d *VariableSubstitutionDecorator) GetDocumentation() Documentation {
	return d.step.GetDocumentation()
}

func (d *VariableSubstitutionDecorator) Validate(context *ValidatorContext) any {
	return d.step.Validate(context)
}

func (d *VariableSubstitutionDecorator) GetDefinition() any {
	originalFunc := d.step.GetDefinition()
	originalFuncValue := reflect.ValueOf(originalFunc)

	if originalFuncValue.Kind() != reflect.Func {
		panic(fmt.Errorf("expected a function, but got %T", originalFunc))
	}

	wrapperFunc := func(args []reflect.Value) []reflect.Value {
		if len(args) == 0 {
			panic("context is required")
		}

		ctxValue := args[0]
		if ctxValue.Type() != reflect.TypeOf((*context.Context)(nil)).Elem() {
			panic("first argument must be context.Context")
		}

		ctx, ok := ctxValue.Interface().(context.Context)
		if !ok {
			panic("context is required")
		}

		scCtx := scenario.FromContext(ctx)
		if scCtx == nil {
			panic("context is required")
		}

		for i := 1; i < len(args); i++ {
			arg := args[i]
			if arg.Type().Kind() == reflect.String {
				newArg := scCtx.ReplaceVariableOccurence(arg.String())
				args[i] = reflect.ValueOf(newArg)
			}
		}

		return originalFuncValue.Call(args)
	}

	return reflect.MakeFunc(originalFuncValue.Type(), wrapperFunc).Interface()
}
