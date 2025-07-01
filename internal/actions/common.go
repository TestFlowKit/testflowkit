package actions

import (
	"slices"
	"strings"
	"testflowkit/internal/steps_definitions/core"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/internal/steps_definitions/frontend"
	"testflowkit/internal/steps_definitions/restapi"
)

func GetAllSteps() []stepbuilder.Step {
	return slices.Concat(frontend.GetAllSteps(), restapi.GetAllSteps())
}

func formatStep(sentence string) string {
	cleanedSentence := strings.TrimPrefix(sentence, "^")
	cleanedSentence = strings.TrimSuffix(cleanedSentence, "$")

	pattern := "^" + core.ConvertWildcards(cleanedSentence) + "$"
	return pattern
}
