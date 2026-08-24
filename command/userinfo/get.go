package userinfo

import (
	"fmt"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/spf13/cobra"
)

func newGetCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get user info",
		Long:  getDescription,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(cmd, outputFormat)
		},
	}

	cmd.Flags().BoolVar(&compact, "compact", false, "Compact output")

	return cmd
}

func runGet(cmd *cobra.Command, outputFormat *string) error {
	defer utils.RecoverFunc()

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	userInfo, err := client.GetUserInfo()
	if err != nil {
		return err
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(userInfo, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			return err
		}
	default:
		fmt.Printf("Name: %s\n", userInfo.Name)
		fmt.Printf("Email: %s\n", userInfo.Email)
		fmt.Printf("Is Premium: %t\n", userInfo.IsPremium)
		fmt.Printf("In Trial: %t\n", userInfo.InTrial)
		fmt.Printf("Profile Picture URL: %s\n", userInfo.ProfilePictureURL)
		fmt.Printf("Max Alias Free Plan: %d\n", userInfo.MaxAliasFreePlan)
	}

	return nil
}

const getDescription = `
Get user info

`
