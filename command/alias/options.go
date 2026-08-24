package alias

import (
	"fmt"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newOptionsCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "options [hostname]",
		Short: "Get alias options",
		Long:  optionsDescription,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runOptions(cmd, outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runOptions(cmd *cobra.Command, outputFormat *string, args []string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	aliasOptions, err := client.GetAliasOptions(args[0])
	if err != nil {
		return err
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(aliasOptions, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
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

	return nil
}

const optionsDescription = `
Get alias options

`
