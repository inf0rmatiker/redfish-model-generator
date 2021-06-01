/* -----------------------------------------------------------------
* c_state_residency.go -
*
* DMTF Redfish CStateResidency resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// CStateResidency - The C-state residency of the processor.
// Modeled after DMTF Redfish schema CStateResidency
type CStateResidency struct {
	// The level of C-state, e.g. C0, C1, C2.
	Level string `json:"Level,omitempty"`

	// The percentage of time that the processor or core has spent in this particular level of C-state.
	ResidencyPercent float64 `json:"ResidencyPercent,omitempty"`

}
