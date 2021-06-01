/* -----------------------------------------------------------------
* npiv.go -
*
* DMTF Redfish NPIV resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// NPIV - N_Port ID Virtualization (NPIV) capabilities for a controller.
// Modeled after DMTF Redfish schema NPIV
type NPIV struct {
	// The maximum number of N_Port ID Virtualization (NPIV) logins allowed simultaneously from all ports on this controller.
	MaxDeviceLogins int `json:"MaxDeviceLogins,omitempty"`

	// The maximum number of N_Port ID Virtualization (NPIV) logins allowed per physical port on this controller.
	MaxPortLogins int `json:"MaxPortLogins,omitempty"`

}
