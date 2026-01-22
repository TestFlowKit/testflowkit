package scenario

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"testflowkit/internal/step_definitions/core"
	"testflowkit/pkg/variables"
)

func (c *Context) GetVariable(name string) (any, bool) {
	value, exists := c.variables[name]
	return value, exists
}

func (c *Context) HasVariable(name string) bool {
	_, exists := c.variables[name]
	return exists
}

func ReplaceVariablesInArray[T any](ctx *Context, data []T) ([]T, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	v := ReplaceVariablesInString(ctx, string(jsonData))

	var parsedData []T
	err = json.Unmarshal([]byte(v), &parsedData)
	if err != nil {
		return nil, err
	}

	return parsedData, nil
}

func ReplaceVariablesInString(ctx *Context, sentence string) string {
	re := regexp.MustCompile(core.VariablePattern)
	matches := re.FindAllStringSubmatch(sentence, -1)
	replacedSentence := sentence

	const correctMatchLength = 2
	for _, v := range matches {
		if len(v) < correctMatchLength {
			log.Printf("Invalid variable match: %s", v)
			continue
		}

		varDef, varName := v[0], strings.TrimSpace(v[1])

		if envKey, ok := strings.CutPrefix(varName, "env."); ok {
			if val, exists := variables.GetEnvVariable(envKey); exists {
				replacedSentence = strings.ReplaceAll(replacedSentence, varDef, val)
				continue
			}
		}

		if value, exists := ctx.variables[varName]; exists {
			replacedSentence = strings.ReplaceAll(replacedSentence, varDef, fmt.Sprintf("%v", value))
		} else {
			log.Printf("Variable '%s' not found in context", varName)
		}
	}

	return replacedSentence
}

func ReplaceVariablesInMap[T any](ctx *Context, data map[string]T) (map[string]T, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	v := ReplaceVariablesInString(ctx, string(jsonData))

	var parsedData map[string]T
	err = json.Unmarshal([]byte(v), &parsedData)
	if err != nil {
		return nil, err
	}

	return parsedData, nil
}
