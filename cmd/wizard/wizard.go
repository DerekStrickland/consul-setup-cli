package wizard

import (
	"flag"
	"fmt"
	"reflect"
	"sync"

	cliErrors "github.com/DerekStrickland/consul-setup-cli/errors"
	proto "github.com/DerekStrickland/consul-setup-cli/prototype"
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

	output, err := c.start()
	if err != nil {
		c.UI.Error(cliErrors.NewWithError("wizard.Run.start", err).Error())
		return 1
	}

	c.UI.Info(fmt.Sprintf("output: %+v", output))

	result, err := yaml.Marshal(&output)
	if err != nil {
		c.UI.Error(cliErrors.NewWithError("wizard.Run.yaml.Marshal", err).Error())
		return 1
	}

	fmt.Printf("--- t dump:\n%+v\n\n", string(result))

	c.UI.Info(string(result))

	return 0
}

func (c *Command) start() (*proto.ConsulHelmValues, error) {
	//return cliErrors.NewWithMessage("wizard.startWizard", cliErrors.NotImplementedError)
	prototype, err := proto.NewPrototype()
	output := proto.ConsulHelmValues{}
	if err != nil {
		return &prototype, cliErrors.NewWithError("wizard.start", err)
	}

	fieldNames := reflection.GetFieldNames(prototype)
	for _, fieldName := range fieldNames {
		query := fmt.Sprintf("Add a %s stanza? (Y/n)", fieldName)
		addStanza, err := c.UI.Ask(query)
		if err != nil {
			return &prototype, err
		}
		if addStanza == "Y" {
			stanzaPrototype, err := reflection.GetStructField(prototype, fieldName)
			if err != nil {
				return &prototype, err
			}

			err = c.buildStanza(output, fieldName, stanzaPrototype)
			if err != nil {
				return &prototype, err
			}

			c.UI.Info(fmt.Sprintf("wizard.Command.start.buildStanza: %+v\n\n", output))
		}
	}

	c.UI.Info(fmt.Sprintf("start: Result %+v", output))
	return &output, nil
}

func (c *Command) buildStanza(target interface{}, fieldName string, prototype interface{}) error {
	prototypeType := reflect.TypeOf(prototype)
	targetStanza := reflect.New(prototypeType).Elem()
	c.UI.Info(fmt.Sprintf("targetStanza is %+v\n\n", targetStanza))

	err := reflection.SetField(targetStanza, "Enabled", true)
	if err != nil {
		return err
	}

	c.UI.Info(fmt.Sprintf("wizard.Command.buildStanza: %+v\n\n", targetStanza))

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
