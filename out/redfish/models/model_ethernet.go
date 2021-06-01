/* -----------------------------------------------------------------
* ethernet.go -
*
* DMTF Redfish Ethernet resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Ethernet - This type describes Ethernet capabilities, status, and configuration of a network device function.
// Modeled after DMTF Redfish schema Ethernet
type Ethernet struct {
	// This is the currently configured MAC address of the (logical port) network device function.
	MACAddress string `json:"MACAddress,omitempty"`

	// The Maximum Transmission Unit (MTU) configured for this network device function.
	MTUSize int `json:"MTUSize,omitempty"`

	// This is the permanent MAC address assigned to this network device function (physical function).
	PermanentMACAddress string `json:"PermanentMACAddress,omitempty"`

}
