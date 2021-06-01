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
	// Reset the content of all UEFI Secure Boot key databases (PK, KEK, DB, DBX) to their default values.
	ResetKeysType_RESET_ALL_KEYS_TO_DEFAULT ResetKeysType = "ResetAllKeysToDefault"

	// Delete the content of all UEFI Secure Boot key databases (PK, KEK, DB, DBX). This puts the system in Setup Mode.
	ResetKeysType_DELETE_ALL_KEYS ResetKeysType = "DeleteAllKeys"

	// Delete the content of the PK UEFI Secure Boot database. This puts the system in Setup Mode.
	ResetKeysType_DELETE_PK ResetKeysType = "DeletePK"

)
