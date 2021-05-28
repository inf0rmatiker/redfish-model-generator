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
	// UEFI Secure Boot is currently in Setup Mode.
	SecureBootModeType_SETUP_MODE SecureBootModeType = "SetupMode"

	// UEFI Secure Boot is currently in User Mode.
	SecureBootModeType_USER_MODE SecureBootModeType = "UserMode"

	// UEFI Secure Boot is currently in Audit Mode.
	SecureBootModeType_AUDIT_MODE SecureBootModeType = "AuditMode"

	// UEFI Secure Boot is currently in Deployed Mode.
	SecureBootModeType_DEPLOYED_MODE SecureBootModeType = "DeployedMode"

)
