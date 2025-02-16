package actions

import (
	_ "embed"
	"etoolse/internal/config"
	"etoolse/pkg/logger"
	"os"
	"text/template"
)

//go:embed boilerplate/cli.boilerplate.yml
var cliConfigTemplate string

//go:embed boilerplate/frontend.boilerplate.yml
var frontTestsConfigTemplate string

type cliConfigVars struct {
	AppName        string
	AppVersion     string
	AppDescription string
}

func Init(appConfig *config.App) {
	logger.Info("init cmd config ...")

	if _, err := os.Stat("cmd.yml"); err == nil {
		logger.Fatal("cmd already initialized", err)
	}

	if _, err := os.Stat("frontend.yml"); err == nil {
		logger.Fatal("cmd already initialized", err)
	}

	vars := cliConfigVars{
		AppName:        appConfig.AppName,
		AppVersion:     appConfig.AppVersion,
		AppDescription: appConfig.AppDescription,
	}

	tmpl, err := template.New("cmd").Parse(cliConfigTemplate)
	if err != nil {
		logger.Fatal("failed to parse cmd config template: ", err)
	}

	f, err := os.Create("cmd.yml")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, vars)
	if err != nil {
		f.Close()
		logger.Fatal("failed to execute cmd config template: ", err)
	}

	err = os.WriteFile("frontend.yml", []byte(frontTestsConfigTemplate), 0600)
	if err != nil {
		f.Close()
		os.Remove("cmd.yml")
		logger.Fatal("failed to write frontend config: ", err)
	}

	f.Close()
}
