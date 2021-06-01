/* -----------------------------------------------------------------
* boot_targets.go -
*
* DMTF Redfish BootTargets resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// BootTargets - A Fibre Channel boot target configured for a network device function.
// Modeled after DMTF Redfish schema BootTargets
type BootTargets struct {
	// The relative priority for this entry in the boot targets array.
	BootPriority int `json:"BootPriority,omitempty"`

	// The Logical Unit Number (LUN) ID to boot from on the device referred to by the corresponding WWPN.
	LUNID string `json:"LUNID,omitempty"`

	// The World-Wide Port Name to boot from.
	WWPN string `json:"WWPN,omitempty"`

}
