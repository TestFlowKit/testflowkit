package core

import (
	"slices"
	"sync"
)

var addLock sync.Mutex

type ValidatorContext struct {
	missingPageErrors    []string
	missingElementErrors []string
}

func (vc *ValidatorContext) AddMissingPage(name string) {
	if slices.Contains(vc.missingPageErrors, name) {
		return
	}
	vc.missingPageErrors = append(vc.missingPageErrors, name)
}

func (vc *ValidatorContext) AddMissingElement(name string) {
	if slices.Contains(vc.missingElementErrors, name) {
		return
	}
	vc.missingElementErrors = append(vc.missingElementErrors, name)
}

func (vc *ValidatorContext) HasErrors() bool {
	return len(vc.missingPageErrors) > 0 || len(vc.missingElementErrors) > 0
}

func (vc *ValidatorContext) GetErrors() []string {
	return append(vc.missingPageErrors, vc.missingElementErrors...)
}

func (vc *ValidatorContext) AddValidationErrors(errors ValidationErrors) {
	addLock.Lock()
	defer addLock.Unlock()

	vc.missingPageErrors = append(vc.missingPageErrors, errors.missingPages...)
	vc.missingElementErrors = append(vc.missingElementErrors, errors.missingElements...)
}
