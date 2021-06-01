/* -----------------------------------------------------------------
* enum_secure_boot_mode_type.go -
*
* DMTF Redfish SecureBootModeType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type SecureBootModeType string

const (
	// Secure Boot is currently in Setup Mode.
	SecureBootModeType_SETUP_MODE SecureBootModeType = "SetupMode"

	// Secure Boot is currently in User Mode.
	SecureBootModeType_USER_MODE SecureBootModeType = "UserMode"

	// Secure Boot is currently in Audit Mode.
	SecureBootModeType_AUDIT_MODE SecureBootModeType = "AuditMode"

	// Secure Boot is currently in Deployed Mode.
	SecureBootModeType_DEPLOYED_MODE SecureBootModeType = "DeployedMode"

)
