package core

import "strings"

type Wildcard = string
type WildcardID = string

const (
	StringID       WildcardID = "{string}"
	NumberID       WildcardID = "{number}"
	NumberWildcard Wildcard   = `(\d+)`
	StringWildcard Wildcard   = `"?([^"]*)"?`
)

var wildcards = map[WildcardID]Wildcard{
	NumberID: NumberWildcard,
	StringID: StringWildcard,
}

func ConvertWildcards(current string) string {
	for id, wildcard := range wildcards {
		current = strings.ReplaceAll(current, id, wildcard)
	}
	return current
}
