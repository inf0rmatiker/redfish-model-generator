/* -----------------------------------------------------------------
 * enum_smtp_connection_protocol.go -
 *
 * DMTF Redfish SMTPConnectionProtocol enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SMTPConnectionProtocol string

const (
	// Clear text.
	SMTPConnectionProtocol_NONE SMTPConnectionProtocol = "None"

	// Auto-detect.
	SMTPConnectionProtocol_AUTO_DETECT SMTPConnectionProtocol = "AutoDetect"

	// StartTLS.
	SMTPConnectionProtocol_START_TLS SMTPConnectionProtocol = "StartTLS"

	// TLS/SSL.
	SMTPConnectionProtocol_TLS_SSL SMTPConnectionProtocol = "TLS_SSL"

)
