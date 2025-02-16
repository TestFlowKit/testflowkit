package logger

import (
	"fmt"
	"os"
)

func Error(msg string, potentialCauses []string, potentialSolutions []string) {
	if potentialCauses == nil {
		potentialCauses = []string{}
	}

	if potentialSolutions == nil {
		potentialSolutions = []string{}
	}

	if len(potentialCauses) == 0 && len(potentialSolutions) == 0 {
		log(erro, msg)
		return
	}

	const format = "%s\n Potential causes: \n%s\n\nPotential solutions: \n%s"
	finalMsg := fmt.Sprintf(format, msg, formatList(potentialCauses), formatList(potentialSolutions))
	log(erro, finalMsg)
}

func Fatal(context string, err error) {
	if err == nil {
		log(fatal, context)
	} else {
		log(fatal, fmt.Sprintf("%s: %s", context, err))
	}
	os.Exit(1)
}
