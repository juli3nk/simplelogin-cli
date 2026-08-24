package domain

import (
	"fmt"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newListCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List domains",
		Long:    listDescription,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(cmd, outputFormat)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Hide table headers")

	return cmd
}

func runList(cmd *cobra.Command, outputFormat *string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	domains, err := client.GetDomains()
	if err != nil {
		return err
	}

	switch *outputFormat {
	case "json":
		if len(domains) == 0 {
			fmt.Printf("%v\n", domains)
			return nil
		}

		if err := display.DisplayData(domains, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default: // table
		if len(domains) == 0 {
			fmt.Println("No domains found.")
			return nil
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

	return nil
}

const listDescription = `
List domains

`
