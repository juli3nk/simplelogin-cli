package userinfo

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
		Short: "Get user info",
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

	userInfo, err := client.GetUserInfo()
	if err != nil {
		log.Fatal(err)
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(userInfo, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("Name: %s\n", userInfo.Name)
		fmt.Printf("Email: %s\n", userInfo.Email)
		fmt.Printf("Is Premium: %t\n", userInfo.IsPremium)
		fmt.Printf("In Trial: %t\n", userInfo.InTrial)
		fmt.Printf("Profile Picture URL: %s\n", userInfo.ProfilePictureURL)
		fmt.Printf("Max Alias Free Plan: %d\n", userInfo.MaxAliasFreePlan)
	}
}

const getDescription = `
Get user info

`
