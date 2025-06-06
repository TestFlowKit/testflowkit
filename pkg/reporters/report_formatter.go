package reporters

import (
	"log"
)

type formatter interface {
	WriteReport(details testSuiteDetails)
}

type disabledFormatter struct {
}

func (f disabledFormatter) WriteReport(details testSuiteDetails) {
	const sentence = "%d tests executed successfully at %s"
	log.Printf(sentence, len(details.Scenarios), details.StartDate)
}
