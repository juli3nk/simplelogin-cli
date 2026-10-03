package alias

import (
	"fmt"
	"strconv"

	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [alias_id]",
		Short: "Delete an alias",
		Long:  deleteDescription,
		Args:  cobra.ExactArgs(1),
		RunE:  runDelete,
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runDelete(cmd *cobra.Command, args []string) error {
	outputFormat, err := cmd.Root().PersistentFlags().GetString("output")
	if err != nil {
		return err
	}

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	aliasID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	alias, err := client.DeleteAlias(aliasID)
	if err != nil {
		return err
	}

	switch outputFormat {
	case "json":
		if err := display.DisplayData(alias, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default:
		fmt.Printf("Alias deleted: %+v\n", alias)
	}

	return nil
}

const deleteDescription = `
Delete an alias

`
