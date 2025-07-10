package visual

import (
	"slices"
	"testflowkit/internal/step_definitions/core/stepbuilder"
	"testflowkit/internal/step_definitions/frontend/visual/table"
)

type steps struct {
}

func GetSteps() []stepbuilder.Step {
	handlers := steps{}

	var otherSteps = []stepbuilder.Step{
		handlers.shouldSeeOnPage(),
		handlers.shouldNotSeeOnPage(),
		handlers.shouldSeeElementWhichContains(),
		handlers.shouldSeeOnPageXElements(),
		handlers.shouldSeeDetailsOnPage(),
		handlers.scrollToElement(),
	}
	return slices.Concat(table.GetSteps(), otherSteps)
}
