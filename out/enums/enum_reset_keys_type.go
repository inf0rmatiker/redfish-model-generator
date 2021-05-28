/* -----------------------------------------------------------------
 * enum_reset_keys_type.go -
 *
 * DMTF Redfish ResetKeysType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ResetKeysType string

const (
	// Reset the contents of all UEFI Secure Boot key databases, including the PK key database, to the default values.
	ResetKeysType_RESET_ALL_KEYS_TO_DEFAULT ResetKeysType = "ResetAllKeysToDefault"

	// Delete the contents of all UEFI Secure Boot key databases, including the PK key database.  This puts the system in Setup Mode.
	ResetKeysType_DELETE_ALL_KEYS ResetKeysType = "DeleteAllKeys"

	// Delete the contents of the PK UEFI Secure Boot database.  This puts the system in Setup Mode.
	ResetKeysType_DELETE_PK ResetKeysType = "DeletePK"

)
