package mailbox

import (
	"github.com/spf13/cobra"
)

var (
	compact   bool
	noHeaders bool
)

func NewCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mailbox",
		Short: "Manage mailboxes",
		Long:  mailboxDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	cmd.AddCommand(
		newCreateCommand(outputFormat),
		newDeleteCommand(outputFormat),
		newListCommand(outputFormat),
	)

	return cmd
}

const mailboxDescription = `
The **simplelogin-cli mailbox** command has subcommands for managing mailboxes.

To see help for a subcommand, use:

    simplelogin-cli mailbox [command] --help

`
