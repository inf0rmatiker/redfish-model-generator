/* -----------------------------------------------------------------
* max_prefix.go -
*
* DMTF Redfish MaxPrefix resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Border Gateway Protocol (BGP) max prefix properties.
type MaxPrefix struct {
	// Maximum prefix number.
	MaxPrefixNumber int `json:"MaxPrefixNumber,omitempty"`

	// Border Gateway Protocol (BGP) restart timer in seconds.
	RestartTimerSeconds int `json:"RestartTimerSeconds,omitempty"`

	// Shutdown threshold status.
	ShutdownThresholdPercentage float64 `json:"ShutdownThresholdPercentage,omitempty"`

	// Threshold warning only status.
	ThresholdWarningOnlyEnabled bool `json:"ThresholdWarningOnlyEnabled,omitempty"`

}
