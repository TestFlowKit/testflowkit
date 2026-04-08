package queryable

// Result is the normalized representation of a value extracted from a response
// body, regardless of whether the source format is JSON or XML.
type Result struct {
	Exists bool
	Raw    string
	Kind   string
	Value  any
}

// Queryable exposes a format-agnostic query API for response assertions.
type Queryable interface {
	Get(path string) (Result, error)
	GetAll(path string) ([]Result, error)
	Exists(path string) (bool, error)
}
