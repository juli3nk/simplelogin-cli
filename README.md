# SimpleLogin CLI

A command-line interface for managing SimpleLogin aliases, domains, contacts, and settings using the SimpleLogin API.

## Features

- **Complete API Coverage**: Full support for all SimpleLogin API endpoints
- **Modern CLI**: Built with Cobra for excellent user experience
- **Multiple Output Formats**: Table and JSON output support
- **Input Validation**: Robust validation using go-playground/validator
- **Pagination Support**: Automatic handling of paginated API responses
- **Error Handling**: Comprehensive error types and user-friendly messages

## Quick Start

### 1. Set up Authentication

```shell
simplelogin-cli auth set-key
```

### 2. Basic Usage

```shell
# List all aliases
simplelogin-cli alias list 0

# Get user information
simplelogin-cli userinfo get

# List mailboxes
simplelogin-cli mailbox list

# List domains
simplelogin-cli domain list

# Get settings
simplelogin-cli setting get
```

## Available Commands

### Aliases

```shell
simplelogin-cli alias activities [alias_id]
simplelogin-cli alias delete [alias_id]      # Delete alias
simplelogin-cli alias get [name]             # Get specific alias
simplelogin-cli alias list [page_id]         # List aliases
simplelogin-cli alias new [alias]            # Create custom alias
simplelogin-cli alias options [hostname]
simplelogin-cli alias random                 # Create random alias
simplelogin-cli alias toggle [alias_id]      # Toggle alias status
simplelogin-cli alias update [alias_id]      # Update alias
```

### Contacts

```shell
simplelogin-cli contact block [contact_id]                   # Get specific 
simplelogin-cli contact create [contact_id] [contact_email]  # Create random alias
simplelogin-cli contact delete [contact_id]                  # Delete alias
simplelogin-cli contact list [alias_id]                      # List contacts for alias
```

### Domains

```shell
simplelogin-cli domain list                # List domains
simplelogin-cli domain trash [domain_id]   # List deleted aliases
simplelogin-cli domain update [domain_id]
```

### Mailboxes

```shell
simplelogin-cli mailbox create [email]       # Create random alias
simplelogin-cli mailbox delete [mailbox_id]  # Delete alias
simplelogin-cli mailbox list                 # List mailboxes
```

### Settings

```shell
simplelogin-cli setting get             # Get current settings
simplelogin-cli setting get-domains     # List available domains
simplelogin-cli setting update          # Update settings
```

### 

```shell
simplelogin-cli stats                   # Get account statistics
```

### User 

```shell
simplelogin-cli userinfo get            # Get user information
simplelogin-cli userinfo update         # Get user information
```

## Output Formats

Most commands support multiple output formats:

```shell
# Table format (default)
simplelogin-cli alias list 0

# JSON format
simplelogin-cli alias list 0 --output json

# Compact output
simplelogin-cli alias list 0 --compact

# No headers
simplelogin-cli alias list 0 --no-headers
```

## Development

### Project Structure

```
├── cmd/simplelogin-cli/    # CLI entry point
├── command/                # CLI commands
├── internal/               # Internal packages
│   ├── config/             # Configuration management
│   └── display/            # Output formatting
├── pkg/simplelogin/        # SimpleLogin API client
└── examples/               # Usage examples
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- **Issues**: [GitHub Issues](https://github.com/juli3nk/simplelogin-cli/issues)
- **Documentation**: [SimpleLogin API Docs](https://github.com/simple-login/app/blob/master/docs/api.md)
- **SimpleLogin**: [https://simplelogin.io](https://simplelogin.io)
