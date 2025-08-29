package setting

import (
	"github.com/spf13/cobra"
)

var (
	compact   bool
	noHeaders bool
)

func NewCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setting",
		Short: "Manage settings",
		Long:  settingDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(
		newGetDomainsCommand(outputFormat),
		newGetCommand(outputFormat),
		newUpdateCommand(outputFormat),
	)

	return cmd
}

const settingDescription = `
The **simplelogin-cli setting** command has subcommands for managing settings.

To see help for a subcommand, use:

    simplelogin-cli setting [command] --help

For full details on using simplelogin visit SimpleLogin's online documentation.

`
