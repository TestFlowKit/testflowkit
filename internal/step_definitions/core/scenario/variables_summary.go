package scenario

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"testflowkit/pkg/logger"
)

const (
	nameColWidth  = 30
	typeColWidth  = 20
	valueMaxWidth = 120
)

func padRight(s string, width int) string {
	if len(s) >= width {
		return s[:width]
	}
	return s + strings.Repeat(" ", width-len(s))
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// GenerateVariablesSummary returns a table-like summary of scenario variables.
// Sensitive names are redacted; structured values are JSON-marshaled where possible
// and truncated for readability.
func (c *Context) GenerateVariablesSummary() string {
	if c == nil || len(c.variables) == 0 {
		return "(no variables)"
	}

	keys := make([]string, 0, len(c.variables))
	for k := range c.variables {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	// header
	fmt.Fprintf(&b, "%s %s %s\n", padRight("NAME", nameColWidth), padRight("TYPE", typeColWidth), "VALUE")
	b.WriteString(strings.Repeat("-", nameColWidth+1+typeColWidth+1+valueMaxWidth))
	b.WriteString("\n")

	for _, k := range keys {
		v := c.variables[k]
		var valStr string
		if logger.IsSensitiveKey(k) {
			valStr = "[REDACTED]"
		} else {
			j, err := json.Marshal(v)
			if err != nil {
				valStr = fmt.Sprintf("%v", v)
			} else {
				valStr = string(j)
			}
		}

		t := "string"
		if v != nil {
			t = fmt.Sprintf("%T", v)
		}

		fmt.Fprintf(&b, "%s %s %s\n",
			padRight(k, nameColWidth),
			padRight(t, typeColWidth),
			truncateString(valStr, valueMaxWidth))
	}

	return b.String()
}
