/* -----------------------------------------------------------------
* snmp_user_info.go -
*
* DMTF Redfish SNMPUserInfo resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The SNMP settings for an account.
type SNMPUserInfo struct {
	// The secret authentication key for SNMPv3.
	AuthenticationKey string `json:"AuthenticationKey,omitempty"`

	// Indicates if the AuthenticationKey property is set.
	AuthenticationKeySet bool `json:"AuthenticationKeySet,omitempty"`

	// The authentication protocol for SNMPv3.
	AuthenticationProtocol SNMPAuthenticationProtocols `json:"AuthenticationProtocol,omitempty"`

	// The secret encryption key used in SNMPv3.
	EncryptionKey string `json:"EncryptionKey,omitempty"`

	// Indicates if the EncryptionKey property is set.
	EncryptionKeySet bool `json:"EncryptionKeySet,omitempty"`

	// The encryption protocol for SNMPv3.
	EncryptionProtocol SNMPEncryptionProtocols `json:"EncryptionProtocol,omitempty"`

}
