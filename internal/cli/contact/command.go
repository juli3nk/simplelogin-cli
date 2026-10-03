package contact

import (
	"github.com/spf13/cobra"
)

var (
	compact   bool
	noHeaders bool
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contact",
		Short: "Manage contacts",
		Long:  contactDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	cmd.AddCommand(
		newBlockCommand(),
		newCreateCommand(),
		newDeleteCommand(),
		newListCommand(),
	)

	return cmd
}

const contactDescription = `
The **simplelogin-cli contact** command has subcommands for managing contacts.

To see help for a subcommand, use:

    simplelogin-cli contact [command] --help

`
