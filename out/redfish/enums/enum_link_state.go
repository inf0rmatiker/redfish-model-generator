/* -----------------------------------------------------------------
* enum_link_state.go -
*
* DMTF Redfish LinkState enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type LinkState string

const (
	// This link is enabled.
	LinkState_ENABLED LinkState = "Enabled"

	// This link is disabled.
	LinkState_DISABLED LinkState = "Disabled"

)
