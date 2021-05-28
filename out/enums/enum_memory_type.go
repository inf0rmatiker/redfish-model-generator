/* -----------------------------------------------------------------
 * enum_memory_type.go -
 *
 * DMTF Redfish MemoryType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type MemoryType string

const (
	// The memory device is comprised of volatile memory.
	MemoryType_DRAM MemoryType = "DRAM"

	// The memory device is comprised of volatile memory backed by non-volatile memory.
	MemoryType_NVDIMM_N MemoryType = "NVDIMM_N"

	// The memory device is comprised of non-volatile memory.
	MemoryType_NVDIMM_F MemoryType = "NVDIMM_F"

	// The memory device is comprised of a combination of non-volatile and volatile memory.
	MemoryType_NVDIMM_P MemoryType = "NVDIMM_P"

	// The memory device is an Intel Optane Persistent Memory Module.
	MemoryType_INTEL_OPTANE MemoryType = "IntelOptane"

)
