/* -----------------------------------------------------------------
* i_pv4_address.go -
*
* DMTF Redfish IPv4Address resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type describes an IPv4 address.
type IPv4Address struct {
	// The IPv4 address.
	Address string `json:"Address,omitempty"`

	// This indicates how the address was determined.
	AddressOrigin IPv4AddressOrigin `json:"AddressOrigin,omitempty"`

	// The IPv4 gateway for this address.
	Gateway string `json:"Gateway,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The IPv4 subnet mask.
	SubnetMask SubnetMask `json:"SubnetMask,omitempty"`

}
