package actionrun

import (
	"regexp"

	"testflowkit/internal/actions/actionutils"
	"testflowkit/internal/step_definitions/frontend"
	"testflowkit/pkg/logger"
)

var frontendStepRegexes = compileFrontendStepRegexes()

func compileFrontendStepRegexes() []*regexp.Regexp {
	steps := frontend.GetAllSteps()
	compiled := make([]*regexp.Regexp, 0, len(steps))

	for _, step := range steps {
		for _, sentence := range step.GetSentences() {
			pattern := "(?i)" + actionutils.FormatStep(sentence)
			re, err := regexp.Compile(pattern)
			if err != nil {
				logger.Warn("Skipping invalid frontend step regex", []string{
					"pattern: " + pattern,
					"error: " + err.Error(),
				})
				continue
			}
			compiled = append(compiled, re)
		}
	}

	return compiled
}

func IsFrontendStepTextMatch(stepText string) bool {
	for _, frontendRegex := range frontendStepRegexes {
		if frontendRegex.MatchString(stepText) {
			return true
		}
	}

	return false
}
