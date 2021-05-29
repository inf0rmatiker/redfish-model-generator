/* -----------------------------------------------------------------
* power_management_policy.go -
*
* DMTF Redfish PowerManagementPolicy resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Power management policy information.
type PowerManagementPolicy struct {
	// Average power budget, in milliwatts.
	AveragePowerBudgetMilliWatts int `json:"AveragePowerBudgetMilliWatts,omitempty"`

	// Maximum TDP in milliwatts.
	MaxTDPMilliWatts int `json:"MaxTDPMilliWatts,omitempty"`

	// Peak power budget, in milliwatts.
	PeakPowerBudgetMilliWatts int `json:"PeakPowerBudgetMilliWatts,omitempty"`

	// An indication of whether the power management policy is enabled.
	PolicyEnabled bool `json:"PolicyEnabled,omitempty"`

}
