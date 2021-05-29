/* -----------------------------------------------------------------
* c_state_residency.go -
*
* DMTF Redfish CStateResidency resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The C-state residency of the processor.
type CStateResidency struct {
	// The C-state level, such as C0, C1, or C2.
	Level string `json:"Level,omitempty"`

	// The percentage of time that the processor or core has spent in this particular level of C-state.
	ResidencyPercent float64 `json:"ResidencyPercent,omitempty"`

}
