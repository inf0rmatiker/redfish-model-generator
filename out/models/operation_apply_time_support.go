/* -----------------------------------------------------------------
* operation_apply_time_support.go -
*
* DMTF Redfish OperationApplyTimeSupport resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The client can request a specific apply time of a create, delete, or action operation of a resource.
type OperationApplyTimeSupport struct {
	// The expiry time of maintenance window in seconds.
	MaintenanceWindowDurationInSeconds int `json:"MaintenanceWindowDurationInSeconds,omitempty"`

	// The location of the maintenance window settings.
	MaintenanceWindowResource idRef `json:"MaintenanceWindowResource,omitempty"`

	// The start time of a maintenance window.
	MaintenanceWindowStartTime string `json:"MaintenanceWindowStartTime,omitempty"`

	// The types of apply times that the client can request when performing a create, delete, or action operation.
	SupportedValues array `json:"SupportedValues"`

}
