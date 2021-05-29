/* -----------------------------------------------------------------
* preferred_apply_time.go -
*
* DMTF Redfish PreferredApplyTime resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The preferred time to apply configuration settings.
type PreferredApplyTime struct {
	// The time when to apply the settings.
	ApplyTime ApplyTime `json:"ApplyTime,omitempty"`

	// The expiry time of maintenance window in seconds.
	MaintenanceWindowDurationInSeconds int `json:"MaintenanceWindowDurationInSeconds,omitempty"`

	// The start time of a maintenance window.
	MaintenanceWindowStartTime string `json:"MaintenanceWindowStartTime,omitempty"`

}
