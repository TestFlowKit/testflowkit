package queryable

import "strings"

// ValueType is the normalized type of a queried response value.
type ValueType string

const (
	TypeNull    ValueType = "null"
	TypeString  ValueType = "string"
	TypeNumber  ValueType = "number"
	TypeInteger ValueType = "integer"
	TypeBoolean ValueType = "boolean"
	TypeObject  ValueType = "object"
	TypeArray   ValueType = "array"
)

// GetValueType returns the normalized type name of a value.
func GetValueType(value any) string {
	if value == nil {
		return string(TypeNull)
	}

	switch v := value.(type) {
	case bool:
		return string(TypeBoolean)
	case float64:
		if v == float64(int64(v)) {
			return string(TypeInteger)
		}
		return string(TypeNumber)
	case string:
		return string(TypeString)
	case map[string]any:
		return string(TypeObject)
	case []any:
		return string(TypeArray)
	default:
		return "unknown"
	}
}

// NormalizeType normalizes supported type aliases to their canonical names.
func NormalizeType(typeName string) string {
	switch strings.ToLower(typeName) {
	case "int":
		return string(TypeInteger)
	case "float":
		return string(TypeNumber)
	case "bool":
		return string(TypeBoolean)
	case "text":
		return string(TypeString)
	case "list":
		return string(TypeArray)
	default:
		return strings.ToLower(typeName)
	}
}
