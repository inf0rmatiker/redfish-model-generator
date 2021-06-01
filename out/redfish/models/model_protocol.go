/* -----------------------------------------------------------------
* protocol.go -
*
* DMTF Redfish Protocol resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

	// Indicates if the protocol is enabled or disabled.
	ProtocolEnabled bool `json:"ProtocolEnabled,omitempty"`

	// Indicates the protocol port.
	Port float64 `json:"Port,omitempty"`

}
