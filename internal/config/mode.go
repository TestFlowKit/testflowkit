package config

type Mode string

const (
	RunMode        Mode = "run"
	InitMode       Mode = "init"
	InstallMode    Mode = "install"
	ValidationMode Mode = "validate"
	VersionMode    Mode = "version"
)
