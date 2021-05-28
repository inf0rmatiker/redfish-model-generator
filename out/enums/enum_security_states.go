/* -----------------------------------------------------------------
 * enum_security_states.go -
 *
 * DMTF Redfish SecurityStates enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type SecurityStates string

const (
	// Secure mode is enabled and access to the data is allowed.
	SecurityStates_ENABLED SecurityStates = "Enabled"

	// Secure mode is disabled.
	SecurityStates_DISABLED SecurityStates = "Disabled"

	// Secure mode is enabled and access to the data is unlocked.
	SecurityStates_UNLOCKED SecurityStates = "Unlocked"

	// Secure mode is enabled and access to the data is locked.
	SecurityStates_LOCKED SecurityStates = "Locked"

	// Secure state is frozen and cannot be modified until reset.
	SecurityStates_FROZEN SecurityStates = "Frozen"

	// Number of attempts to unlock the memory exceeded limit.
	SecurityStates_PASSPHRASELIMIT SecurityStates = "Passphraselimit"

)
