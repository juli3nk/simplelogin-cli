package contact

import (
	"fmt"
	"strconv"

	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create [alias_id] [contact_email]",
		Aliases: []string{"c"},
		Short:   "Create contact",
		Long:    createDescription,
		Args:    cobra.ExactArgs(2),
		RunE:    runCreate,
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runCreate(cmd *cobra.Command, args []string) error {
	outputFormat, err := cmd.Root().PersistentFlags().GetString("output")
	if err != nil {
		return err
	}

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	aliasID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	contact, err := client.CreateAliasContact(aliasID, args[1])
	if err != nil {
		return err
	}

	switch outputFormat {
	case "json":
		if err := display.DisplayData(contact, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default:
		fmt.Printf("Contact ID: %d\n", contact.ID)
		fmt.Printf("Contact email: %s\n", contact.Contact)
		fmt.Printf("Contact creation date: %s\n", contact.CreationDate)
		fmt.Printf("Contact creation timestamp: %d\n", contact.CreationTimestamp)
		fmt.Printf("Contact last email sent date: %s\n", contact.LastEmailSentDate)
		fmt.Printf("Contact last email sent timestamp: %d\n", contact.LastEmailSentTimestamp)
		fmt.Printf("Contact reverse alias: %s\n", contact.ReverseAlias)
		fmt.Printf("Contact block forward: %t\n", contact.BlockForward)
		fmt.Printf("Contact existed: %t\n", contact.Existed)
	}

	return nil
}

const createDescription = `
Create contact

`
