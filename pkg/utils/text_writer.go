package utils

import (
	"fmt"
)

type TextWriter struct {
	textContent string
}

func (t *TextWriter) Write(b []byte) (int, error) {
	if len(t.textContent) == 0 {
		t.textContent = string(b)
		return len(b), nil
	}

	t.textContent = fmt.Sprintf("%s\n%s", t.textContent, string(b))
	return len(b), nil
}

func (t *TextWriter) String() string {
	return t.textContent
}
