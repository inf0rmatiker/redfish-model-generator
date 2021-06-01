/* -----------------------------------------------------------------
* enum_operation_apply_time.go -
*
* DMTF Redfish OperationApplyTime enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type OperationApplyTime string

const (
	// The requested operation is applied immediately.
	OperationApplyTime_IMMEDIATE OperationApplyTime = "Immediate"

	// The requested operation is applied on a reset.
	OperationApplyTime_ON_RESET OperationApplyTime = "OnReset"

	// The requested operation is applied within the administrator-specified maintenance window.
	OperationApplyTime_AT_MAINTENANCE_WINDOW_START OperationApplyTime = "AtMaintenanceWindowStart"

	// The requested operation is applied after a reset but within the administrator-specified maintenance window.
	OperationApplyTime_IN_MAINTENANCE_WINDOW_ON_RESET OperationApplyTime = "InMaintenanceWindowOnReset"

	// The requested operation is applied when the StartUpdate action of the update service is invoked.
	OperationApplyTime_ON_START_UPDATE_REQUEST OperationApplyTime = "OnStartUpdateRequest"

)
