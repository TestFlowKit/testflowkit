package reporters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newReport(formatType string) Report {
	return New("app", "1.0.0", formatType)
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
