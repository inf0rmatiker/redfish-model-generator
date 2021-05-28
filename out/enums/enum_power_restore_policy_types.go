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
	// The system always powers on when power is applied.
	PowerRestorePolicyTypes_ALWAYS_ON PowerRestorePolicyTypes = "AlwaysOn"

	// The system always remains powered off when power is applied.
	PowerRestorePolicyTypes_ALWAYS_OFF PowerRestorePolicyTypes = "AlwaysOff"

	// The system returns to its last on or off power state when power is applied.
	PowerRestorePolicyTypes_LAST_STATE PowerRestorePolicyTypes = "LastState"

)
