package logger

import (
	"fmt"
)

func Warn(msg string, actionsExpected []string) {
	if len(actionsExpected) == 0 {
		log(warn, msg)
	} else {
		const format = "%s\nActions expected: \n%s"
		log(warn, fmt.Sprintf(format, msg, formatList(actionsExpected)))
	}
}
