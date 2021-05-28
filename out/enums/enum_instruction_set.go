/* -----------------------------------------------------------------
 * enum_instruction_set.go -
 *
 * DMTF Redfish InstructionSet enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type InstructionSet string

const (
	// x86 32-bit.
	InstructionSet_X86 InstructionSet = "x86"

	// x86 64-bit.
	InstructionSet_X86-64 InstructionSet = "x86-64"

	// Intel IA-64.
	InstructionSet_IA-64 InstructionSet = "IA-64"

	// ARM 32-bit.
	InstructionSet_ARM-A32 InstructionSet = "ARM-A32"

	// ARM 64-bit.
	InstructionSet_ARM-A64 InstructionSet = "ARM-A64"

	// MIPS 32-bit.
	InstructionSet_MIPS32 InstructionSet = "MIPS32"

	// MIPS 64-bit.
	InstructionSet_MIPS64 InstructionSet = "MIPS64"

	// PowerISA-64 or PowerISA-32.
	InstructionSet_POWER_ISA InstructionSet = "PowerISA"

	// OEM-defined.
	InstructionSet_OEM InstructionSet = "OEM"

)
