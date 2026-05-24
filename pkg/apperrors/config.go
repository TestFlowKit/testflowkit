package apperrors

import "errors"

// ErrNoConfigPath is returned when no configuration file path is provided.
var ErrNoConfigPath = errors.New("no config path provided")

// ErrNoModeProvided is returned when the CLI is invoked without a subcommand.
var ErrNoModeProvided = errors.New("no mode provided")

// ErrConfigNotLoaded is returned when Get() is called before Load().
var ErrConfigNotLoaded = errors.New("configuration not loaded - call Load first")

// ErrConfigAlreadyExists is returned when init is run and a config file already exists.
var ErrConfigAlreadyExists = errors.New("configuration file already exists")

// ErrFrontendElementsRequired is returned when frontend elements configuration is absent.
var ErrFrontendElementsRequired = errors.New("frontend elements configuration is required")

// ErrFrontendPagesRequired is returned when frontend pages configuration is absent.
var ErrFrontendPagesRequired = errors.New("frontend pages configuration is required")
