package reporters

// Cucumber JSON report schema: a list of features, each with scenario "elements".

type cucumberFeature struct {
	URI      string            `json:"uri"`
	ID       string            `json:"id"`
	Keyword  string            `json:"keyword"`
	Name     string            `json:"name"`
	Elements []cucumberElement `json:"elements"`
}

type cucumberElement struct {
	ID      string         `json:"id"`
	Keyword string         `json:"keyword"`
	Type    string         `json:"type"`
	Name    string         `json:"name"`
	Tags    []cucumberTag  `json:"tags,omitempty"`
	Steps   []cucumberStep `json:"steps"`
}

type cucumberTag struct {
	Name string `json:"name"`
}

type cucumberStep struct {
	Keyword    string              `json:"keyword"`
	Name       string              `json:"name"`
	Result     cucumberResult      `json:"result"`
	Embeddings []cucumberEmbedding `json:"embeddings,omitempty"`
}

type cucumberResult struct {
	Status       string `json:"status"`
	Duration     int64  `json:"duration"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type cucumberEmbedding struct {
	MimeType string `json:"mime_type"`
	Data     string `json:"data"`
}
