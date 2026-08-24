package setting

import (
	"fmt"

	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/juli3nk/simplelogin-cli/pkg/simplelogin"
	"github.com/spf13/cobra"
)

var (
	aliasGenerator           string
	notification             bool
	randomAliasDefaultDomain string
	senderFormat             string
	randomAliasSuffix        string
)

func newUpdateCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update setting",
		Long:  updateDescription,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpdate(cmd, outputFormat)
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&compact, "compact", false, "Compact output")

	flags.StringVarP(&aliasGenerator, "alias-generator", "a", "", "Alias generator")
	flags.BoolVarP(&notification, "notification", "n", false, "Notification")
	flags.StringVarP(&randomAliasDefaultDomain, "random-alias-default-domain", "d", "", "Random alias default domain")
	flags.StringVarP(&senderFormat, "sender-format", "s", "", "Sender format")
	flags.StringVarP(&randomAliasSuffix, "random-alias-suffix", "r", "", "Random alias suffix")

	return cmd
}

func runUpdate(cmd *cobra.Command, outputFormat *string) error {
	if aliasGenerator == "" && !notification && randomAliasDefaultDomain == "" && senderFormat == "" && randomAliasSuffix == "" {
		return fmt.Errorf("no update provided")
	}

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	settingInput := simplelogin.Setting{}
	if aliasGenerator != "" {
		settingInput.AliasGenerator = aliasGenerator
	}
	if notification {
		settingInput.Notification = notification
	}
	if randomAliasDefaultDomain != "" {
		settingInput.RandomAliasDefaultDomain = randomAliasDefaultDomain
	}
	if senderFormat != "" {
		settingInput.SenderFormat = senderFormat
	}
	if randomAliasSuffix != "" {
		settingInput.RandomAliasSuffix = randomAliasSuffix
	}

	setting, err := client.UpdateSetting(settingInput)
	if err != nil {
		return err
	}

	switch *outputFormat {
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

const updateDescription = `
Update setting

`
