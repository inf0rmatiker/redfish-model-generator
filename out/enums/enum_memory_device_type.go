/* -----------------------------------------------------------------
 * enum_memory_device_type.go -
 *
 * DMTF Redfish MemoryDeviceType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type MemoryDeviceType string

const (
	// DDR.
	MemoryDeviceType_DDR MemoryDeviceType = "DDR"

	// DDR2.
	MemoryDeviceType_DDR2 MemoryDeviceType = "DDR2"

	// DDR3.
	MemoryDeviceType_DDR3 MemoryDeviceType = "DDR3"

	// DDR4.
	MemoryDeviceType_DDR4 MemoryDeviceType = "DDR4"

	// DDR4 SDRAM.
	MemoryDeviceType_DDR4_SDRAM MemoryDeviceType = "DDR4_SDRAM"

	// DDR4E SDRAM.
	MemoryDeviceType_DDR4_E_SDRAM MemoryDeviceType = "DDR4E_SDRAM"

	// LPDDR4 SDRAM.
	MemoryDeviceType_LPDDR4_SDRAM MemoryDeviceType = "LPDDR4_SDRAM"

	// DDR3 SDRAM.
	MemoryDeviceType_DDR3_SDRAM MemoryDeviceType = "DDR3_SDRAM"

	// LPDDR3 SDRAM.
	MemoryDeviceType_LPDDR3_SDRAM MemoryDeviceType = "LPDDR3_SDRAM"

	// DDR2 SDRAM.
	MemoryDeviceType_DDR2_SDRAM MemoryDeviceType = "DDR2_SDRAM"

	// DDR2 SDRAM FB_DIMM.
	MemoryDeviceType_DDR2_SDRAM_FB_DIMM MemoryDeviceType = "DDR2_SDRAM_FB_DIMM"

	// DDR2 SDRAM FB_DIMM PROBE.
	MemoryDeviceType_DDR2_SDRAM_FB_DIMM_PROBE MemoryDeviceType = "DDR2_SDRAM_FB_DIMM_PROBE"

	// DDR SGRAM.
	MemoryDeviceType_DDR_SGRAM MemoryDeviceType = "DDR_SGRAM"

	// DDR SDRAM.
	MemoryDeviceType_DDR_SDRAM MemoryDeviceType = "DDR_SDRAM"

	// ROM.
	MemoryDeviceType_ROM MemoryDeviceType = "ROM"

	// SDRAM.
	MemoryDeviceType_SDRAM MemoryDeviceType = "SDRAM"

	// EDO.
	MemoryDeviceType_EDO MemoryDeviceType = "EDO"

	// Fast Page Mode.
	MemoryDeviceType_FAST_PAGE_MODE MemoryDeviceType = "FastPageMode"

	// Pipelined Nibble.
	MemoryDeviceType_PIPELINED_NIBBLE MemoryDeviceType = "PipelinedNibble"

	// Logical Non-volatile device.
	MemoryDeviceType_LOGICAL MemoryDeviceType = "Logical"

	// High Bandwidth Memory.
	MemoryDeviceType_HBM MemoryDeviceType = "HBM"

	// The second generation of High Bandwidth Memory.
	MemoryDeviceType_HBM2 MemoryDeviceType = "HBM2"

)
