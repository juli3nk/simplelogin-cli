package alias

import (
	"fmt"
	"log"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/juli3nk/simplelogin-cli/pkg/simplelogin"
	"github.com/spf13/cobra"
)

func newOptionsCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "options [hostname]",
		Short: "Get alias options",
		Long:  optionsDescription,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runOptions(outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runOptions(outputFormat *string, args []string) {
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

	aliasOptions, err := client.GetAliasOptions(args[0])
	if err != nil {
		log.Fatal(err)
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(aliasOptions, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("Can Create: %t\n", aliasOptions.CanCreate)
		fmt.Printf("Prefix Suggestion: %s\n", aliasOptions.PrefixSuggestion)
		for _, suffix := range aliasOptions.Suffixes {
			fmt.Printf("Suffix: %s\n", suffix.Suffix)
			fmt.Printf("Signed Suffix: %s\n", suffix.SignedSuffix)
			fmt.Printf("Is Custom: %t\n", suffix.IsCustom)
			fmt.Printf("Is Premium: %t\n", suffix.IsPremium)
		}
	}
}

const optionsDescription = `
Get alias options

`
