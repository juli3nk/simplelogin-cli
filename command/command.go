package command

import (
	"github.com/spf13/cobra"

	"github.com/juli3nk/simplelogin-cli/command/alias"
	"github.com/juli3nk/simplelogin-cli/command/auth"
	"github.com/juli3nk/simplelogin-cli/command/contact"
	"github.com/juli3nk/simplelogin-cli/command/domain"
	"github.com/juli3nk/simplelogin-cli/command/mailbox"
	"github.com/juli3nk/simplelogin-cli/command/setting"
	"github.com/juli3nk/simplelogin-cli/command/stats"
	"github.com/juli3nk/simplelogin-cli/command/userinfo"
)

var usageTemplate = `{{ .Short | trim }}

Usage:{{ if .Runnable }}
  {{ if .HasAvailableFlags }}{{ appendIfNotPresent .UseLine "[flags]" }}{{ else }}{{ .UseLine }}{{ end }}{{ end }}{{ if .HasAvailableSubCommands }}
  {{ .CommandPath }} [command]{{ end }}{{ if gt .Aliases 0 }}

Aliases:
{{ .NameAndAliases }}{{ end }}{{ if .HasExample }}

Examples:
{{ .Example }}{{ end }}{{ if .HasAvailableSubCommands}}

Available Commands:{{ range .Commands }}{{ if .IsAvailableCommand }}
{{ rpad .Name .NamePadding }} {{ .Short }}{{ end }}{{ end }}{{ end }}{{ if .HasAvailableLocalFlags }}

Flags:
{{ .LocalFlags.FlagUsages | trimRightSpace }}{{ end }}{{ if .HasAvailableInheritedFlags }}

Global Flags:
{{ .InheritedFlags.FlagUsages | trimRightSpace }}{{ end }}{{ if .HasHelpSubCommands }}

Additional help topics: {{ range .Commands }}{{ if .IsHelpCommand }}
{{ rpad .CommandPath .CommandPathPadding }} {{ .Short }}{{ end }}{{ end }}{{ end }}{{ if .HasAvailableSubCommands }}

Use "{{ .CommandPath }} [command] --help" for more information about a command.{{ end }}
`

var helpTemplate = `
{{ if or .Runnable .HasSubCommands }}{{ .UsageString }}{{ end }}`

var outputFormat string

func NewSimpleLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "simplelogin-cli",
		Short: "SimpleLogin CLI",
		Long:  "SimpleLogin CLI",
	}

	cmd.SetHelpTemplate(helpTemplate)
	cmd.SetUsageTemplate(usageTemplate)

	cmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "Output format (table, json)")

	cmd.AddCommand(alias.NewCommand(&outputFormat))
	cmd.AddCommand(auth.NewCommand())
	cmd.AddCommand(contact.NewCommand(&outputFormat))
	cmd.AddCommand(domain.NewCommand(&outputFormat))
	cmd.AddCommand(mailbox.NewCommand(&outputFormat))
	cmd.AddCommand(setting.NewCommand(&outputFormat))
	cmd.AddCommand(stats.NewCommand(&outputFormat))
	cmd.AddCommand(userinfo.NewCommand(&outputFormat))

	return cmd
}
