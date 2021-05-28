/* -----------------------------------------------------------------
 * enum_apply_time.go -
 *
 * DMTF Redfish ApplyTime enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ApplyTime string

const (
	// Apply immediately.
	ApplyTime_IMMEDIATE ApplyTime = "Immediate"

	// Apply on a reset.
	ApplyTime_ON_RESET ApplyTime = "OnReset"

	// Apply during an administrator-specified maintenance window.
	ApplyTime_AT_MAINTENANCE_WINDOW_START ApplyTime = "AtMaintenanceWindowStart"

	// Apply after a reset but within an administrator-specified maintenance window.
	ApplyTime_IN_MAINTENANCE_WINDOW_ON_RESET ApplyTime = "InMaintenanceWindowOnReset"

)
