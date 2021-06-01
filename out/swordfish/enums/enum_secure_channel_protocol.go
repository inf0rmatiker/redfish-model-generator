/* -----------------------------------------------------------------
* enum_secure_channel_protocol.go -
*
* SNIA Swordfish SecureChannelProtocol enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type SecureChannelProtocol string

const (
	// No encryption.
	SecureChannelProtocol_NONE SecureChannelProtocol = "None"

	// Transport Layer Security (TLS), as defined by IETF RFC 5246.
	SecureChannelProtocol_TLS SecureChannelProtocol = "TLS"

	// Internet Protocol Security (IPsec), as defined by IETF RFC 2401.
	SecureChannelProtocol_I_PSEC SecureChannelProtocol = "IPsec"

	// RPC access to the Generic Security Services Application Programming Interface (GSS-API), as defined by IETF RPC 2203.
	SecureChannelProtocol_RPCSEC_GSS SecureChannelProtocol = "RPCSEC_GSS"

)
