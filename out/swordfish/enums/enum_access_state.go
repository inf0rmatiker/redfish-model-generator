/* -----------------------------------------------------------------
* enum_access_state.go -
*
* SNIA Swordfish AccessState enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type AccessState string

const (
	// The endpoint(s) are in an Active/Optimized state.
	AccessState_OPTIMIZED AccessState = "Optimized"

	// The endpoint(s) are in an Active/NonOptimized state.
	AccessState_NON_OPTIMIZED AccessState = "NonOptimized"

	// The endpoint(s) are in a Standby state.
	AccessState_STANDBY AccessState = "Standby"

	// The endpoint(s) are unavailable.
	AccessState_UNAVAILABLE AccessState = "Unavailable"

	// The endpoint(s) are transitioning to a new AccessState.
	AccessState_TRANSITIONING AccessState = "Transitioning"

)
