package simplelogin

import "context"

// APIClient is the public interface for the SimpleLogin API client.
// It is implemented by *simplelogin.Client and can be used to mock the client
// in tests of CLI commands or other consumers.
//
// Implementations must provide WithContext so callers can inject cancellation
// and timeout control.
type APIClient interface {
	WithContext(ctx context.Context) APIClient

	// Alias
	GetAliasOptions(hostname string) (*AliasOptions, error)
	CreateCustomAlias(hostname string, options AliasCreateCustomOptions) (*Alias, error)
	CreateRandomAlias(hostname, mode, note string) (*Alias, error)
	GetAliases(options AliasListOptions, pageID int) ([]Alias, error)
	GetAllAliases(options AliasListOptions) ([]Alias, error)
	GetAlias(aliasID int) (*Alias, error)
	DeleteAlias(aliasID int) (bool, error)
	ToggleAlias(aliasID int) (*AliasToggleResponse, error)
	GetAliasActivities(aliasID, pageID int) ([]AliasActivity, error)
	GetAllAliasActivities(aliasID int) ([]AliasActivity, error)
	UpdateAlias(aliasID int, options AliasUpdateOptions) error
	GetAliasContacts(aliasID, pageID int) ([]AliasContact, error)
	GetAllAliasContacts(aliasID int) ([]AliasContact, error)
	CreateAliasContact(aliasID int, contact string) (*AliasContactCreateResponse, error)

	// Contact
	DeleteContact(contactID int) (*ContactDeleteResponse, error)
	ToggleContact(contactID int) (*ContactBlockResponse, error)

	// Domain
	GetDomains() ([]Domain, error)
	UpdateDomain(domainID int, updateDomain UpdateDomain) (*Domain, error)
	GetDeletedAliasesDomain(domainID int) ([]TrashAlias, error)

	// Mailbox
	GetMailboxes() ([]Mailbox, error)
	CreateMailbox(email string) (*MailboxCreateResponse, error)
	DeleteMailbox(mailboxID int, options MailboxDeleteOptions) error

	// Setting
	GetSetting() (*Setting, error)
	UpdateSetting(setting Setting) (*Setting, error)
	GetSettingDomains() ([]SettingDomain, error)

	// Stats
	GetStats() (*Stats, error)

	// User
	GetUserInfo() (*UserInfo, error)
	UpdateUserInfo(userInfo UserInfoUpdate) (*UserInfo, error)
}
