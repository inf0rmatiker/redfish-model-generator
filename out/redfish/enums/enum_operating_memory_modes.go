/* -----------------------------------------------------------------
* enum_operating_memory_modes.go -
*
* DMTF Redfish OperatingMemoryModes enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type OperatingMemoryModes string

const (
	// Volatile memory.
	OperatingMemoryModes_VOLATILE OperatingMemoryModes = "Volatile"

	// Persistent memory, byte-accessible through system address space.
	OperatingMemoryModes_PMEM OperatingMemoryModes = "PMEM"

	// Block-accessible system memory.
	OperatingMemoryModes_BLOCK OperatingMemoryModes = "Block"

)
