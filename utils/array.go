package utils

import "strings"

type String struct{}

func (str String) SplitAndTrim(s, sep string) []string {
	var arr []string
	for _, v := range strings.Split(s, sep) {
		arr = append(arr, strings.TrimSpace(v))
	}

	return arr
}
