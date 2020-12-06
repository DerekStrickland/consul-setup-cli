package template

import (
	"testing"

	cliErrors "github.com/DerekStrickland/consul-setup-cli/errors"
	"github.com/mitchellh/cli"
	"github.com/stretchr/testify/require"
)

func TestDebugNotSpecifiedErr(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	ui := cli.NewMockUi()
	cmd := Command{
		UI: ui,
	}
	returnCode := cmd.Run([]string{})
	require.Equal(1, returnCode, ui.ErrorWriter.String())

	require.Contains(
		ui.ErrorWriter.String(),
		cliErrors.NewWithMessage("template.TestDebugNotSpecifiedErr", cliErrors.MissingArgumentError).Error(),
	)
}
