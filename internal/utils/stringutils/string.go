package stringutils

import (
	"fmt"
	"strings"
)

func SplitAndTrim(s, sep string) []string {
	var arr []string
	for _, v := range strings.Split(s, sep) {
		arr = append(arr, strings.TrimSpace(v))
	}

	return arr
}

func SuffixWithUnderscore(str, suffix string) string {
	return fmt.Sprintf("%s_%s", strings.Trim(str, " "), strings.Trim(suffix, " "))
}
