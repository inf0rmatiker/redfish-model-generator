/* -----------------------------------------------------------------
* boot_targets.go -
*
* DMTF Redfish BootTargets resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A Fibre Channel boot target configured for a network device function.
type BootTargets struct {
	// The relative priority for this entry in the boot targets array.
	BootPriority int `json:"BootPriority,omitempty"`

	// The logical unit number (LUN) ID from which to boot on the device to which the corresponding WWPN refers.
	LUNID string `json:"LUNID,omitempty"`

	// The World Wide Port Name (WWPN) from which to boot.
	WWPN string `json:"WWPN,omitempty"`

}
