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

func newGetCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get setting",
		Long:  getDescription,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runGet(outputFormat)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runGet(outputFormat *string) {
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

	setting, err := client.GetSetting()
	if err != nil {
		log.Fatal(err)
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(setting, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("Alias Generator: %s\n", setting.AliasGenerator)
		fmt.Printf("Notification: %t\n", setting.Notification)
		fmt.Printf("Random Alias Default Domain: %s\n", setting.RandomAliasDefaultDomain)
		fmt.Printf("Sender Format: %s\n", setting.SenderFormat)
		fmt.Printf("Random Alias Suffix: %s\n", setting.RandomAliasSuffix)
	}
}

const getDescription = `
Get setting

`
