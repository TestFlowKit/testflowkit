package configschema

const (
	schemaTypeObject  = "object"
	schemaTypeMap     = "map"
	schemaTypeArray   = "array"
	schemaTypeString  = "string"
	schemaTypeInteger = "integer"
	schemaTypeBoolean = "boolean"
	schemaTypeNumber  = "number"
	schemaTypeAny     = "any"
)

// Export is the top-level JSON document returned by ExportJSON.
type Export struct {
	Root    string `json:"root"`
	Version string `json:"version"`
	Schema  *Node  `json:"schema"`
}

// Node describes a single property in the testflowkit.yml schema tree.
type Node struct {
	Name            string           `json:"name,omitempty"`
	YAMLKey         string           `json:"yaml_key"`
	Type            string           `json:"type"`
	GoType          string           `json:"go_type,omitempty"`
	Description     string           `json:"description"`
	Details         string           `json:"details,omitempty"`
	TypeDescription string           `json:"type_description,omitempty"`
	Required        bool             `json:"required,omitempty"`
	Constraints     string           `json:"constraints,omitempty"`
	Enum            []string         `json:"enum,omitempty"`
	KeyType         string           `json:"key_type,omitempty"`
	Properties      map[string]*Node `json:"properties,omitempty"`
	Value           *Node            `json:"value,omitempty"`
	Items           *Node            `json:"items,omitempty"`
}

type docIndex struct {
	typeDocs  map[string]string
	fieldDocs map[string]map[string]string
}
