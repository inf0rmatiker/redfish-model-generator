/* -----------------------------------------------------------------
 * enum_snmp_authentication_protocols.go -
 *
 * DMTF Redfish SNMPAuthenticationProtocols enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SNMPAuthenticationProtocols string

const (
	// No authentication.
	SNMPAuthenticationProtocols_NONE SNMPAuthenticationProtocols = "None"

	// Trap community string authentication.
	SNMPAuthenticationProtocols_COMMUNITY_STRING SNMPAuthenticationProtocols = "CommunityString"

	// HMAC-MD5-96 authentication.
	SNMPAuthenticationProtocols_HMAC_MD5 SNMPAuthenticationProtocols = "HMAC_MD5"

	// HMAC-SHA-96 authentication.
	SNMPAuthenticationProtocols_HMAC_SHA96 SNMPAuthenticationProtocols = "HMAC_SHA96"

)
