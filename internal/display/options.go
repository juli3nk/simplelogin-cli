package display

import (
	"encoding/json"
	"fmt"
	"os"
)

// OutputFormat defines the output format for data display
type OutputFormat string

const (
	FormatTable OutputFormat = "table"
	FormatJSON  OutputFormat = "json"
)

// DisplayOptions configures how data should be displayed
type DisplayOptions struct {
	Format    OutputFormat
	TableOpts *TableOptions
	Compact   bool
	Color     bool
}

// DefaultDisplayOptions returns sensible defaults
func DefaultDisplayOptions() *DisplayOptions {
	return &DisplayOptions{
		Format:    FormatTable,
		TableOpts: DefaultTableOptions(),
		Compact:   false,
		Color:     false,
	}
}

// DisplayData displays data in the specified format.
//
// JSON output is handled here. Table output is intentionally not implemented
// because rendering a table requires command-specific headers and column
// configuration. Callers that want table output must build and render the table
// themselves using the helpers in this package (NewTable, DefaultTableOptions,
// etc.).
func DisplayData(data interface{}, options *DisplayOptions) error {
	if options == nil {
		options = DefaultDisplayOptions()
	}

	switch options.Format {
	case FormatJSON:
		return displayJSON(data, options.Compact)
	case FormatTable:
		return fmt.Errorf("table format must be rendered by the caller")
	default:
		return fmt.Errorf("unsupported output format: %s", options.Format)
	}
}

// displayJSON outputs data as JSON
func displayJSON(data interface{}, compact bool) error {
	encoder := json.NewEncoder(os.Stdout)
	if !compact {
		encoder.SetIndent("", "  ")
	}
	return encoder.Encode(data)
}

// CompactTableOptions returns options for a more compact table display
func CompactTableOptions() *TableOptions {
	return &TableOptions{
		NoHeaders: false,
		Borders:   false,
		Center:    false,
		Color:     false,
	}
}

// MinimalTableOptions returns options for a minimal table display
func MinimalTableOptions() *TableOptions {
	return &TableOptions{
		NoHeaders: true,
		Borders:   false,
		Center:    false,
		Color:     false,
	}
}
