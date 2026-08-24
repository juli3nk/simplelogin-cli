package setting

import (
	"fmt"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newGetDomainsCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-domains",
		Short: "Get setting domains",
		Long:  getDomainsDescription,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetDomains(cmd, outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Hide table headers")

	return cmd
}

func runGetDomains(cmd *cobra.Command, outputFormat *string, args []string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	domains, err := client.GetSettingDomains()
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

		table := display.NewTable(nil)
		table.SetHeader([]string{"Domain", "IsCustom"})

		for _, domain := range domains {
			table.Append([]string{
				domain.Domain,
				fmt.Sprintf("%t", domain.IsCustom),
			})
		}

		table.Render()
		fmt.Printf("\nTotal: %d domains\n", len(domains))
	}

	return nil
}

const getDomainsDescription = `
Get setting domains

`
