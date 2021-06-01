/* -----------------------------------------------------------------
* maintenance_window.go -
*
* DMTF Redfish MaintenanceWindow resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MaintenanceWindow - The maintenance window assignment for applying settings or operations to a Resource.
// Modeled after DMTF Redfish schema MaintenanceWindow
type MaintenanceWindow struct {
	// The expiry time of maintenance window in seconds.
	MaintenanceWindowDurationInSeconds int `json:"MaintenanceWindowDurationInSeconds"`

	// The start time of a maintenance window.
	MaintenanceWindowStartTime string `json:"MaintenanceWindowStartTime"`

}
