package domain

import (
	"fmt"
	"strconv"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/juli3nk/simplelogin-cli/pkg/simplelogin"
	"github.com/spf13/cobra"
)

var (
	catchAll               bool
	randomPrefixGeneration bool
	name                   string
	mailboxIds             []int
)

func newUpdateCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update [domain_id]",
		Aliases: []string{"up"},
		Short:   "Update domain",
		Long:    updateDescription,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpdate(cmd, outputFormat, args)
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&compact, "compact", false, "Compact output")

	flags.BoolVarP(&catchAll, "catch-all", "c", false, "Catch all")
	flags.BoolVarP(&randomPrefixGeneration, "random-prefix-generation", "r", false, "Random prefix generation")
	flags.StringVarP(&name, "name", "n", "", "Name")
	flags.IntSliceVarP(&mailboxIds, "mailbox-ids", "m", []int{}, "Mailbox IDs")

	return cmd
}

func runUpdate(cmd *cobra.Command, outputFormat *string, args []string) error {
	defer utils.RecoverFunc()

	if !catchAll && !randomPrefixGeneration && name == "" && len(mailboxIds) == 0 {
		return fmt.Errorf("no update provided")
	}

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	domainID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	domainInput := simplelogin.UpdateDomain{}
	if catchAll {
		domainInput.CatchAll = true
	}
	if randomPrefixGeneration {
		domainInput.RandomPrefixGeneration = true
	}
	if name != "" {
		domainInput.Name = name
	}
	if len(mailboxIds) > 0 {
		domainInput.MailboxIds = mailboxIds
	}

	domain, err := client.UpdateDomain(domainID, domainInput)
	if err != nil {
		return err
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(domain, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default:
		fmt.Printf("CatchAll: %t\n", domain.CatchAll)
		fmt.Printf("CreationDate: %s\n", domain.CreationDate)
		fmt.Printf("CreationTimestamp: %d\n", domain.CreationTimestamp)
		fmt.Printf("DomainName: %s\n", domain.DomainName)
		fmt.Printf("ID: %d\n", domain.ID)
		fmt.Printf("IsVerified: %t\n", domain.IsVerified)
		fmt.Printf("Mailboxes: %v\n", domain.Mailboxes)
		fmt.Printf("Name: %s\n", domain.Name)
		fmt.Printf("NbAlias: %d\n", domain.NbAlias)
		fmt.Printf("RandomPrefixGeneration: %t\n", domain.RandomPrefixGeneration)
	}

	return nil
}

const updateDescription = `
Update domain

`
