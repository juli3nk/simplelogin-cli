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
	createNewAliasPrefix  string
	createNewSignedSuffix string
	createNewMailboxIds   []int
	createNewNote         string
	createNewName         string
)

func newCreateNewCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new [hostname]",
		Short: "Create new alias",
		Long:  createNewDescription,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runCreateNew(outputFormat, args)
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&compact, "compact", false, "Compact output")

	flags.StringVarP(&createNewAliasPrefix, "alias-prefix", "a", "", "The first part of the alias that user can choose")
	flags.StringVarP(&createNewSignedSuffix, "signed-suffix", "s", "", "Should be one of the suffixes returned in the GET /api/v5/alias/options endpoint")
	flags.IntSliceVarP(&createNewMailboxIds, "mailbox-ids", "m", []int{}, "List of mailbox_id that 'owns' this alias")
	flags.StringVar(&createNewNote, "note", "", "Alias note")
	flags.StringVar(&createNewName, "name", "", "Alias name")

	return cmd
}

func runCreateNew(outputFormat *string, args []string) {
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

	input := simplelogin.AliasCreateCustomOptions{
		AliasPrefix:  createNewAliasPrefix,
		SignedSuffix: createNewSignedSuffix,
		MailboxIDs:   createNewMailboxIds,
	}
	if createNewNote != "" {
		input.Note = createNewNote
	}
	if createNewName != "" {
		input.Name = createNewName
	}

	alias, err := client.CreateCustomAlias(hostname, input)
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

const createNewDescription = `
Create new alias

`
