/* -----------------------------------------------------------------
* processor_memory.go -
*
* DMTF Redfish ProcessorMemory resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ProcessorMemory - This type describes the memory directly attached or integrated within a processor.
// Modeled after DMTF Redfish schema ProcessorMemory
type ProcessorMemory struct {
	// The memory capacity in MiB.
	CapacityMiB int `json:"CapacityMiB,omitempty"`

	// An indication of whether this memory is integrated within the processor.
	IntegratedMemory bool `json:"IntegratedMemory,omitempty"`

	// The type of memory used by this processor.
	MemoryType ProcessorMemoryType `json:"MemoryType,omitempty"`

	// The operating speed of the memory in MHz.
	SpeedMHz int `json:"SpeedMHz,omitempty"`

}
