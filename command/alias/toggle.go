package alias

import (
	"fmt"
	"strconv"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newToggleCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "toggle [alias_id]",
		Aliases: []string{"t"},
		Short:   "Toggle alias",
		Long:    toggleDescription,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runToggle(cmd, outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runToggle(cmd *cobra.Command, outputFormat *string, args []string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	aliasID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	alias, err := client.ToggleAlias(aliasID)
	if err != nil {
		return err
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(alias, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default:
		fmt.Printf("Alias enabled: %t\n", alias.Enabled)
	}

	return nil
}

const toggleDescription = `
Toggle alias

`
