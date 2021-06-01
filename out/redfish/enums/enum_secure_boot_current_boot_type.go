/* -----------------------------------------------------------------
* enum_secure_boot_current_boot_type.go -
*
* DMTF Redfish SecureBootCurrentBootType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type SecureBootCurrentBootType string

const (
	// Secure Boot is currently enabled.
	SecureBootCurrentBootType_ENABLED SecureBootCurrentBootType = "Enabled"

	// Secure Boot is currently disabled.
	SecureBootCurrentBootType_DISABLED SecureBootCurrentBootType = "Disabled"

)
