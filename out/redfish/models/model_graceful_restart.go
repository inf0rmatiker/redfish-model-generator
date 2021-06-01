/* -----------------------------------------------------------------
* graceful_restart.go -
*
* DMTF Redfish GracefulRestart resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// GracefulRestart - Border Gateway Protocol (BGP) graceful restart properties.
// Modeled after DMTF Redfish schema GracefulRestart
type GracefulRestart struct {
	// Border Gateway Protocol (BGP) graceful restart status.
	GracefulRestartEnabled bool `json:"GracefulRestartEnabled,omitempty"`

	// Graceful restart helper mode status.
	HelperModeEnabled bool `json:"HelperModeEnabled,omitempty"`

	// Stale route timer in seconds.
	StaleRoutesTimeSeconds int `json:"StaleRoutesTimeSeconds,omitempty"`

	// Graceful restart timer in seconds.
	TimeSeconds int `json:"TimeSeconds,omitempty"`

}
