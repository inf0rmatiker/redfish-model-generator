/* -----------------------------------------------------------------
* lldp_receive.go -
*
* DMTF Redfish LLDPReceive resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Link Layer Data Protocol (LLDP) data received from the remote partner across this link.
type LLDPReceive struct {
	// Link Layer Data Protocol (LLDP) chassis ID received from the remote partner across this link.
	ChassisId string `json:"ChassisId,omitempty"`

	// The type of identifier used for the chassis ID received from the remote partner across this link.
	ChassisIdSubtype IEEE802IdSubtype `json:"ChassisIdSubtype,omitempty"`

	// The IPv4 management address received from the remote partner across this link.
	ManagementAddressIPv4 string `json:"ManagementAddressIPv4,omitempty"`

	// The IPv6 management address received from the remote partner across this link.
	ManagementAddressIPv6 string `json:"ManagementAddressIPv6,omitempty"`

	// The management MAC address received from the remote partner across this link.
	ManagementAddressMAC string `json:"ManagementAddressMAC,omitempty"`

	// The management VLAN ID received from the remote partner across this link.
	ManagementVlanId int `json:"ManagementVlanId,omitempty"`

	// A colon delimited string of hexadecimal octets identifying a port.
	PortId string `json:"PortId,omitempty"`

	// The port ID subtype received from the remote partner across this link.
	PortIdSubtype IEEE802IdSubtype `json:"PortIdSubtype,omitempty"`

}
