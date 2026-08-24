package contact

import (
	"fmt"
	"strconv"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newListCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list [alias_id]",
		Aliases: []string{"ls"},
		Short:   "List contacts",
		Long:    listDescription,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(cmd, outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Hide table headers")

	return cmd
}

func runList(cmd *cobra.Command, outputFormat *string, args []string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	aliasID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	contacts, err := client.GetAllAliasContacts(aliasID)
	if err != nil {
		return err
	}

	// Handle different output formats
	switch *outputFormat {
	case "json":
		if len(contacts) == 0 {
			fmt.Printf("%v\n", contacts)
			return nil
		}

		if err := display.DisplayData(contacts, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default: // table
		if len(contacts) == 0 {
			fmt.Println("No contacts found for this alias.")
			return nil
		}

		tableOpts := display.DefaultTableOptions()
		if noHeaders {
			tableOpts.NoHeaders = true
		}
		if compact {
			tableOpts = display.CompactTableOptions()
		}

		table := display.NewTable(tableOpts)
		if !noHeaders {
			table.SetHeader([]string{"ID", "Contact", "Created", "Last Email", "Reverse Alias", "Blocked"})
		}

		for _, contact := range contacts {
			table.Append([]string{
				display.FormatID(contact.ID),
				display.FormatEmail(contact.Contact, 30),
				display.FormatDate(contact.CreationDate),
				display.FormatDate(contact.LastEmailSentDate),
				display.FormatEmail(contact.ReverseAlias, 25),
				display.FormatBool(contact.BlockForward),
			})
		}

		table.Render()
		fmt.Printf("\nTotal: %d contacts\n", len(contacts))
	}

	return nil
}

const listDescription = `
List contacts

`
