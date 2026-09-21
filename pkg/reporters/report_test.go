package reporters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newReport(formatType string) *Report {
	return New(formatType)
}
func TestReportShouldBeDisabledBecauseReportFormatNotRecognized(t *testing.T) {
	report := newReport("")
	_, isDisabled := report.formatter.(disabledFormatter)

	assert.True(t, isDisabled)
}

func TestHTMLReportInstantiation(t *testing.T) {
	report := newReport("html")
	_, isHTMLFormatter := report.formatter.(htmlReportFormatter)

	assert.True(t, isHTMLFormatter)
}

func TestJSONReportInstantiation(t *testing.T) {
	_, ok := newReport("json").formatter.(jsonReportFormatter)
	assert.True(t, ok)
}

func TestJUnitReportInstantiation(t *testing.T) {
	_, ok := newReport("junit").formatter.(junitReportFormatter)
	assert.True(t, ok)
}
