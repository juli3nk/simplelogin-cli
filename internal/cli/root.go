package cli

import (
	"github.com/spf13/cobra"

	"github.com/juli3nk/simplelogin-cli/internal/cli/alias"
	"github.com/juli3nk/simplelogin-cli/internal/cli/auth"
	"github.com/juli3nk/simplelogin-cli/internal/cli/contact"
	"github.com/juli3nk/simplelogin-cli/internal/cli/domain"
	"github.com/juli3nk/simplelogin-cli/internal/cli/mailbox"
	"github.com/juli3nk/simplelogin-cli/internal/cli/setting"
	"github.com/juli3nk/simplelogin-cli/internal/cli/stats"
	"github.com/juli3nk/simplelogin-cli/internal/cli/userinfo"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "simplelogin-cli",
		Short: "SimpleLogin CLI",
		Long:  "Manage SimpleLogin aliases, domains and mailboxes from the command line.",
	}

	cmd.PersistentFlags().StringP("output", "o", "table", "Output format (table,json)")

	cmd.AddCommand(alias.NewCommand())
	cmd.AddCommand(auth.NewCommand())
	cmd.AddCommand(contact.NewCommand())
	cmd.AddCommand(domain.NewCommand())
	cmd.AddCommand(mailbox.NewCommand())
	cmd.AddCommand(setting.NewCommand())
	cmd.AddCommand(stats.NewCommand())
	cmd.AddCommand(userinfo.NewCommand())

	cmd.AddCommand(NewVersionCommand())

	return cmd
}

func Execute() error {
	return NewCommand().Execute()
}
