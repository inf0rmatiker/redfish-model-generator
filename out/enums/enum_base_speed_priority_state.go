/* -----------------------------------------------------------------
 * enum_base_speed_priority_state.go -
 *
 * DMTF Redfish BaseSpeedPriorityState enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type BaseSpeedPriorityState string

const (
	// Base speed priority is enabled.
	BaseSpeedPriorityState_ENABLED BaseSpeedPriorityState = "Enabled"

	// Base speed priority is disabled.
	BaseSpeedPriorityState_DISABLED BaseSpeedPriorityState = "Disabled"

)
