package mailbox

import (
	"fmt"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newCreateCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create [email]",
		Aliases: []string{"c"},
		Short:   "Create mailbox",
		Long:    createDescription,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(cmd, outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runCreate(cmd *cobra.Command, outputFormat *string, args []string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	mailbox, err := client.CreateMailbox(args[0])
	if err != nil {
		return err
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(mailbox, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default:
		fmt.Printf("ID: %d\n", mailbox.ID)
		fmt.Printf("Email: %s\n", mailbox.Email)
		fmt.Printf("Verified: %t\n", mailbox.Verified)
		fmt.Printf("Default: %t\n", mailbox.Default)
	}

	return nil
}

const createDescription = `
Create mailbox

`
