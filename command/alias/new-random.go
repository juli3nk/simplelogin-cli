package alias

import (
	"fmt"
	"log"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/internal/display"
	"github.com/juli3nk/simplelogin-cli/pkg/simplelogin"
	"github.com/spf13/cobra"
)

var (
	createRandomMode string
	createRandomNote string
)

func newCreateRandomCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "random [hostname]",
		Short: "Create random alias",
		Long:  createRandomDescription,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runCreateRandom(outputFormat, args)
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&compact, "compact", false, "Compact output")

	flags.StringVarP(&createRandomMode, "mode", "m", "", "The mode of the alias")
	flags.StringVar(&createRandomNote, "note", "", "Alias note")

	return cmd
}

func runCreateRandom(outputFormat *string, args []string) {
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

	hostname := args[0]

	alias, err := client.CreateRandomAlias(hostname, createRandomMode, createRandomNote)
	if err != nil {
		log.Fatal(err)
	}

	switch *outputFormat {
	case "json":
		if err := display.DisplayData(alias, &display.DisplayOptions{
			Format:  display.FormatJSON,
			Compact: compact,
		}); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("CreationDate: %s\n", alias.CreationDate)
		fmt.Printf("CreationTimestamp: %d\n", alias.CreationTimestamp)
		fmt.Printf("Email: %s\n", alias.Email)
		fmt.Printf("Name: %s\n", alias.Name)
		fmt.Printf("Enabled: %t\n", alias.Enabled)
		fmt.Printf("ID: %d\n", alias.ID)
		fmt.Printf("Mailbox: %+v\n", alias.Mailbox)
		for _, mailbox := range alias.Mailboxes {
			fmt.Printf("Mailbox: %+v\n", mailbox)
		}
		fmt.Printf("LatestActivity: %+v\n", alias.LatestActivity)
		fmt.Printf("NbBlock: %d\n", alias.NbBlock)
		fmt.Printf("NbForward: %d\n", alias.NbForward)
		fmt.Printf("NbReply: %d\n", alias.NbReply)
		fmt.Printf("Note: %s\n", alias.Note)
		fmt.Printf("Pinned: %t\n", alias.Pinned)
	}
}

const createRandomDescription = `
Create random alias

`
