package alias

import (
	"fmt"
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

func newUpdateCommand(outputFormat *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update [alias_id]",
		Aliases: []string{"up"},
		Short:   "Update alias",
		Long:    updateDescription,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpdate(cmd, outputFormat, args)
		},
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

func runUpdate(cmd *cobra.Command, outputFormat *string, args []string) error {
	defer utils.RecoverFunc()

	if note == "" && name == "" && len(mailboxIds) == 0 && !disablePGP && !pinned {
		return fmt.Errorf("no update provided")
	}

	client, err := config.ClientFactoryWithContext(cmd.Context())
	if err != nil {
		return err
	}

	aliasID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
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
		return err
	}

	return nil
}

const updateDescription = `
Update alias

`
