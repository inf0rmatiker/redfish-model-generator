/* -----------------------------------------------------------------
* protocol.go -
*
* DMTF Redfish Protocol resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The settings for a network protocol associated with a manager.
type Protocol struct {
	// The protocol port.
	Port int `json:"Port,omitempty"`

	// An indication of whether the protocol is enabled.
	ProtocolEnabled bool `json:"ProtocolEnabled,omitempty"`

}
