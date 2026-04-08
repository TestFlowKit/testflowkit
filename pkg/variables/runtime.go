package variables

import (
	"log"
	"regexp"
	"strings"
	"testflowkit/pkg/rand"
)

// ReplaceInString replaces all variable placeholders in the format {{variableName}} with their actual values.
func (p *Parser) ReplaceInString(input string) string {
	// Pattern to match {{variableName}}
	re := regexp.MustCompile(`\{\{([^}]+)\}\}`)

	return re.ReplaceAllStringFunc(input, func(match string) string {
		// Extract variable name from {{variableName}}
		varName := strings.TrimSpace(match[2 : len(match)-2])

		// Resolve rand expressions before store lookup
		if rand.IsRandExpression(varName) {
			generated, err := rand.Generate(varName)
			if err != nil {
				log.Printf("rand generation failed for '%s': %v", varName, err)
				return match
			}
			log.Printf("[rand] generated '%s' → %s", match, generated)
			return generated
		}

		// Get variable value from context
		value, exists := p.store.GetGraphQLVariable(varName)
		if !exists {
			return match // Return original if variable not found
		}

		// Convert value to string
		str, err := p.SerializeValue(value)
		if err != nil {
			return match // Return original if serialization fails
		}

		return str
	})
}

// ReplaceInBytes replaces variable placeholders in byte slice (useful for request bodies).
func (p *Parser) ReplaceInBytes(input []byte) []byte {
	replaced := p.ReplaceInString(string(input))
	return []byte(replaced)
}

// ReplaceInMap replaces variable placeholders in all values of a map (useful for headers and query params).
func (p *Parser) ReplaceInMap(input map[string]string) map[string]string {
	result := make(map[string]string)
	for key, value := range input {
		result[key] = p.ReplaceInString(value)
	}
	return result
}

type Store interface {
	GetGraphQLVariable(name string) (any, bool)
}
