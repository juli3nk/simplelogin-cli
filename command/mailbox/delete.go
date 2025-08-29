package mailbox

import (
	"fmt"
	"log"
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
		Run: func(cmd *cobra.Command, args []string) {
			runDelete(outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	cmd.Flags().StringVarP(&transferAliasesTo, "transfer-aliases-to", "t", "", "Transfer aliases to mailbox email")

	return cmd
}

func runDelete(outputFormat *string, args []string) {
	defer utils.RecoverFunc()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	apiKey, err := config.LoadApiKey()
	if err != nil {
		log.Fatal(err)
	}

	client, err := simplelogin.NewClient(cfg.ApiURL, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	mailboxID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	transferAliasesToID, err := strconv.Atoi(transferAliasesTo)
	if err != nil {
		log.Fatal(err)
	}

	mailboxDeleteOptions := simplelogin.MailboxDeleteOptions{
		TransferAliasesTo: &transferAliasesToID,
	}

	err = client.DeleteMailbox(mailboxID, mailboxDeleteOptions)
	if err != nil {
		log.Fatal(err)
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(map[string]bool{"deleted": true}, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("Mailbox deleted\n")
	}
}

const deleteDescription = `
Delete mailbox

`
