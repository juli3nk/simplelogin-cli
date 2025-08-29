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

func newCreateCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create [alias_id] [contact_email]",
		Aliases: []string{"c"},
		Short:   "Create contact",
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

	aliasID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	contact, err := client.CreateAliasContact(aliasID, args[1])
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
}

const createDescription = `
Create contact

`
