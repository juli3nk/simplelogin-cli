package alias

import (
	"fmt"
	"strconv"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/juli3nk/simplelogin-cli/pkg/simplelogin"
	"github.com/spf13/cobra"
)

var (
	aliasListPinned   bool
	aliasListDisabled bool
	aliasListEnabled  bool
	aliasListQuery    string
)

func newListCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list [page_id]",
		Aliases: []string{"ls"},
		Short:   "List aliases",
		Long:    listDescription,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(cmd, outputFormat, args)
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&compact, "compact", false, "Compact output")
	flags.BoolVar(&noHeaders, "no-headers", false, "Hide table headers")

	flags.BoolVarP(&aliasListPinned, "pinned", "p", false, "Pinned aliases")
	flags.BoolVarP(&aliasListDisabled, "disabled", "d", false, "Disabled aliases")
	flags.BoolVarP(&aliasListEnabled, "enabled", "e", false, "Enabled aliases")
	flags.StringVarP(&aliasListQuery, "query", "q", "", "Query aliases")

	return cmd
}

func runList(cmd *cobra.Command, outputFormat *string, args []string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	opts := simplelogin.AliasListOptions{
		Pinned:   aliasListPinned,
		Disabled: aliasListDisabled,
		Enabled:  aliasListEnabled,
		Query:    aliasListQuery,
	}

	pageID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	aliases, err := client.GetAliases(opts, pageID)
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
		if compact {
			tableOpts = display.CompactTableOptions()
		}

		table := display.NewTable(tableOpts)
		table.SetHeader([]string{"ID", "Email", "Status", "Note", "Created", "Forwards", "Replies", "Blocks"})

		for _, alias := range aliases {
			status := "✓ Enabled"
			if !alias.Enabled {
				status = "✗ Disabled"
			}
			if alias.Pinned {
				status = "📌 Pinned"
			}

			table.Append([]string{
				display.FormatID(alias.ID),
				display.FormatEmail(alias.Email, 35),
				status,
				display.FormatDate(alias.Note),
				display.FormatDate(alias.CreationDate),
				display.FormatID(alias.NbForward),
				display.FormatID(alias.NbReply),
				display.FormatID(alias.NbBlock),
			})
		}

		table.Render()
		fmt.Printf("\nTotal: %d aliases\n", len(aliases))
	}

	return nil
}

const listDescription = `
List aliases

`
