package stringutils

import (
	"fmt"
	"regexp"
	"strings"
)

func SplitAndTrim(s, sep string) []string {
	var arr []string
	for v := range strings.SplitSeq(s, sep) {
		arr = append(arr, strings.TrimSpace(v))
	}

	return arr
}

func Inline(str string) string {
	return strings.Join(SplitAndTrim(str, "\n"), " ")
}

func ContainsIgnoreLineBreaks(str, substr string) bool {
	newStr := Inline(str)
	newSubstr := Inline(substr)
	return strings.Contains(newStr, newSubstr)
}

func SuffixWithUnderscore(str, suffix string) string {
	return fmt.Sprintf("%s_%s", strings.Trim(strings.ToLower(str), " "), strings.Trim(suffix, " "))
}

func SnakeCase(label string) string {
	return strings.ToLower(strings.ReplaceAll(label, " ", "_"))
}

func NormalizeWhitespace(text string) string {
	// Remove leading and trailing whitespace
	sanitized := text
	sanitized = strings.TrimSpace(sanitized)

	// Replace multiple consecutive whitespace characters with a single space
	sanitized = regexp.MustCompile(`\s+`).ReplaceAllString(sanitized, " ")

	sanitized = strings.ReplaceAll(sanitized, "\n", " ")

	return sanitized
}
