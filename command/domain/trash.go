package domain

import (
	"fmt"
	"strconv"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newTrashCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trash [domain_id]",
		Short: "Trash domains",
		Long:  trashDescription,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTrash(cmd, outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Hide table headers")

	return cmd
}

func runTrash(cmd *cobra.Command, outputFormat *string, args []string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	domainID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	aliases, err := client.GetDeletedAliasesDomain(domainID)
	if err != nil {
		return err
	}

	switch *outputFormat {
	case "json":
		if len(aliases) == 0 {
			fmt.Printf("%v\n", aliases)
			return nil
		}

		if err := display.DisplayData(aliases, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default: // table
		if len(aliases) == 0 {
			fmt.Println("No aliases found.")
			return nil
		}

		tableOpts := display.DefaultTableOptions()
		if noHeaders {
			tableOpts.NoHeaders = true
		}

		table := display.NewTable(tableOpts)

		table.SetHeader([]string{"Alias", "Deletion Timestamp"})

		for _, alias := range aliases {
			table.Append([]string{alias.Alias, strconv.Itoa(alias.DeletionTimestamp)})
		}

		table.Render()
	}

	return nil
}

const trashDescription = `
Trash domains

`
