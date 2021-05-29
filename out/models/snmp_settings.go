/* -----------------------------------------------------------------
* snmp_settings.go -
*
* DMTF Redfish SNMPSettings resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Settings for an SNMP event destination.
type SNMPSettings struct {
	// The secret authentication key for SNMPv3.
	AuthenticationKey string `json:"AuthenticationKey,omitempty"`

	// The authentication protocol for SNMPv3.
	AuthenticationProtocol SNMPAuthenticationProtocols `json:"AuthenticationProtocol,omitempty"`

	// The secret authentication key for SNMPv3.
	EncryptionKey string `json:"EncryptionKey,omitempty"`

	// The encryption protocol for SNMPv3.
	EncryptionProtocol SNMPEncryptionProtocols `json:"EncryptionProtocol,omitempty"`

	// The SNMP trap community string.
	TrapCommunity string `json:"TrapCommunity,omitempty"`

}
