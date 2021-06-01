/* -----------------------------------------------------------------
* enum_boot_source_override_enabled.go -
*
* DMTF Redfish BootSourceOverrideEnabled enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type BootSourceOverrideEnabled string

const (
	// The system boots normally.
	BootSourceOverrideEnabled_DISABLED BootSourceOverrideEnabled = "Disabled"

	// On its next boot cycle, the system boots one time to the boot source override target.  Then, the BootSourceOverrideEnabled value is reset to `Disabled`.
	BootSourceOverrideEnabled_ONCE BootSourceOverrideEnabled = "Once"

	// The system boots to the target specified in the BootSourceOverrideTarget property until this property is `Disabled`.
	BootSourceOverrideEnabled_CONTINUOUS BootSourceOverrideEnabled = "Continuous"

)
