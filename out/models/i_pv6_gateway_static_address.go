/* -----------------------------------------------------------------
* i_pv6_gateway_static_address.go -
*
* DMTF Redfish IPv6GatewayStaticAddress resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This type represents a single IPv6 static address to be assigned on a network interface.
type IPv6GatewayStaticAddress struct {
	// A valid IPv6 address.
	Address string `json:"Address"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The IPv6 network prefix length, in bits, for this address.
	PrefixLength PrefixLength `json:"PrefixLength,omitempty"`

}
