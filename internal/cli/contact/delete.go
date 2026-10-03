package contact

import (
	"fmt"
	"strconv"

	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete [contact_id]",
		Aliases: []string{"del"},
		Short:   "Delete contact",
		Long:    deleteDescription,
		Args:    cobra.ExactArgs(1),
		RunE:    runDelete,
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runDelete(cmd *cobra.Command, args []string) error {
	outputFormat, err := cmd.Root().PersistentFlags().GetString("output")
	if err != nil {
		return err
	}

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	contactID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	contact, err := client.DeleteContact(contactID)
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
		fmt.Printf("Contact deleted: %t\n", contact.Deleted)
	}

	return nil
}

const deleteDescription = `
Delete contact

`
