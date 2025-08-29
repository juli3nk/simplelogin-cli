package domain

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

func newTrashCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trash [domain_id]",
		Short: "Trash domains",
		Long:  trashDescription,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runTrash(outputFormat, args)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Hide table headers")

	return cmd
}

func runTrash(outputFormat *string, args []string) {
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

	domainID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	aliases, err := client.GetDeletedAliasesDomain(domainID)
	if err != nil {
		log.Fatal(err)
	}

	switch *outputFormat {
	case "json":
		if len(aliases) == 0 {
			fmt.Printf("%v\n", aliases)
			return
		}

		if err := display.DisplayData(aliases, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			log.Fatal(err)
		}
	default: // table
		if len(aliases) == 0 {
			fmt.Println("No aliases found.")
			return
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
}

const trashDescription = `
Trash domains

`
