package wizard

import (
	"flag"
	"fmt"
	"sync"

	proto "github.com/DerekStrickland/consul-setup-cli/prototype"
	cliErrors "github.com/DerekStrickland/consul-setup-cli/errors"
	reflection "github.com/DerekStrickland/consul-setup-cli/reflection"

	"github.com/hashicorp/consul/command/flags"
	"github.com/mitchellh/cli"

	yaml "gopkg.in/yaml.v2"
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

	output, err = c.start()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	result, err := yaml.Marshal(&output)
	if err != nil {
		c.UI.Error(cliErrors.NewWithError("wizard.Run.yaml.Marshal", err))
		return 1
	}
	c.UI.Info(string(result))

	return 0
}

func (c *Command) start() (*proto.YamlValues, error) {
	//return cliErrors.NewWithMessage("wizard.startWizard", cliErrors.NotImplementedError)
	prototype, err := proto.NewPrototype()
	output := proto.YamlValues{}
	if err != nil {
		return prototype, cliErrors.NewWithError("wizard.start", err)
	}

	fields := reflection.GetFields(prototype)
	for _, field := range fields {
		addStanza := c.UI.Ask(fmt.Sprintf("Add a %s stanza? (Y/n)"), field.Name))
		if addStanza == "Y" {
			proto.AddDefaultField(output, field.Name)
		}
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

const synopsis = `Run the consul setup interactive wizard that prompts (and infrorms)
	the user about their setup choices when generating the helm values file.`
const help = `
Usage: consul-setup-cli wizard [options]

  Runs setup wizard.
`
