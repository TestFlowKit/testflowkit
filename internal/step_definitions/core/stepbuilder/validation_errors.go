package stepbuilder

type ValidationErrors struct {
	missingPages    []string
	missingElements []string
	missingFiles    []string
	otherErrors     []string
	undefinedSteps  []string
}

func (ve *ValidationErrors) AddError(s string) {
	ve.otherErrors = append(ve.otherErrors, s)
}

func (ve *ValidationErrors) AddMissingPage(name string) {
	ve.missingPages = append(ve.missingPages, name)
}

func (ve *ValidationErrors) AddMissingElement(name string) {
	ve.missingElements = append(ve.missingElements, name)
}

func (ve *ValidationErrors) AddMissingFile(name string) {
	ve.missingFiles = append(ve.missingFiles, name)
}

func (ve *ValidationErrors) AddUndefinedStep(text string) {
	ve.undefinedSteps = append(ve.undefinedSteps, text)
}

func (ve *ValidationErrors) HasErrors() bool {
	frontErrors := len(ve.missingPages) > 0 || len(ve.missingElements) > 0 || len(ve.missingFiles) > 0
	otherErrors := len(ve.otherErrors) > 0 || len(ve.undefinedSteps) > 0
	return frontErrors || otherErrors
}
