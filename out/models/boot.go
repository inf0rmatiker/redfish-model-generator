/* -----------------------------------------------------------------
* boot.go -
*
* DMTF Redfish Boot resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The boot information for this resource.
type Boot struct {
	// Ordered array of boot source aliases representing the persistent boot order associated with this computer system.
	AliasBootOrder array `json:"AliasBootOrder,omitempty"`

	// The BootOptionReference of the Boot Option to perform a one-time boot from when BootSourceOverrideTarget is `UefiBootNext`.
	BootNext string `json:"BootNext,omitempty"`

	// The link to the collection of the UEFI boot options associated with this computer system.
	BootOptions BootOptionCollection `json:"BootOptions,omitempty"`

	// An array of BootOptionReference strings that represent the persistent boot order for with this computer system.
	BootOrder []string `json:"BootOrder,omitempty"`

	// The name of the boot order property that the system uses for the persistent boot order.
	BootOrderPropertySelection BootOrderTypes `json:"BootOrderPropertySelection,omitempty"`

	// The state of the boot source override feature.
	BootSourceOverrideEnabled BootSourceOverrideEnabled `json:"BootSourceOverrideEnabled,omitempty"`

	// The BIOS boot mode to use when the system boots from the BootSourceOverrideTarget boot source.
	BootSourceOverrideMode BootSourceOverrideMode `json:"BootSourceOverrideMode,omitempty"`

	// The current boot source to use at the next boot instead of the normal boot device, if BootSourceOverrideEnabled is `true`.
	BootSourceOverrideTarget BootSource `json:"BootSourceOverrideTarget,omitempty"`

	// The link to a collection of certificates used for booting through HTTPS by this computer system.
	Certificates CertificateCollection `json:"Certificates,omitempty"`

	// The URI to boot from when BootSourceOverrideTarget is set to `UefiHttp`.
	HttpBootUri string `json:"HttpBootUri,omitempty"`

	// The UEFI device path of the device from which to boot when BootSourceOverrideTarget is `UefiTarget`.
	UefiTargetBootSourceOverride string `json:"UefiTargetBootSourceOverride,omitempty"`

}
