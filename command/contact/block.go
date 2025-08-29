package contact

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

func newBlockCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "block [contact_id]",
		Aliases: []string{"b"},
		Short:   "Block contact",
		Long:    blockDescription,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runBlock(outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runBlock(outputFormat *string, args []string) {
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

	contactID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	contact, err := client.ToggleContact(contactID)
	if err != nil {
		log.Fatal(err)
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(contact, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("Contact blocked: %t\n", contact.BlockForward)
	}
}

const blockDescription = `
Block contact

`
