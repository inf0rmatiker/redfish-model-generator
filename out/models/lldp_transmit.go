/* -----------------------------------------------------------------
* lldp_transmit.go -
*
* DMTF Redfish LLDPTransmit resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Link Layer Data Protocol (LLDP) data being transmitted on this link.
type LLDPTransmit struct {
	// Link Layer Data Protocol (LLDP) chassis ID.
	ChassisId string `json:"ChassisId,omitempty"`

	// The type of identifier used for the chassis ID.
	ChassisIdSubtype IEEE802IdSubtype `json:"ChassisIdSubtype,omitempty"`

	// The IPv4 management address to be transmitted from this endpoint.
	ManagementAddressIPv4 string `json:"ManagementAddressIPv4,omitempty"`

	// The IPv6 management address to be transmitted from this endpoint.
	ManagementAddressIPv6 string `json:"ManagementAddressIPv6,omitempty"`

	// The management MAC address to be transmitted from this endpoint.
	ManagementAddressMAC string `json:"ManagementAddressMAC,omitempty"`

	// The management VLAN ID to be transmitted from this endpoint.
	ManagementVlanId int `json:"ManagementVlanId,omitempty"`

	// A colon delimited string of hexadecimal octets identifying a port to be transmitted from this endpoint.
	PortId string `json:"PortId,omitempty"`

	// The port ID subtype to be transmitted from this endpoint.
	PortIdSubtype IEEE802IdSubtype `json:"PortIdSubtype,omitempty"`

}
