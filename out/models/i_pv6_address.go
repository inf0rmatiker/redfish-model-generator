/* -----------------------------------------------------------------
* i_pv6_address.go -
*
* DMTF Redfish IPv6Address resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type describes an IPv6 address.
type IPv6Address struct {
	// The IPv6 address.
	Address string `json:"Address,omitempty"`

	// This indicates how the address was determined.
	AddressOrigin IPv6AddressOrigin `json:"AddressOrigin,omitempty"`

	// The current RFC4862-defined state of this address.
	AddressState AddressState `json:"AddressState,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The IPv6 address prefix Length.
	PrefixLength PrefixLength `json:"PrefixLength,omitempty"`

}
