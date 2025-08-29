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

func newActivitiesCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "activities [alias_id] [page_id]",
		Aliases: []string{"act"},
		Short:   "List alias activities",
		Long:    activitiesDescription,
		Args:    cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			runActivities(outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Hide table headers")

	return cmd
}

func runActivities(outputFormat *string, args []string) {
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

	pageID, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	activities, err := client.GetAliasActivities(aliasID, pageID)
	if err != nil {
		log.Fatal(err)
	}

	switch *outputFormat {
	case "json":
		if len(activities) == 0 {
			fmt.Printf("%v\n", activities)
			return
		}

		if err := display.DisplayData(activities, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			log.Fatal(err)
		}
	default: // table
		if len(activities) == 0 {
			fmt.Println("No activities found.")
			return
		}

		tableOpts := display.DefaultTableOptions()
		if noHeaders {
			tableOpts.NoHeaders = true
		}
		if compact {
			tableOpts = display.CompactTableOptions()
		}

		table := display.NewTable(tableOpts)
		table.SetHeader([]string{"Action", "From", "Timestamp", "To", "Reverse Alias", "Reverse Alias Address"})

		for _, activity := range activities {
			table.Append([]string{
				activity.Action,
				activity.From,
				display.FormatID(activity.Timestamp),
				activity.To,
				activity.ReverseAlias,
				activity.ReverseAliasAddress,
			})
		}

		table.Render()
		fmt.Printf("\nTotal: %d activities\n", len(activities))
	}
}

const activitiesDescription = `
List alias activities

`
