package auth

import (
	"fmt"
	"log"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/spf13/cobra"
)

func newSetKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-key [key]",
		Short: "Set API key",
		Long:  setApiKeyDescription,
		Args:  cobra.ExactArgs(1),
		Run:   runSetApiKey,
	}

	return cmd
}

func runSetApiKey(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	if err := config.SaveApiKey(args[0]); err != nil {
		log.Fatalf("Failed to save API key: %v", err)
	}

	fmt.Println("API key saved successfully")
}

const setApiKeyDescription = `
Set API Key

`
