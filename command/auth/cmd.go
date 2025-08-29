package auth

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Manage auth",
		Long:  authDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(
		newSetKeyCommand(),
	)

	return cmd
}

const authDescription = `
The **simplelogin-cli auth** command has subcommands for managing authentication.

To see help for a subcommand, use:

    simplelogin-cli auth [command] --help

For full details on using simplelogin visit SimpleLogin's online documentation.

`
