package userinfo

import (
	"github.com/spf13/cobra"
)

var compact bool

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "userinfo",
		Short: "Manage user info",
		Long:  userinfoDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	cmd.AddCommand(
		newGetCommand(),
		newUpdateCommand(),
	)

	return cmd
}

const userinfoDescription = `
The **simplelogin-cli userinfo** command has subcommands for managing user info.

To see help for a subcommand, use:

    simplelogin-cli userinfo [command] --help

`
