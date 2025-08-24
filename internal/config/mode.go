package config

type Mode string

const (
	RunMode        Mode = "run"
	InitMode       Mode = "init"
	ValidationMode Mode = "validate"
	VersionMode    Mode = "version"
)
