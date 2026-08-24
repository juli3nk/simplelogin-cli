package userinfo

import (
	"fmt"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/juli3nk/simplelogin-cli/pkg/simplelogin"
	"github.com/spf13/cobra"
)

var (
	name           string
	profilePicture string
)

func newUpdateCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update user info",
		Long:  updateDescription,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpdate(cmd, outputFormat)
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&compact, "compact", false, "Compact output")

	flags.StringVarP(&name, "name", "n", "", "Name")
	flags.StringVarP(&profilePicture, "profile-picture", "p", "", "Profile picture")

	return cmd
}

func runUpdate(cmd *cobra.Command, outputFormat *string) error {
	defer utils.RecoverFunc()

	if name == "" && profilePicture == "" {
		return fmt.Errorf("no update provided")
	}

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	userInfoUpdate := simplelogin.UserInfoUpdate{}

	if name != "" {
		userInfoUpdate.Name = name
	}

	if profilePicture != "" {
		userInfoUpdate.ProfilePicture = profilePicture
	}

	userInfo, err := client.UpdateUserInfo(userInfoUpdate)
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

const updateDescription = `
Update user info

`
