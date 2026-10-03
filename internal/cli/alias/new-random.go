package alias

import (
	"fmt"

	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

var (
	createRandomMode string
	createRandomNote string
)

func newCreateRandomCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "random [hostname]",
		Short: "Create random alias",
		Long:  createRandomDescription,
		Args:  cobra.ExactArgs(1),
		RunE:  runCreateRandom,
	}

	flags := cmd.Flags()
	flags.BoolVar(&compact, "compact", false, "Compact output")

	flags.StringVarP(&createRandomMode, "mode", "m", "", "The mode of the alias")
	flags.StringVar(&createRandomNote, "note", "", "Alias note")

	return cmd
}

func runCreateRandom(cmd *cobra.Command, args []string) error {
	outputFormat, err := cmd.Root().PersistentFlags().GetString("output")
	if err != nil {
		return err
	}

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	hostname := args[0]

	alias, err := client.CreateRandomAlias(hostname, createRandomMode, createRandomNote)
	if err != nil {
		return err
	}

	switch outputFormat {
	case "json":
		if err := display.DisplayData(alias, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default:
		fmt.Printf("CreationDate: %s\n", alias.CreationDate)
		fmt.Printf("CreationTimestamp: %d\n", alias.CreationTimestamp)
		fmt.Printf("Email: %s\n", alias.Email)
		fmt.Printf("Name: %s\n", alias.Name)
		fmt.Printf("Enabled: %t\n", alias.Enabled)
		fmt.Printf("ID: %d\n", alias.ID)
		fmt.Printf("Mailbox: %+v\n", alias.Mailbox)
		for _, mailbox := range alias.Mailboxes {
			fmt.Printf("Mailbox: %+v\n", mailbox)
		}
		fmt.Printf("LatestActivity: %+v\n", alias.LatestActivity)
		fmt.Printf("NbBlock: %d\n", alias.NbBlock)
		fmt.Printf("NbForward: %d\n", alias.NbForward)
		fmt.Printf("NbReply: %d\n", alias.NbReply)
		fmt.Printf("Note: %s\n", alias.Note)
		fmt.Printf("Pinned: %t\n", alias.Pinned)
	}

	return nil
}

const createRandomDescription = `
Create random alias

`
