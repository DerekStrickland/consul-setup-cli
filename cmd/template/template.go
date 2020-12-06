package template

import (
	"flag"
	"sync"

	cliErrors "github.com/DerekStrickland/consul-setup-cli/errors"

	"github.com/hashicorp/consul/command/flags"
	"github.com/mitchellh/cli"
)

// Command encapsulates this operation.
type Command struct {
	UI        cli.Ui
	flags     *flag.FlagSet
	flagDebug bool
	once      sync.Once
	help      string
}

func (c *Command) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)

	c.flags.BoolVar(&c.flagDebug, "debug", false,
		"Show debug output while processing")

	c.help = flags.Usage(help, c.flags)
}

// Run fulfills the command.
func (c *Command) Run(args []string) int {
	c.once.Do(c.init)
	var err error
	if err = c.flags.Parse(args); err != nil {
		return 1
	}

	if len(c.flags.Args()) > 0 {
		c.UI.Error(cliErrors.NonFlagArgumentsError)
		return 1
	}

	err = c.applyTemplate()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	return 0
}

func (c *Command) applyTemplate() error {
	if 1 == 1 {
		return cliErrors.NewWithMessage("template.applyTemplate", cliErrors.NotImplementedError)
	}
	return nil
}

// Synopsis returns a short description of the command
func (c *Command) Synopsis() string { return synopsis }

// Help returns the full CLI help text
func (c *Command) Help() string {
	c.once.Do(c.init)
	return c.help
}

const synopsis = `Run the consul setup wizard template functions.

Users can:
	- Install from a saved template
	- Install from a built in template (demo, insecure, secure)
	- Install add-on components from a template to the target installation (Prometheus, Grafan, Jaeger)
.`
const help = `
Usage: consul-setup-cli template [options]

  Runs applies template to current context.
`
