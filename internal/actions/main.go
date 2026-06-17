package actions

import (
	"fmt"
	"testflowkit/internal/actions/actionexportconfigschema"
	"testflowkit/internal/actions/actionexportsteps"
	"testflowkit/internal/actions/actioninit"
	"testflowkit/internal/actions/actioninstall"
	"testflowkit/internal/actions/actionrun"
	"testflowkit/internal/actions/actionvalidate"
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
)

func Execute(cfg *config.Config, cfgErr error, mode config.Mode) {
	modes := map[config.Mode]func(*config.Config, error){
		config.RunMode:        actionrun.Execute,
		config.InitMode:       actioninit.Execute,
		config.InstallMode:    actioninstall.Execute,
		config.ValidationMode: actionvalidate.Execute,
		config.VersionMode:    version,
		config.ExportStepDefinitionsMode: func(_ *config.Config, _ error) {
			actionexportsteps.Execute("json")
		},
		config.ExportConfigSchemaMode: func(cfg *config.Config, _ error) {
			version := ""
			if cfg != nil {
				version = cfg.GetVersion()
			}
			actionexportconfigschema.Execute("json", version)
		},
	}

	if action, ok := modes[mode]; ok {
		action(cfg, cfgErr)
	} else {
		logger.Fatal(fmt.Sprintf("unknown mode: %s", mode), nil)
	}
}
