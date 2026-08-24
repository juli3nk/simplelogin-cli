package alias

import (
	"fmt"
	"strconv"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newGetCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [name]",
		Short: "Get an alias",
		Long:  getDescription,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(cmd, outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runGet(cmd *cobra.Command, outputFormat *string, args []string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	aliasID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	alias, err := client.GetAlias(aliasID)
	if err != nil {
		return err
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(alias, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default:
		fmt.Printf("Creation Date: %s\n", alias.CreationDate)
		fmt.Printf("Creation Timestamp: %d\n", alias.CreationTimestamp)
		fmt.Printf("Email: %s\n", alias.Email)
		fmt.Printf("Name: %s\n", alias.Name)
		fmt.Printf("Enabled: %t\n", alias.Enabled)
		fmt.Printf("ID: %d\n", alias.ID)
		fmt.Printf("Mailbox: %+v\n", alias.Mailbox)
		fmt.Printf("Mailboxes: %+v\n", alias.Mailboxes)
		fmt.Printf("Latest Activity: %+v\n", alias.LatestActivity)
		fmt.Printf("Nb Block: %d\n", alias.NbBlock)
		fmt.Printf("Nb Forward: %d\n", alias.NbForward)
		fmt.Printf("Nb Reply: %d\n", alias.NbReply)
		fmt.Printf("Note: %s\n", alias.Note)
		fmt.Printf("Pinned: %t\n", alias.Pinned)
	}

	return nil
}

const getDescription = `
Get an alias

`
