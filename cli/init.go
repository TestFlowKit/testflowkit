package cli

import (
	_ "embed"
	"etoolse/config"
	"log"
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

func Init(appConfig *config.AppConfig) {
	log.Println("Init cli config ...")

	if _, err := os.Stat("cli.yml"); err == nil {
		log.Fatal("cli already initialized")
	}

	if _, err := os.Stat("frontend.yml"); err == nil {
		log.Fatal("cli already initialized")
	}

	vars := cliConfigVars{
		AppName:        appConfig.AppName,
		AppVersion:     appConfig.AppVersion,
		AppDescription: appConfig.AppDescription,
	}

	tmpl, err := template.New("cli").Parse(cliConfigTemplate)
	if err != nil {
		log.Fatal("failed to parse cli config template: ", err)
	}

	f, err := os.Create("cli.yml")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, vars)
	if err != nil {
		f.Close()
		log.Fatal("failed to execute cli config template: ", err)
	}

	err = os.WriteFile("frontend.yml", []byte(frontTestsConfigTemplate), 0600)
	if err != nil {
		f.Close()
		os.Remove("cli.yml")
		log.Fatal("failed to write frontend config: ", err)
	}

	f.Close()
}
