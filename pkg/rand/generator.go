package rand

import (
	"fmt"
	"strings"
)

// Generator is a function that generates a random string value given optional parameters.
type Generator func(opts map[string]string) (string, error)

var registry = map[string]Generator{
	"uuid":  genUUID,
	"email": genEmail,
	"phone": genPhone,
	"int":   genInt,
	"date":  genDate,
	"words": genWords,
	"regex": genRegex,
}

// Generate parses and executes a rand expression, returning the generated string value.
// The expr must start with "rand:" (e.g. "rand:uuid", "rand:phone:country=FR,format=e164").
func Generate(expr string) (string, error) {
	typeName, opts, err := Parse(strings.TrimSpace(expr))
	if err != nil {
		return "", err
	}
	gen, ok := registry[typeName]
	if !ok {
		return "", fmt.Errorf("rand: unknown generator type %q", typeName)
	}
	return gen(opts)
}
