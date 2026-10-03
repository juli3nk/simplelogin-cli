package setting

import (
	"fmt"

	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get setting",
		Long:  getDescription,
		Args:  cobra.NoArgs,
		RunE:  runGet,
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runGet(cmd *cobra.Command, args []string) error {
	outputFormat, err := cmd.Root().PersistentFlags().GetString("output")
	if err != nil {
		return err
	}

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	setting, err := client.GetSetting()
	if err != nil {
		return err
	}

	switch outputFormat {
	case "json":
		if err := display.DisplayData(setting, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default:
		fmt.Printf("Alias Generator: %s\n", setting.AliasGenerator)
		fmt.Printf("Notification: %t\n", setting.Notification)
		fmt.Printf("Random Alias Default Domain: %s\n", setting.RandomAliasDefaultDomain)
		fmt.Printf("Sender Format: %s\n", setting.SenderFormat)
		fmt.Printf("Random Alias Suffix: %s\n", setting.RandomAliasSuffix)
	}

	return nil
}

const getDescription = `
Get setting

`
