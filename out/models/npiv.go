/* -----------------------------------------------------------------
* npiv.go -
*
* DMTF Redfish NPIV resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// N_Port ID Virtualization (NPIV) capabilities for a controller.
type NPIV struct {
	// The maximum number of N_Port ID Virtualization (NPIV) logins allowed simultaneously from all ports on this controller.
	MaxDeviceLogins int `json:"MaxDeviceLogins,omitempty"`

	// The maximum number of N_Port ID Virtualization (NPIV) logins allowed per physical port on this controller.
	MaxPortLogins int `json:"MaxPortLogins,omitempty"`

}
