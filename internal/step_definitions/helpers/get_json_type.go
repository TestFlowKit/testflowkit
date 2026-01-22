package helpers

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

// GetJSONType returns the JSON type of a value.
func GetJSONType(value any) string {
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

func NormalizeType(typeName string) string {
	switch typeName {
	case "int":
		return string(TypeInteger)
	case "float":
		return string(TypeNumber)
	case "bool":
		return string(TypeBoolean)
	default:
		return typeName
	}
}
