/* -----------------------------------------------------------------
* base_speed_priority_settings.go -
*
* DMTF Redfish BaseSpeedPrioritySettings resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The clock speed for a set of cores.
type BaseSpeedPrioritySettings struct {
	// The clock speed to configure the set of cores in MHz.
	BaseSpeedMHz int `json:"BaseSpeedMHz,omitempty"`

	// The number of cores to configure with a specified speed.
	CoreCount int `json:"CoreCount,omitempty"`

	// The identifier of the cores to configure with the specified speed.
	CoreIDs []integer `json:"CoreIDs,omitempty"`

}
