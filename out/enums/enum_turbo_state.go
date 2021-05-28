/* -----------------------------------------------------------------
 * enum_turbo_state.go -
 *
 * DMTF Redfish TurboState enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type TurboState string

const (
	// Turbo is enabled.
	TurboState_ENABLED TurboState = "Enabled"

	// Turbo is disabled.
	TurboState_DISABLED TurboState = "Disabled"

)
