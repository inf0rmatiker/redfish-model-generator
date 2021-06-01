/* -----------------------------------------------------------------
* account_service.go -
*
* DMTF Redfish AccountService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// AccountService - The AccountService schema defines an account service.  The properties are common to, and enable management of, all user accounts.  The properties include the password requirements and control features, such as account lockout.  The schema also contains links to the manager accounts and roles.
// Modeled after DMTF Redfish schema AccountService
type AccountService struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The period of time, in seconds, between the last failed login attempt and the reset of the lockout threshold counter.  This value must be less than or equal to the AccountLockoutDuration value.  A reset sets the counter to `0`.
	AccountLockoutCounterResetAfter int `json:"AccountLockoutCounterResetAfter,omitempty"`

	// The period of time, in seconds, that an account is locked after the number of failed login attempts reaches the account lockout threshold, within the period between the last failed login attempt and the reset of the lockout threshold counter.  If this value is `0`, no lockout will occur.  If the AccountLockoutCounterResetEnabled value is `false`, this property is ignored.
	AccountLockoutDuration int `json:"AccountLockoutDuration,omitempty"`

	// The number of allowed failed login attempts before a user account is locked for a specified duration.  If `0`, the account is never locked.
	AccountLockoutThreshold int `json:"AccountLockoutThreshold,omitempty"`

	// The collection of manager accounts.
	Accounts map[string]interface{} `json:"Accounts,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The first Active Directory external account provider that this account service supports.
	ActiveDirectory ExternalAccountProvider `json:"ActiveDirectory,omitempty"`

	// The additional external account providers that this account service uses.
	AdditionalExternalAccountProviders map[string]interface{} `json:"AdditionalExternalAccountProviders,omitempty"`

	// The number of authorization failures per account that are allowed before the failed attempt is logged to the manager log.
	AuthFailureLoggingThreshold int `json:"AuthFailureLoggingThreshold,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The first LDAP external account provider that this account service supports.
	LDAP ExternalAccountProvider `json:"LDAP,omitempty"`

	// An indication of how the service uses the accounts collection within this account service as part of authentication.  The enumerated values describe the details for each mode.
	LocalAccountAuth LocalAccountAuth `json:"LocalAccountAuth,omitempty"`

	// The maximum password length for this account service.
	MaxPasswordLength int `json:"MaxPasswordLength,omitempty"`

	// The minimum password length for this account service.
	MinPasswordLength int `json:"MinPasswordLength,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The link to the mapping of the privileges required to complete a requested operation on a URI associated with this service.
	PrivilegeMap map[string]interface{} `json:"PrivilegeMap,omitempty"`

	// The collection of Redfish roles.
	Roles map[string]interface{} `json:"Roles,omitempty"`

	// An indication of whether the account service is enabled.  If `true`, it is enabled.  If `false`, it is disabled and users cannot be created, deleted, or modified, and new sessions cannot be started.  However, established sessions might still continue to run.  Any service, such as the session service, that attempts to access the disabled account service fails.  However, this does not affect HTTP Basic Authentication connections.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

}
