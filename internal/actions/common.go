package actions

import (
	"strings"
	"testflowkit/internal/step_definitions/core"
)

func formatStep(sentence string) string {
	cleanedSentence := strings.TrimPrefix(sentence, "^")
	cleanedSentence = strings.TrimSuffix(cleanedSentence, "$")

	pattern := "^" + core.ConvertWildcards(cleanedSentence) + "$"
	return pattern
}
