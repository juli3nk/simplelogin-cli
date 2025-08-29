package alias

import (
	"fmt"
	"log"
	"strconv"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/juli3nk/simplelogin-cli/pkg/simplelogin"
	"github.com/spf13/cobra"
)

func newToggleCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "toggle [alias_id]",
		Aliases: []string{"t"},
		Short:   "Toggle alias",
		Long:    toggleDescription,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runToggle(outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runToggle(outputFormat *string, args []string) {
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

	aliasID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	alias, err := client.ToggleAlias(aliasID)
	if err != nil {
		log.Fatal(err)
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(alias, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("Alias enabled: %t\n", alias.Enabled)
	}
}

const toggleDescription = `
Toggle alias

`
