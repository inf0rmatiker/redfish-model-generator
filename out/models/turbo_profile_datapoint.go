/* -----------------------------------------------------------------
* turbo_profile_datapoint.go -
*
* DMTF Redfish TurboProfileDatapoint resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The turbo profile for a set of active cores.
type TurboProfileDatapoint struct {
	// The number of active cores to be configured with the specified maximum clock speed.
	ActiveCoreCount int `json:"ActiveCoreCount,omitempty"`

	// The maximum turbo clock speed that correspond to the number of active cores in MHz.
	MaxSpeedMHz int `json:"MaxSpeedMHz,omitempty"`

}
