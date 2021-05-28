/* -----------------------------------------------------------------
 * enum_reset_to_defaults_type.go -
 *
 * DMTF Redfish ResetToDefaultsType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ResetToDefaultsType string

const (
	// Reset all settings to factory defaults.
	ResetToDefaultsType_RESET_ALL ResetToDefaultsType = "ResetAll"

	// Reset all settings except network and local user names/passwords to factory defaults.
	ResetToDefaultsType_PRESERVE_NETWORK_AND_USERS ResetToDefaultsType = "PreserveNetworkAndUsers"

	// Reset all settings except network settings to factory defaults.
	ResetToDefaultsType_PRESERVE_NETWORK ResetToDefaultsType = "PreserveNetwork"

)
