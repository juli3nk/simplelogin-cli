package mailbox

import (
	"fmt"
	"strconv"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/juli3nk/simplelogin-cli/pkg/simplelogin"
	"github.com/spf13/cobra"
)

var (
	transferAliasesTo string
)

func newDeleteCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete [mailbox_id]",
		Aliases: []string{"del"},
		Short:   "Delete mailbox",
		Long:    deleteDescription,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDelete(cmd, outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	cmd.Flags().StringVarP(&transferAliasesTo, "transfer-aliases-to", "t", "", "Transfer aliases to mailbox email")

	return cmd
}

func runDelete(cmd *cobra.Command, outputFormat *string, args []string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	mailboxID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	mailboxDeleteOptions := simplelogin.MailboxDeleteOptions{}
	if transferAliasesTo != "" {
		transferAliasesToID, err := strconv.Atoi(transferAliasesTo)
		if err != nil {
			return err
		}
		mailboxDeleteOptions.TransferAliasesTo = &transferAliasesToID
	}

	err = client.DeleteMailbox(mailboxID, mailboxDeleteOptions)
	if err != nil {
		return err
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(map[string]bool{"deleted": true}, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default:
		fmt.Printf("Mailbox deleted\n")
	}

	return nil
}

const deleteDescription = `
Delete mailbox

`
