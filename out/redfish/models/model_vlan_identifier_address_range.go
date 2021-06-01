/* -----------------------------------------------------------------
* vlan_identifier_address_range.go -
*
* DMTF Redfish VLANIdentifierAddressRange resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// VLANIdentifierAddressRange - VLAN tag related addressing for an Ethernet fabric or for end host networks.
// Modeled after DMTF Redfish schema VLANIdentifierAddressRange
type VLANIdentifierAddressRange struct {
	// Virtual LAN (VLAN) tag lower value.
	Lower int `json:"Lower,omitempty"`

	// Virtual LAN (VLAN) tag upper value.
	Upper int `json:"Upper,omitempty"`

}
