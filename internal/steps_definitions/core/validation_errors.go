package core

type ValidationErrors struct {
	missingPages    []string
	missingElements []string
	undefinedSteps  []string
}

func (ve *ValidationErrors) AddMissingPage(name string) {
	ve.missingPages = append(ve.missingPages, name)
}

func (ve *ValidationErrors) AddMissingElement(name string) {
	ve.missingElements = append(ve.missingElements, name)
}

func (ve *ValidationErrors) AddUndefinedStep(text string) {
	ve.undefinedSteps = append(ve.undefinedSteps, text)
}

func (ve *ValidationErrors) HasErrors() bool {
	return len(ve.missingPages) > 0 || len(ve.missingElements) > 0
}
