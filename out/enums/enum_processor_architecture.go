/* -----------------------------------------------------------------
 * enum_processor_architecture.go -
 *
 * DMTF Redfish ProcessorArchitecture enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ProcessorArchitecture string

const (
	// x86 or x86-64.
	ProcessorArchitecture_X86 ProcessorArchitecture = "x86"

	// Intel Itanium.
	ProcessorArchitecture_IA-64 ProcessorArchitecture = "IA-64"

	// ARM.
	ProcessorArchitecture_ARM ProcessorArchitecture = "ARM"

	// MIPS.
	ProcessorArchitecture_MIPS ProcessorArchitecture = "MIPS"

	// Power.
	ProcessorArchitecture_POWER ProcessorArchitecture = "Power"

	// OEM-defined.
	ProcessorArchitecture_OEM ProcessorArchitecture = "OEM"

)
