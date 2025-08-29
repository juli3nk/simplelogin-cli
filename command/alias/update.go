package alias

import (
	"log"
	"strconv"

	"github.com/juli3nk/go-utils"
	"github.com/juli3nk/simplelogin-cli/internal/config"
	"github.com/juli3nk/simplelogin-cli/pkg/simplelogin"
	"github.com/spf13/cobra"
)

var (
	note       string
	name       string
	mailboxIds []int
	disablePGP bool
	pinned     bool
)

func newUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update [alias_id]",
		Aliases: []string{"up"},
		Short:   "Update alias",
		Long:    updateDescription,
		Args:    cobra.ExactArgs(1),
		Run:     runUpdate,
	}

	flags := cmd.Flags()
	flags.BoolVar(&compact, "compact", false, "Compact output")

	flags.StringVar(&note, "note", "", "Note")
	flags.StringVar(&name, "name", "", "Name")
	flags.IntSliceVarP(&mailboxIds, "mailbox-ids", "m", []int{}, "Mailbox IDs")
	flags.BoolVarP(&disablePGP, "disable-pgp", "d", false, "Disable PGP")
	flags.BoolVarP(&pinned, "pinned", "p", false, "Pinned")

	return cmd
}

func runUpdate(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	if note == "" && name == "" && len(mailboxIds) == 0 && !disablePGP && !pinned {
		log.Fatal("No update provided")
	}

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

	aliasID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	aliasInput := simplelogin.AliasUpdateOptions{}
	if note != "" {
		aliasInput.Note = note
	}
	if name != "" {
		aliasInput.Name = name
	}
	if len(mailboxIds) > 0 {
		aliasInput.MailboxIDs = mailboxIds
	}
	if disablePGP {
		aliasInput.DisablePGP = true
	}
	if pinned {
		aliasInput.Pinned = true
	}

	err = client.UpdateAlias(aliasID, aliasInput)
	if err != nil {
		log.Fatal(err)
	}
}

const updateDescription = `
Update alias

`
