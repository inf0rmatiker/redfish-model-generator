/* -----------------------------------------------------------------
* ldap_service.go -
*
* DMTF Redfish LDAPService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// LDAPService - The settings required to parse a generic LDAP service.
// Modeled after DMTF Redfish schema LDAPService
type LDAPService struct {
	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The required settings to search an external LDAP service.
	SearchSettings LDAPSearchSettings `json:"SearchSettings,omitempty"`

}
