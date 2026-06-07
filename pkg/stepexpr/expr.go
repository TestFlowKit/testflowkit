package stepexpr

import "strings"

type Wildcard = string
type WildcardID = string

const VariablePattern = `\{\{\s*([^}]+)\s*\}\}`

var (
	stringID WildcardID = "{string}"
	numberID WildcardID = "{number}"
	intID    WildcardID = "{int}"
	// Accept quoted/unquoted values and allow escaped quotes inside quoted values (e.g. \"priority\").
	stringWildcard Wildcard = `"?((?:\\.|[^"])*)"?`
	// Match integer/float numeric values (quoted or unquoted).
	numberWildcard Wildcard = `"?(-?\d+(?:\.\d+)?)"?`

	// Match integer numeric values only; quoted integers are not accepted.
	intWildcard Wildcard = `(-?\d+)`
)

var wildcards = map[WildcardID]Wildcard{
	numberID: numberWildcard,
	stringID: stringWildcard,
	intID:    intWildcard,
}

func ConvertWildcards(current string) string {
	for id, wildcard := range wildcards {
		current = strings.ReplaceAll(current, id, wildcard)
	}
	return current
}
