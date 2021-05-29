/* -----------------------------------------------------------------
* core_metrics.go -
*
* DMTF Redfish CoreMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The processor core metrics.
type CoreMetrics struct {
	// The C-state residency of this core in the processor.
	CStateResidency array `json:"CStateResidency,omitempty"`

	// The cache metrics of this core in the processor.
	CoreCache array `json:"CoreCache,omitempty"`

	// The processor core identifier.
	CoreId string `json:"CoreId,omitempty"`

	// The number of stalled cycles due to I/O operations.
	IOStallCount float64 `json:"IOStallCount,omitempty"`

	// The number of instructions per clock cycle of this core.
	InstructionsPerCycle float64 `json:"InstructionsPerCycle,omitempty"`

	// The number of stalled cycles due to memory operations.
	MemoryStallCount float64 `json:"MemoryStallCount,omitempty"`

	// The unhalted cycles count of this core.
	UnhaltedCycles float64 `json:"UnhaltedCycles,omitempty"`

}
