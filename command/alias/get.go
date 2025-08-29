package alias

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

func newGetCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [name]",
		Short: "Get an alias",
		Long:  getDescription,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runGet(outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runGet(outputFormat *string, args []string) {
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

	aliasID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	alias, err := client.GetAlias(aliasID)
	if err != nil {
		log.Fatal(err)
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(alias, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			log.Fatal(err)
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
}

const getDescription = `
Get an alias

`
