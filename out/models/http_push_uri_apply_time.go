/* -----------------------------------------------------------------
* http_push_uri_apply_time.go -
*
* DMTF Redfish HttpPushUriApplyTime resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The settings for when to apply HttpPushUri-provided software.
type HttpPushUriApplyTime struct {
	// The time when to apply the HttpPushUri-provided software update.
	ApplyTime ApplyTime `json:"ApplyTime,omitempty"`

	// The expiry time, in seconds, of the maintenance window.
	MaintenanceWindowDurationInSeconds int `json:"MaintenanceWindowDurationInSeconds,omitempty"`

	// The start time of a maintenance window.
	MaintenanceWindowStartTime string `json:"MaintenanceWindowStartTime,omitempty"`

}
