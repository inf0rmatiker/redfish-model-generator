/* -----------------------------------------------------------------
* enum_power_restore_policy_types.go -
*
* DMTF Redfish PowerRestorePolicyTypes enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type PowerRestorePolicyTypes string

const (
	// Always power on when external power is applied.
	PowerRestorePolicyTypes_ALWAYS_ON PowerRestorePolicyTypes = "AlwaysOn"

	// Always remain powered off when external power is applied.
	PowerRestorePolicyTypes_ALWAYS_OFF PowerRestorePolicyTypes = "AlwaysOff"

	// Return to the last power state (on or off) when external power is applied.
	PowerRestorePolicyTypes_LAST_STATE PowerRestorePolicyTypes = "LastState"

)
