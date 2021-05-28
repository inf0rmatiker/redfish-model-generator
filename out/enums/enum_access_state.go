/* -----------------------------------------------------------------
 * enum_access_state.go -
 *
 * DMTF Redfish AccessState enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type AccessState string

const (
	// The resource is in an active and optimized state.
	AccessState_OPTIMIZED AccessState = "Optimized"

	// The resource is in an active and non-optimized state.
	AccessState_NON_OPTIMIZED AccessState = "NonOptimized"

	// The resource is in a standby state.
	AccessState_STANDBY AccessState = "Standby"

	// The resource is in an unavailable state.
	AccessState_UNAVAILABLE AccessState = "Unavailable"

	// The resource is transitioning to a new state.
	AccessState_TRANSITIONING AccessState = "Transitioning"

)
