package logger

import (
	"fmt"
)

func Warn(msg string, actionsExpected []string) {
	const format = "%s\nActions expected: \n%s"
	log(warn, fmt.Sprintf(format, msg, formatList(actionsExpected)))
}
