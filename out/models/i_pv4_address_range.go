/* -----------------------------------------------------------------
* i_pv4_address_range.go -
*
* DMTF Redfish IPv4AddressRange resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// IPv4 related address range for an Ethernet fabric.
type IPv4AddressRange struct {
	// Lower IPv4 network address.
	Lower string `json:"Lower,omitempty"`

	// Upper IPv4 network address.
	Upper string `json:"Upper,omitempty"`

}
