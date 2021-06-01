/* -----------------------------------------------------------------
* enum_memory_mirroring.go -
*
* DMTF Redfish MemoryMirroring enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type MemoryMirroring string

const (
	// The system supports DIMM mirroring at the system level.  Individual DIMMs are not paired for mirroring in this mode.
	MemoryMirroring_SYSTEM MemoryMirroring = "System"

	// The system supports DIMM mirroring at the DIMM level.  Individual DIMMs can be mirrored.
	MemoryMirroring_DIMM MemoryMirroring = "DIMM"

	// The system supports a hybrid mirroring at the system and DIMM levels.  Individual DIMMs can be mirrored.
	MemoryMirroring_HYBRID MemoryMirroring = "Hybrid"

	// The system does not support DIMM mirroring.
	MemoryMirroring_NONE MemoryMirroring = "None"

)
