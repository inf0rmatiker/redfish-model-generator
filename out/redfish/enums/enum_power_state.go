/* -----------------------------------------------------------------
* enum_power_state.go -
*
* DMTF Redfish PowerState enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type PowerState string

const (
	// The state is powered On.
	PowerState_ON PowerState = "On"

	// The state is powered Off.
	PowerState_OFF PowerState = "Off"

	// A temporary state between Off and On.
	PowerState_POWERING_ON PowerState = "PoweringOn"

	// A temporary state between On and Off.
	PowerState_POWERING_OFF PowerState = "PoweringOff"

)
