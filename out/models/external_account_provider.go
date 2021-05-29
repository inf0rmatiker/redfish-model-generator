/* -----------------------------------------------------------------
* external_account_provider.go -
*
* DMTF Redfish ExternalAccountProvider resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The external account provider services that can provide accounts for this manager to use for authentication.
type ExternalAccountProvider struct {
	// The type of external account provider to which this service connects.
	AccountProviderType AccountProviderTypes `json:"AccountProviderType,omitempty"`

	// The authentication information for the external account provider.
	Authentication Authentication `json:"Authentication,omitempty"`

	// The link to a collection of certificates that the external account provider uses.
	Certificates CertificateCollection `json:"Certificates,omitempty"`

	// The additional mapping information needed to parse a generic LDAP service.
	LDAPService LDAPService `json:"LDAPService,omitempty"`

	// Indicates if the Password property is set.
	PasswordSet bool `json:"PasswordSet,omitempty"`

	// The authentication priority for the external account provider.
	Priority int `json:"Priority,omitempty"`

	// The mapping rules to convert the external account providers account information to the local Redfish role.
	RemoteRoleMapping array `json:"RemoteRoleMapping,omitempty"`

	// The addresses of the user account providers to which this external account provider links.  The format of this field depends on the type of external account provider.
	ServiceAddresses []string `json:"ServiceAddresses,omitempty"`

	// An indication of whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The additional information needed to parse a TACACS+ services.
	TACACSplusService TACACSplusService `json:"TACACSplusService,omitempty"`

}