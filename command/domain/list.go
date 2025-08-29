package domain

import (
	"fmt"
	"log"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/juli3nk/simplelogin-cli/pkg/simplelogin"
	"github.com/spf13/cobra"
)

func newListCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List domains",
		Long:    listDescription,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runList(outputFormat)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Hide table headers")

	return cmd
}

func runList(outputFormat *string) {
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

	domains, err := client.GetDomains()
	if err != nil {
		log.Fatal(err)
	}

	switch *outputFormat {
	case "json":
		if len(domains) == 0 {
			fmt.Printf("%v\n", domains)
			return
		}

		if err := display.DisplayData(domains, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			log.Fatal(err)
		}
	default: // table
		if len(domains) == 0 {
			fmt.Println("No domains found.")
			return
		}

		tableOpts := display.DefaultTableOptions()
		if noHeaders {
			tableOpts.NoHeaders = true
		}

		table := display.NewTable(tableOpts)

		table.SetHeader([]string{"Domain", "ID", "Verified", "Nb Alias"})

		for _, domain := range domains {
			table.Append([]string{
				domain.DomainName,
				display.FormatID(domain.ID),
				display.FormatBool(domain.IsVerified),
				display.FormatID(domain.NbAlias),
			})
		}

		table.Render()
	}
}

const listDescription = `
List domains

`
