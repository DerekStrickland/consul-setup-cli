package main

import (
	"os"

	cmdTemplate "github.com/DerekStrickland/consul-setup-cli/cmd/template"
	cmdWizard "github.com/DerekStrickland/consul-setup-cli/cmd/wizard"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all available consul-k8s commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}

	coloredUI := &cli.ColoredUi{
		OutputColor: cli.UiColorMagenta,
		InfoColor:   cli.UiColorBlue,
		ErrorColor:  cli.UiColorRed,
		WarnColor:   cli.UiColorYellow,
		Ui:          ui,
	}

	Commands = map[string]cli.CommandFactory{
		"wizard": func() (cli.Command, error) {
			return &cmdWizard.Command{UI: coloredUI}, nil
		},
		"template": func() (cli.Command, error) {
			return &cmdTemplate.Command{UI: coloredUI}, nil
		},
	}
}

func helpFunc() cli.HelpFunc {
	var keys []string
	for key := range Commands {
		keys = append(keys, key)
	}
	return cli.FilteredHelpFunc(keys, cli.BasicHelpFunc("lnx"))
}
