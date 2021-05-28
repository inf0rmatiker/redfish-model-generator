/* -----------------------------------------------------------------
 * enum_memory_media.go -
 *
 * DMTF Redfish MemoryMedia enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type MemoryMedia string

const (
	// DRAM media.
	MemoryMedia_DRAM MemoryMedia = "DRAM"

	// NAND media.
	MemoryMedia_NAND MemoryMedia = "NAND"

	// Intel 3D XPoint media.
	MemoryMedia_INTEL3_DX_POINT MemoryMedia = "Intel3DXPoint"

	// Proprietary media.
	MemoryMedia_PROPRIETARY MemoryMedia = "Proprietary"

)
