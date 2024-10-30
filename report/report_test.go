package report

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newReport(enabled bool, formatType string) Report {
	return New("app", "1.0.0", enabled, formatType)
}
func TestDisabledReportInstantiation(t *testing.T) {
	report := newReport(false, "")
	_, isDisabled := report.formatter.(disabledFormatter)

	assert.True(t, isDisabled)
}

func TestHTMLReportInstantiation(t *testing.T) {
	report := newReport(true, "html")
	_, isHTMLFormatter := report.formatter.(htmlReportFormatter)

	assert.True(t, isHTMLFormatter)
}

func TestGetReporterPanicBecauseFormatterNotFound(t *testing.T) {
	assert.Panics(t, func() {
		newReport(true, "ttt")
	})
}
