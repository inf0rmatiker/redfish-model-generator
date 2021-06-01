/* -----------------------------------------------------------------
* enum_transfer_protocol_type.go -
*
* DMTF Redfish TransferProtocolType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type TransferProtocolType string

const (
	// Common Internet File System (CIFS).
	TransferProtocolType_CIFS TransferProtocolType = "CIFS"

	// File Transfer Protocol (FTP).
	TransferProtocolType_FTP TransferProtocolType = "FTP"

	// Secure File Transfer Protocol (SFTP).
	TransferProtocolType_SFTP TransferProtocolType = "SFTP"

	// Hypertext Transfer Protocol (HTTP).
	TransferProtocolType_HTTP TransferProtocolType = "HTTP"

	// Hypertext Transfer Protocol Secure (HTTPS).
	TransferProtocolType_HTTPS TransferProtocolType = "HTTPS"

	// Network File System (NFS).
	TransferProtocolType_NSF TransferProtocolType = "NSF"

	// Secure Copy Protocol (SCP).
	TransferProtocolType_SCP TransferProtocolType = "SCP"

	// Trivial File Transfer Protocol (TFTP).
	TransferProtocolType_TFTP TransferProtocolType = "TFTP"

	// A manufacturer-defined protocol.
	TransferProtocolType_OEM TransferProtocolType = "OEM"

	// Network File System (NFS).
	TransferProtocolType_NFS TransferProtocolType = "NFS"

)
