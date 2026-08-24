package alias

import (
	"fmt"
	"strconv"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newActivitiesCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "activities [alias_id] [page_id]",
		Aliases: []string{"act"},
		Short:   "List alias activities",
		Long:    activitiesDescription,
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runActivities(cmd, outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Hide table headers")

	return cmd
}

func runActivities(cmd *cobra.Command, outputFormat *string, args []string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	aliasID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	pageID, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	activities, err := client.GetAliasActivities(aliasID, pageID)
	if err != nil {
		return err
	}

	switch *outputFormat {
	case "json":
		if len(activities) == 0 {
			fmt.Printf("%v\n", activities)
			return nil
		}

		if err := display.DisplayData(activities, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default: // table
		if len(activities) == 0 {
			fmt.Println("No activities found.")
			return nil
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

	return nil
}

const activitiesDescription = `
List alias activities

`
