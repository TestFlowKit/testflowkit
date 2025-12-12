package stepbuilder

import (
	"fmt"
	"slices"
	"strings"
	"sync"
	"testflowkit/internal/utils/stringutils"
	"testflowkit/pkg/logger"
)

var addLock sync.Mutex
var addStepLock sync.Mutex

const categoryTypeIdent = 2
const variableIdent = 3

type ValidatorContext struct {
	missingPages    []string
	missingElements []string
	undefinedSteps  []string
}

func (vc *ValidatorContext) addMissingPage(label string) {
	key := stringutils.SnakeCase(label)

	if slices.Contains(vc.missingPages, key) {
		return
	}
	vc.missingPages = append(vc.missingPages, key)
}

func (vc *ValidatorContext) addMissingElement(label string) {
	key := stringutils.SnakeCase(label)
	if slices.Contains(vc.missingElements, key) {
		return
	}
	vc.missingElements = append(vc.missingElements, key)
}

func (vc *ValidatorContext) HasErrors() bool {
	return vc.HasMissingPages() || vc.HasMissingElements() || vc.HasUndefinedSteps()
}

func (vc *ValidatorContext) HasMissingPages() bool {
	return len(vc.missingPages) > 0
}

func (vc *ValidatorContext) HasMissingElements() bool {
	return len(vc.missingElements) > 0
}

func (vc *ValidatorContext) HasUndefinedSteps() bool {
	return len(vc.undefinedSteps) > 0
}

func (vc *ValidatorContext) GetUndefinedSteps() []string {
	return vc.undefinedSteps
}

func (vc *ValidatorContext) GetElementsErrorsFormatted() string {
	lines := []string{
		"Add the following elements to the configuration file:",
		logger.GetIndents(1) + "elements:",
		logger.GetIndents(categoryTypeIdent) + "common:",
		logger.GetIndents(variableIdent) + "...",
	}

	idnt := logger.GetIndents(variableIdent)
	elementFormat := idnt + "%s:\n" + idnt + " -"

	for _, element := range vc.missingElements {
		lines = append(lines, fmt.Sprintf(elementFormat, element))
	}
	return strings.Join(lines, "\n")
}

func (vc *ValidatorContext) GetPagesErrorsFormatted() string {
	lines := []string{
		"Add the following pages to the configuration file:",
		logger.GetIndents(1) + "global:",
		logger.GetIndents(categoryTypeIdent) + "pages:",
		logger.GetIndents(variableIdent) + "...",
	}

	pageFormat := logger.GetIndents(variableIdent) + "%s:"

	for _, page := range vc.missingPages {
		lines = append(lines, fmt.Sprintf(pageFormat, page))
	}
	return strings.Join(lines, "\n")
}

func (vc *ValidatorContext) AddValidationErrors(errors ValidationErrors) {
	addLock.Lock()
	defer addLock.Unlock()

	for _, mp := range errors.missingPages {
		vc.addMissingPage(mp)
	}

	for _, me := range errors.missingElements {
		vc.addMissingElement(me)
	}
}

func (vc *ValidatorContext) AddUndefinedStep(text string) {
	addStepLock.Lock()
	defer addStepLock.Unlock()

	if slices.Contains(vc.undefinedSteps, text) {
		return
	}

	vc.undefinedSteps = append(vc.undefinedSteps, text)
}
