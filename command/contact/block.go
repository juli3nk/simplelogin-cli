package contact

import (
	"fmt"
	"strconv"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newBlockCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "block [contact_id]",
		Aliases: []string{"b"},
		Short:   "Block contact",
		Long:    blockDescription,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runBlock(cmd, outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runBlock(cmd *cobra.Command, outputFormat *string, args []string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	contactID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	contact, err := client.ToggleContact(contactID)
	if err != nil {
		return err
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(contact, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default:
		fmt.Printf("Contact blocked: %t\n", contact.BlockForward)
	}

	return nil
}

const blockDescription = `
Block contact

`
