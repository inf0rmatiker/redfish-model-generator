/* -----------------------------------------------------------------
* snmp_protocol.go -
*
* DMTF Redfish SNMPProtocol resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The settings for a network protocol associated with a manager.
type SNMPProtocol struct {
	// The authentication protocol used for SNMP access to this manager.
	AuthenticationProtocol SNMPAuthenticationProtocols `json:"AuthenticationProtocol,omitempty"`

	// The access level of the SNMP community.
	CommunityAccessMode SNMPCommunityAccessMode `json:"CommunityAccessMode,omitempty"`

	// The SNMP community strings.
	CommunityStrings array `json:"CommunityStrings,omitempty"`

	// Indicates if access via SNMPv1 is enabled.
	EnableSNMPv1 bool `json:"EnableSNMPv1,omitempty"`

	// Indicates if access via SNMPv2c is enabled.
	EnableSNMPv2c bool `json:"EnableSNMPv2c,omitempty"`

	// Indicates if access via SNMPv3 is enabled.
	EnableSNMPv3 bool `json:"EnableSNMPv3,omitempty"`

	// The encryption protocol used for SNMPv3 access to this manager.
	EncryptionProtocol SNMPEncryptionProtocols `json:"EncryptionProtocol,omitempty"`

	// The engine ID.
	EngineId EngineId `json:"EngineId,omitempty"`

	// Indicates if the community strings should be hidden.
	HideCommunityStrings bool `json:"HideCommunityStrings,omitempty"`

	// The protocol port.
	Port int `json:"Port,omitempty"`

	// An indication of whether the protocol is enabled.
	ProtocolEnabled bool `json:"ProtocolEnabled,omitempty"`

}
