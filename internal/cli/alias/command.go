package alias

import (
	"github.com/spf13/cobra"
)

var (
	compact   bool
	noHeaders bool
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alias",
		Short: "Manage aliases",
		Long:  aliasDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	cmd.AddCommand(
		newActivitiesCommand(),
		newCreateNewCommand(),
		newCreateRandomCommand(),
		newDeleteCommand(),
		newGetCommand(),
		newListCommand(),
		newOptionsCommand(),
		newToggleCommand(),
		newUpdateCommand(),
	)

	return cmd
}

const aliasDescription = `
The **simplelogin-cli alias** command has subcommands for managing aliases.

To see help for a subcommand, use:

    simplelogin-cli alias [command] --help

`
