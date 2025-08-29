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

func newCreateCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create [email]",
		Aliases: []string{"c"},
		Short:   "Create mailbox",
		Long:    createDescription,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runCreate(outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runCreate(outputFormat *string, args []string) {
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

	mailbox, err := client.CreateMailbox(args[0])
	if err != nil {
		log.Fatal(err)
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(mailbox, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("ID: %d\n", mailbox.ID)
		fmt.Printf("Email: %s\n", mailbox.Email)
		fmt.Printf("Verified: %t\n", mailbox.Verified)
		fmt.Printf("Default: %t\n", mailbox.Default)
	}
}

const createDescription = `
Create mailbox

`
