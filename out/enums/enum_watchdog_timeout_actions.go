/* -----------------------------------------------------------------
 * enum_watchdog_timeout_actions.go -
 *
 * DMTF Redfish WatchdogTimeoutActions enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type WatchdogTimeoutActions string

const (
	// No action taken.
	WatchdogTimeoutActions_NONE WatchdogTimeoutActions = "None"

	// Reset the system.
	WatchdogTimeoutActions_RESET_SYSTEM WatchdogTimeoutActions = "ResetSystem"

	// Power cycle the system.
	WatchdogTimeoutActions_POWER_CYCLE WatchdogTimeoutActions = "PowerCycle"

	// Power down the system.
	WatchdogTimeoutActions_POWER_DOWN WatchdogTimeoutActions = "PowerDown"

	// Perform an OEM-defined action.
	WatchdogTimeoutActions_OEM WatchdogTimeoutActions = "OEM"

)
