package core

import (
	"strings"
)

type Wildcard = string
type WildcardID = string

const VariablePattern = `\{\{\s*([^}]+)\s*\}\}`

var (
	stringID WildcardID = "{string}"
	numberID WildcardID = "{number}"
	wildcard Wildcard   = `"?([^"]*)"?`
)

var wildcards = map[WildcardID]Wildcard{
	numberID: wildcard,
	stringID: wildcard,
}

func ConvertWildcards(current string) string {
	for id, wildcard := range wildcards {
		current = strings.ReplaceAll(current, id, wildcard)
	}
	return current
}
