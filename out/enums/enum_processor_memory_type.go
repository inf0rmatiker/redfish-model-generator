/* -----------------------------------------------------------------
 * enum_processor_memory_type.go -
 *
 * DMTF Redfish ProcessorMemoryType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ProcessorMemoryType string

const (
	// L1 cache.
	ProcessorMemoryType_L1_CACHE ProcessorMemoryType = "L1Cache"

	// L2 cache.
	ProcessorMemoryType_L2_CACHE ProcessorMemoryType = "L2Cache"

	// L3 cache.
	ProcessorMemoryType_L3_CACHE ProcessorMemoryType = "L3Cache"

	// L4 cache.
	ProcessorMemoryType_L4_CACHE ProcessorMemoryType = "L4Cache"

	// L5 cache.
	ProcessorMemoryType_L5_CACHE ProcessorMemoryType = "L5Cache"

	// L6 cache.
	ProcessorMemoryType_L6_CACHE ProcessorMemoryType = "L6Cache"

	// L7 cache.
	ProcessorMemoryType_L7_CACHE ProcessorMemoryType = "L7Cache"

	// High Bandwidth Memory.
	ProcessorMemoryType_HBM1 ProcessorMemoryType = "HBM1"

	// The second generation of High Bandwidth Memory.
	ProcessorMemoryType_HBM2 ProcessorMemoryType = "HBM2"

	// The third generation of High Bandwidth Memory.
	ProcessorMemoryType_HBM3 ProcessorMemoryType = "HBM3"

	// Synchronous graphics RAM.
	ProcessorMemoryType_SGRAM ProcessorMemoryType = "SGRAM"

	// Synchronous graphics random-access memory.
	ProcessorMemoryType_GDDR ProcessorMemoryType = "GDDR"

	// Double data rate type two synchronous graphics random-access memory.
	ProcessorMemoryType_GDDR2 ProcessorMemoryType = "GDDR2"

	// Double data rate type three synchronous graphics random-access memory.
	ProcessorMemoryType_GDDR3 ProcessorMemoryType = "GDDR3"

	// Double data rate type four synchronous graphics random-access memory.
	ProcessorMemoryType_GDDR4 ProcessorMemoryType = "GDDR4"

	// Double data rate type five synchronous graphics random-access memory.
	ProcessorMemoryType_GDDR5 ProcessorMemoryType = "GDDR5"

	// Double data rate type five X synchronous graphics random-access memory.
	ProcessorMemoryType_GDDR5_X ProcessorMemoryType = "GDDR5X"

	// Double data rate type six synchronous graphics random-access memory.
	ProcessorMemoryType_GDDR6 ProcessorMemoryType = "GDDR6"

	// Double data rate synchronous dynamic random-access memory.
	ProcessorMemoryType_DDR ProcessorMemoryType = "DDR"

	// Double data rate type two synchronous dynamic random-access memory.
	ProcessorMemoryType_DDR2 ProcessorMemoryType = "DDR2"

	// Double data rate type three synchronous dynamic random-access memory.
	ProcessorMemoryType_DDR3 ProcessorMemoryType = "DDR3"

	// Double data rate type four synchronous dynamic random-access memory.
	ProcessorMemoryType_DDR4 ProcessorMemoryType = "DDR4"

	// Double data rate type five synchronous dynamic random-access memory.
	ProcessorMemoryType_DDR5 ProcessorMemoryType = "DDR5"

	// Synchronous dynamic random-access memory.
	ProcessorMemoryType_SDRAM ProcessorMemoryType = "SDRAM"

	// Static random-access memory.
	ProcessorMemoryType_SRAM ProcessorMemoryType = "SRAM"

	// Flash memory.
	ProcessorMemoryType_FLASH ProcessorMemoryType = "Flash"

	// OEM-defined.
	ProcessorMemoryType_OEM ProcessorMemoryType = "OEM"

)
