/* -----------------------------------------------------------------
* i_pv6_address_policy_entry.go -
*
* DMTF Redfish IPv6AddressPolicyEntry resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The entry in the RFC6724-defined address selection policy table.
type IPv6AddressPolicyEntry struct {
	// The IPv6 label, as defined in RFC6724, section 2.1.
	Label int `json:"Label,omitempty"`

	// The IPv6 precedence, as defined in RFC6724, section 2.1.
	Precedence int `json:"Precedence,omitempty"`

	// The IPv6 address prefix, as defined in RFC6724, section 2.1.
	Prefix string `json:"Prefix,omitempty"`

}
