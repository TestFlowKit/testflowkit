package logger

import "fmt"

func Info(message string) {
	log(info, message)
}

func Infof(format string, args ...interface{}) {
	log(info, fmt.Sprintf(format, args...))
}
