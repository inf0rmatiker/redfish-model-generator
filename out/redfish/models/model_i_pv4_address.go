/* -----------------------------------------------------------------
* i_pv4_address.go -
*
* DMTF Redfish IPv4Address resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This is the IPv4 Address.
	Address string `json:"Address,omitempty"`

	// This is the IPv4 Subnet mask.
	SubnetMask SubnetMask `json:"SubnetMask,omitempty"`

	// This indicates how the address was determined.
	AddressOrigin IPv4AddressOrigin `json:"AddressOrigin,omitempty"`

	// This is the IPv4 gateway for this address.
	Gateway string `json:"Gateway,omitempty"`

}
