/* -----------------------------------------------------------------
* enum_boot_source_override_mode.go -
*
* DMTF Redfish BootSourceOverrideMode enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type BootSourceOverrideMode string

const (
	// The system boots in non-UEFI boot mode to the boot source override target.
	BootSourceOverrideMode_LEGACY BootSourceOverrideMode = "Legacy"

	// The system boots in UEFI boot mode to the boot source override target.
	BootSourceOverrideMode_UEFI BootSourceOverrideMode = "UEFI"

)
