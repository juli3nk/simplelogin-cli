package mailbox

import (
	"fmt"
	"log"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/juli3nk/simplelogin-cli/pkg/simplelogin"
	"github.com/spf13/cobra"
)

func newListCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List mailboxes",
		Long:    listDescription,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runList(outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Hide table headers")

	return cmd
}

func runList(outputFormat *string, args []string) {
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

	mailboxes, err := client.GetMailboxes()
	if err != nil {
		log.Fatal(err)
	}

	switch *outputFormat {
	case "json":
		if len(mailboxes) == 0 {
			fmt.Printf("%v\n", mailboxes)
			return
		}

		if err := display.DisplayData(mailboxes, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			log.Fatal(err)
		}
	default: // table
		if len(mailboxes) == 0 {
			fmt.Println("No mailboxes found.")
			return
		}

		table := display.NewTable(nil)
		table.SetHeader([]string{"ID", "Email", "Default", "Verified", "Aliases", "Created"})

		for _, mailbox := range mailboxes {
			defaultStatus := ""
			if mailbox.Default {
				defaultStatus = "✓"
			}

			table.Append([]string{
				display.FormatID(mailbox.ID),
				display.FormatEmail(mailbox.Email, 40),
				defaultStatus,
				display.FormatBool(mailbox.Verified),
				display.FormatID(mailbox.NBAlias),
				display.FormatTimestamp(mailbox.CreationTimestamp),
			})
		}

		table.Render()
		fmt.Printf("\nTotal: %d mailboxes\n", len(mailboxes))
	}
}

const listDescription = `
List mailboxes

`
