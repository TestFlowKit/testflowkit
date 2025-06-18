package stepbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorContext_AddMissingElement_dontAddDuplicate(t *testing.T) {
	vc := ValidatorContext{}

	vc.addMissingElement("element")
	vc.addMissingElement("element")

	assert.Len(t, vc.missingElements, 1)
}

func TestValidatorContext_AddMissingPage_dontAddDuplicate(t *testing.T) {
	vc := ValidatorContext{}

	vc.addMissingPage("page")
	vc.addMissingPage("page")

	assert.Len(t, vc.missingPages, 1)
}

func TestValidatorContext_AddMissingElement_Convert_to_key_before_adding(t *testing.T) {
	vc := ValidatorContext{}

	vc.addMissingElement("element one")

	assert.Contains(t, vc.missingElements, "element_one")
}

func TestValidatorContext_AddMissingPage_Convert_to_key_before_adding(t *testing.T) {
	vc := ValidatorContext{}

	vc.addMissingPage("page one")

	assert.Contains(t, vc.missingPages, "page_one")
}
