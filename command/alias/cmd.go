package alias

import (
	"github.com/spf13/cobra"
)

var (
	compact   bool
	noHeaders bool
)

func NewCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alias",
		Short: "Manage aliases",
		Long:  aliasDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(
		newActivitiesCommand(outputFormat),
		newCreateNewCommand(outputFormat),
		newCreateRandomCommand(outputFormat),
		newDeleteCommand(outputFormat),
		newGetCommand(outputFormat),
		newListCommand(outputFormat),
		newOptionsCommand(outputFormat),
		newToggleCommand(outputFormat),
		newUpdateCommand(),
	)

	return cmd
}

const aliasDescription = `
The **simplelogin-cli alias** command has subcommands for managing aliases.

To see help for a subcommand, use:

    simplelogin-cli alias [command] --help

For full details on using simplelogin visit SimpleLogin's online documentation.

`
