/* -----------------------------------------------------------------
* boot.go -
*
* DMTF Redfish Boot resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Boot - The boot information for this resource.
// Modeled after DMTF Redfish schema Boot
type Boot struct {
	// The state of the boot source override feature.
	BootSourceOverrideEnabled BootSourceOverrideEnabled `json:"BootSourceOverrideEnabled,omitempty"`

	// The BIOS boot mode to use when the system boots from the BootSourceOverrideTarget boot source.
	BootSourceOverrideMode BootSourceOverrideMode `json:"BootSourceOverrideMode,omitempty"`

	// The current boot source to use at the next boot instead of the normal boot device, if BootSourceOverrideEnabled is `true`.
	BootSourceOverrideTarget map[string]interface{} `json:"BootSourceOverrideTarget,omitempty"`

	// The UEFI device path of the device from which to boot when BootSourceOverrideTarget is `UefiTarget`.
	UefiTargetBootSourceOverride string `json:"UefiTargetBootSourceOverride,omitempty"`

}
