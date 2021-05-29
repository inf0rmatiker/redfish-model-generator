/* -----------------------------------------------------------------
* ldap_search_settings.go -
*
* DMTF Redfish LDAPSearchSettings resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The settings to search a generic LDAP service.
type LDAPSearchSettings struct {
	// The base distinguished names to use to search an external LDAP service.
	BaseDistinguishedNames []string `json:"BaseDistinguishedNames,omitempty"`

	// The attribute name that contains the LDAP group name entry.
	GroupNameAttribute string `json:"GroupNameAttribute,omitempty"`

	// The attribute name that contains the groups for a user on the LDAP user entry.
	GroupsAttribute string `json:"GroupsAttribute,omitempty"`

	// The attribute name that contains the LDAP user name entry.
	UsernameAttribute string `json:"UsernameAttribute,omitempty"`

}
