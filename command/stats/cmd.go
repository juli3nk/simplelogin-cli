package stats

import (
	"fmt"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

var (
	compact bool
)

func NewCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stats",
		Short: "Manage stats",
		Long:  statsDescription,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStats(cmd, outputFormat)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runStats(cmd *cobra.Command, outputFormat *string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	stats, err := client.GetStats()
	if err != nil {
		return err
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(stats, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default:
		fmt.Printf("Nb Alias: %d\n", stats.NBAlias)
		fmt.Printf("Nb Block: %d\n", stats.NBBlock)
		fmt.Printf("Nb Forward: %d\n", stats.NBForward)
		fmt.Printf("Nb Reply: %d\n", stats.NBReply)
	}

	return nil
}

const statsDescription = `
Display account statistics.

The **simplelogin-cli stats** command shows the number of aliases, blocked,
forwarded and replied emails for the authenticated account.

Use --output json to get the result as JSON.
`
