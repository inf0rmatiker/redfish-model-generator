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
	// The components within the chassis have power.
	PowerState_ON PowerState = "On"

	// The components within the chassis have no power, except some components might continue to have AUX power, such as the management controller.
	PowerState_OFF PowerState = "Off"

	// A temporary state between off and on.  The components within the chassis can take time to process the power on action.
	PowerState_POWERING_ON PowerState = "PoweringOn"

	// A temporary state between on and off.  The components within the chassis can take time to process the power off action.
	PowerState_POWERING_OFF PowerState = "PoweringOff"

)
