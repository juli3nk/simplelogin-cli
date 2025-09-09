package domain

import (
	"github.com/spf13/cobra"
)

var (
	compact   bool
	noHeaders bool
)

func NewCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "domain",
		Short: "Manage domains",
		Long:  domainDescription,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmd.AddCommand(
		newListCommand(outputFormat),
		newTrashCommand(outputFormat),
		newUpdateCommand(outputFormat),
	)

	return cmd
}

const domainDescription = `
The **simplelogin-cli domain** command has subcommands for managing domains.

To see help for a subcommand, use:

    simplelogin-cli domain [command] --help

`
