/* -----------------------------------------------------------------
* enum_breaker_states.go -
*
* DMTF Redfish BreakerStates enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type BreakerStates string

const (
	// The breaker is powered on.
	BreakerStates_NORMAL BreakerStates = "Normal"

	// The breaker has been tripped.
	BreakerStates_TRIPPED BreakerStates = "Tripped"

	// The breaker is off.
	BreakerStates_OFF BreakerStates = "Off"

)
