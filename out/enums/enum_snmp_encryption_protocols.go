/* -----------------------------------------------------------------
 * enum_snmp_encryption_protocols.go -
 *
 * DMTF Redfish SNMPEncryptionProtocols enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SNMPEncryptionProtocols string

const (
	// No encryption.
	SNMPEncryptionProtocols_NONE SNMPEncryptionProtocols = "None"

	// CBC-DES encryption.
	SNMPEncryptionProtocols_CBC_DES SNMPEncryptionProtocols = "CBC_DES"

	// CFB128-AES-128 encryption.
	SNMPEncryptionProtocols_CFB128_AES128 SNMPEncryptionProtocols = "CFB128_AES128"

)
