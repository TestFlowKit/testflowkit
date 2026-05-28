package main

// Schema is the top-level envelope written to config-schema.json.
type Schema struct {
	SchemaVersion string             `json:"schema_version"`
	GeneratedBy   string             `json:"generated_by"`
	Files         map[string]FileDoc `json:"files"`
}

// FileDoc describes one configuration file (testflowkit.yml, testflowkit.agent.yml ...).
type FileDoc struct {
	Description   string    `json:"description"`
	Interpolation string    `json:"interpolation,omitempty"`
	Sections      []Section `json:"sections"`
}

// Section is a top-level YAML key with its child field definitions.
// Polymorphic sections (e.g. security_schemes) carry Variants instead of Fields.
type Section struct {
	Key           string    `json:"key"`
	Required      bool      `json:"required"`
	Description   string    `json:"description"`
	Type          string    `json:"type,omitempty"`
	Discriminator string    `json:"discriminator,omitempty"`
	Fields        []Field   `json:"fields,omitempty"`
	Variants      []Variant `json:"variants,omitempty"`
}

// Field describes a single YAML field inside a section or nested object.
type Field struct {
	Key           string      `json:"key"`
	Type          string      `json:"type"`
	Required      bool        `json:"required,omitempty"`
	RequiredIf    string      `json:"required_if,omitempty"`
	Default       interface{} `json:"default,omitempty"`
	Enum          []string    `json:"enum,omitempty"`
	Min           interface{} `json:"min,omitempty"`
	Max           interface{} `json:"max,omitempty"`
	Interpolation bool        `json:"interpolation,omitempty"`
	Description   string      `json:"description"`
	Fields        []Field     `json:"fields,omitempty"`
}

// Variant describes the fields active for one value of the discriminator.
type Variant struct {
	When   string  `json:"when"`
	Fields []Field `json:"fields"`
}
