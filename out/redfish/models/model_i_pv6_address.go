/* -----------------------------------------------------------------
* i_pv6_address.go -
*
* DMTF Redfish IPv6Address resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This is the IPv6 Address.
	Address string `json:"Address,omitempty"`

	// This is the IPv6 Address Prefix Length.
	PrefixLength PrefixLength `json:"PrefixLength,omitempty"`

	// This indicates how the address was determined.
	AddressOrigin IPv6AddressOrigin `json:"AddressOrigin,omitempty"`

	// The current state of this address as defined in RFC 4862.
	AddressState AddressState `json:"AddressState,omitempty"`

}
