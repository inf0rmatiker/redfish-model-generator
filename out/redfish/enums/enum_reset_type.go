/* -----------------------------------------------------------------
* enum_reset_type.go -
*
* DMTF Redfish ResetType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ResetType string

const (
	// Turn on the unit.
	ResetType_ON ResetType = "On"

	// Turn off the unit immediately (non-graceful shutdown).
	ResetType_FORCE_OFF ResetType = "ForceOff"

	// Shut down gracefully and power off.
	ResetType_GRACEFUL_SHUTDOWN ResetType = "GracefulShutdown"

	// Shut down gracefully and restart the system.
	ResetType_GRACEFUL_RESTART ResetType = "GracefulRestart"

	// Shut down immediately and non-gracefully and restart the system.
	ResetType_FORCE_RESTART ResetType = "ForceRestart"

	// Generate a diagnostic interrupt, which is usually an NMI on x86 systems, to stop normal operations, complete diagnostic actions, and, typically, halt the system.
	ResetType_NMI ResetType = "Nmi"

	// Turn on the unit immediately.
	ResetType_FORCE_ON ResetType = "ForceOn"

	// Simulate the pressing of the physical power button on this unit.
	ResetType_PUSH_POWER_BUTTON ResetType = "PushPowerButton"

	// Power cycle the unit.  Behaves like a full power removal, followed by a power restore to the resource.
	ResetType_POWER_CYCLE ResetType = "PowerCycle"

)
