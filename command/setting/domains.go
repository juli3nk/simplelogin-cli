package setting

import (
	"fmt"
	"log"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/juli3nk/simplelogin-cli/pkg/simplelogin"
	"github.com/spf13/cobra"
)

func newGetDomainsCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-domains",
		Short: "Get setting domains",
		Long:  getDomainsDescription,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runGetDomains(outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Hide table headers")

	return cmd
}

func runGetDomains(outputFormat *string, args []string) {
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

	domains, err := client.GetSettingDomains()
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
}

const getDomainsDescription = `
Get setting domains

`
