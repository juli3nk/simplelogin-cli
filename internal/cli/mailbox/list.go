package mailbox

import (
	"fmt"

	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List mailboxes",
		Long:    listDescription,
		Args:    cobra.NoArgs,
		RunE:    runList,
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Hide table headers")

	return cmd
}

func runList(cmd *cobra.Command, args []string) error {
	outputFormat, err := cmd.Root().PersistentFlags().GetString("output")
	if err != nil {
		return err
	}

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	mailboxes, err := client.GetMailboxes()
	if err != nil {
		return err
	}

	switch outputFormat {
	case "json":
		if len(mailboxes) == 0 {
			fmt.Printf("%v\n", mailboxes)
			return nil
		}

		if err := display.DisplayData(mailboxes, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default: // table
		if len(mailboxes) == 0 {
			fmt.Println("No mailboxes found.")
			return nil
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

	return nil
}

const listDescription = `
List mailboxes

`
