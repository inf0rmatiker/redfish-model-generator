/* -----------------------------------------------------------------
* enum_file_protocol.go -
*
* SNIA Swordfish FileProtocol enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type FileProtocol string

const (
	// NFSv3, as defined in RFC 1813.
	FileProtocol_NF_SV3 FileProtocol = "NFSv3"

	// NFSv4, as defined in RFC 7530.
	FileProtocol_NF_SV4_0 FileProtocol = "NFSv4_0"

	// NFSv4.1, as defined in RFC 5661.
	FileProtocol_NF_SV4_1 FileProtocol = "NFSv4_1"

	// Server Message Block version 2.0.
	FileProtocol_SM_BV2_0 FileProtocol = "SMBv2_0"

	// Server Message Block version 2.1.
	FileProtocol_SM_BV2_1 FileProtocol = "SMBv2_1"

	// Server Message Block version 3.0.
	FileProtocol_SM_BV3_0 FileProtocol = "SMBv3_0"

	// Server Message Block version 3.0.2.
	FileProtocol_SM_BV3_0_2 FileProtocol = "SMBv3_0_2"

	// Server Message Block version 3.1.1.
	FileProtocol_SM_BV3_1_1 FileProtocol = "SMBv3_1_1"

)
