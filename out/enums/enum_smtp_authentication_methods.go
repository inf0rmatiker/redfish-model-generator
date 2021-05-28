/* -----------------------------------------------------------------
 * enum_smtp_authentication_methods.go -
 *
 * DMTF Redfish SMTPAuthenticationMethods enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SMTPAuthenticationMethods string

const (
	// No authentication.
	SMTPAuthenticationMethods_NONE SMTPAuthenticationMethods = "None"

	// Auto-detect.
	SMTPAuthenticationMethods_AUTO_DETECT SMTPAuthenticationMethods = "AutoDetect"

	// PLAIN authentication.
	SMTPAuthenticationMethods_PLAIN SMTPAuthenticationMethods = "Plain"

	// LOGIN authentication.
	SMTPAuthenticationMethods_LOGIN SMTPAuthenticationMethods = "Login"

	// CRAM-MD5 authentication.
	SMTPAuthenticationMethods_CRAM_MD5 SMTPAuthenticationMethods = "CRAM_MD5"

)
