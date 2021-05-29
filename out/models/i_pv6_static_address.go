/* -----------------------------------------------------------------
* i_pv6_static_address.go -
*
* DMTF Redfish IPv6StaticAddress resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type represents a single IPv6 static address to be assigned on a network interface.
type IPv6StaticAddress struct {
	// A valid IPv6 address.
	Address string `json:"Address"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The prefix length, in bits, of this IPv6 address.
	PrefixLength PrefixLength `json:"PrefixLength"`

}
