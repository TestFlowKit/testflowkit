package rand

import (
	"errors"
	"fmt"
	"strings"
)

const randPrefix = "rand:"

// IsRandExpression reports whether the trimmed variable name is a rand expression.
func IsRandExpression(name string) bool {
	return strings.HasPrefix(name, randPrefix)
}

// Parse parses a rand expression like "rand:uuid", "rand:email:domain=kil.com",
// "rand:phone:country=FR,format=e164", "rand:regex:pattern=[A-Z]{3}-\d{4}".
// It returns the generator type and the parsed options map.
func Parse(expr string) (typeName string, opts map[string]string, err error) {
	if !IsRandExpression(expr) {
		return "", nil, fmt.Errorf("not a rand expression: %q", expr)
	}

	rest := strings.TrimPrefix(expr, randPrefix)
	if rest == "" {
		return "", nil, fmt.Errorf("rand expression has no type: %q", expr)
	}

	// Split type from options on first ':'
	idx := strings.Index(rest, ":")
	if idx == -1 {
		// No options segment
		return rest, map[string]string{}, nil
	}

	typeName = rest[:idx]
	optStr := rest[idx+1:]

	if typeName == "" {
		return "", nil, fmt.Errorf("rand expression has empty type: %q", expr)
	}

	opts, err = parseOptions(typeName, optStr)
	return typeName, opts, err
}

// parseOptions parses the key=value,key=value options string.
// Special case for "regex": the pattern may contain ',' and ':', so we extract
// the value of pattern= as the rest-of-string after "pattern=".
func parseOptions(typeName, optStr string) (map[string]string, error) {
	if optStr == "" {
		return map[string]string{}, nil
	}

	if typeName == "regex" {
		return parseRegexOptions(optStr)
	}

	return parseKeyValuePairs(optStr)
}

func parseRegexOptions(optStr string) (map[string]string, error) {
	const patternKey = "pattern="

	idx := strings.Index(optStr, patternKey)
	if idx == -1 {
		return nil, fmt.Errorf("rand:regex requires a pattern= option, got %q", optStr)
	}

	before := strings.TrimRight(optStr[:idx], ",")
	pattern := optStr[idx+len(patternKey):]
	if pattern == "" {
		return nil, errors.New("rand:regex pattern= value is empty")
	}

	opts := map[string]string{"pattern": pattern}
	if before == "" {
		return opts, nil
	}

	extra, err := parseKeyValuePairs(before)
	if err != nil {
		return nil, err
	}
	for k, v := range extra {
		opts[k] = v
	}

	return opts, nil
}

func parseKeyValuePairs(optStr string) (map[string]string, error) {
	opts := make(map[string]string)
	parts := strings.Split(optStr, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		eqIdx := strings.Index(part, "=")
		if eqIdx == -1 {
			return nil, fmt.Errorf("invalid option %q: expected key=value format", part)
		}
		key := strings.TrimSpace(part[:eqIdx])
		value := strings.TrimSpace(part[eqIdx+1:])
		if key == "" {
			return nil, fmt.Errorf("empty key in option %q", part)
		}
		opts[key] = value
	}
	return opts, nil
}
