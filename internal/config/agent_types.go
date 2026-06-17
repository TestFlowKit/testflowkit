package config

// AgentStepCatalog overrides how the MCP server resolves the step catalog.
type AgentStepCatalog struct {
	// File is a local step catalog JSON path (reserved; CLI export is used today).
	File string `yaml:"file"`

	// URL is a remote step catalog JSON URL (reserved; CLI export is used today).
	URL string `yaml:"url"`
}

// AgentConfig holds MCP IDE agent settings; ignored by tkit run.
type AgentConfig struct {
	// DefaultTagsForDraft is added to every AI-drafted scenario, e.g. "@wip @ai-generated".
	DefaultTagsForDraft string `yaml:"default_tags_for_draft"`

	// RunCommand is the suggested command to run draft tests (informational only).
	RunCommand string `yaml:"run_command"`

	// StepCatalog overrides step catalog resolution; omit to use the tkit CLI export.
	StepCatalog *AgentStepCatalog `yaml:"step_catalog"`
}
