package display

import (
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

// TableOptions configures table display
type TableOptions struct {
	NoHeaders bool
	Borders   bool
	Center    bool
	Color     bool
}

// DefaultTableOptions returns sensible defaults for table display
func DefaultTableOptions() *TableOptions {
	return &TableOptions{
		NoHeaders: false,
		Borders:   true,
		Center:    false,
		Color:     false,
	}
}

// NewTable creates a new table writer with the given options
func NewTable(options *TableOptions) *tablewriter.Table {
	if options == nil {
		options = DefaultTableOptions()
	}

	table := tablewriter.NewWriter(os.Stdout)

	if options.Borders {
		table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
		table.SetCenterSeparator("|")
		table.SetColumnSeparator("|")
		table.SetRowSeparator("-")
	} else {
		table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
	}

	if options.Center {
		table.SetAlignment(tablewriter.ALIGN_CENTER)
	} else {
		table.SetAlignment(tablewriter.ALIGN_LEFT)
	}

	if !options.NoHeaders {
		table.SetHeaderLine(true)
	}

	return table
}

// FormatBool formats a boolean value for display
func FormatBool(b bool) string {
	if b {
		return "✓"
	}
	return "✗"
}

// FormatTimestamp formats a Unix timestamp for display
func FormatTimestamp(timestamp int) string {
	if timestamp == 0 {
		return "-"
	}
	return time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
}

// FormatDate formats a date string for display
func FormatDate(date string) string {
	if date == "" {
		return "-"
	}
	return date
}

// FormatID formats an ID for display
func FormatID(id int) string {
	return strconv.Itoa(id)
}

// FormatEmail formats an email for display with truncation if needed
func FormatEmail(email string, maxLength int) string {
	if email == "" {
		return "-"
	}
	if len(email) <= maxLength {
		return email
	}
	return email[:maxLength-3] + "..."
}
