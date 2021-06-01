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
	// The system is powered on.
	PowerState_ON PowerState = "On"

	// The system is powered off, although some components might continue to have AUX power such as management controller.
	PowerState_OFF PowerState = "Off"

	// A temporary state between off and on.  This temporary state can be very short.
	PowerState_POWERING_ON PowerState = "PoweringOn"

	// A temporary state between on and off.  The power off action can take time while the OS is in the shutdown process.
	PowerState_POWERING_OFF PowerState = "PoweringOff"

)
