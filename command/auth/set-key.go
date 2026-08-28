package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/spf13/cobra"
)

func newSetKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-key",
		Short: "Set API key",
		Long:  setApiKeyDescription,
		Args:  cobra.NoArgs,
		RunE:  runSetApiKey,
	}

	return cmd
}

func runSetApiKey(cmd *cobra.Command, args []string) error {
	defer utils.RecoverFunc()

	apiKey, err := readAPIKey()
	if err != nil {
		return err
	}

	if err := config.SaveAPIKey(apiKey); err != nil {
		return fmt.Errorf("failed to save API key: %w", err)
	}

	fmt.Println("API key saved successfully")
	return nil
}

func readAPIKey() (string, error) {
	fmt.Print("Enter API key: ")

	reader := bufio.NewReader(os.Stdin)
	apiKey, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read API key: %w", err)
	}

	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		return "", fmt.Errorf("API key cannot be empty")
	}

	return apiKey, nil
}

// false positive: this is a help text string, not a real credential
//
//nolint:gosec
const setApiKeyDescription = `
Set the API key used to authenticate with SimpleLogin.

The key is read from stdin so it is not exposed in shell history or process
listings. If the OS keyring is unavailable, it is stored in
~/.config/simplelogin-cli/credentials.json with 0600 permissions.

Example:

    echo "your-api-key" | simplelogin-cli auth set-key
`
