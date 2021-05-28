/* -----------------------------------------------------------------
 * enum_processor_type.go -
 *
 * DMTF Redfish ProcessorType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ProcessorType string

const (
	// A CPU.
	ProcessorType_CPU ProcessorType = "CPU"

	// A GPU.
	ProcessorType_GPU ProcessorType = "GPU"

	// An FPGA.
	ProcessorType_FPGA ProcessorType = "FPGA"

	// A DSP.
	ProcessorType_DSP ProcessorType = "DSP"

	// An accelerator.
	ProcessorType_ACCELERATOR ProcessorType = "Accelerator"

	// A core in a processor.
	ProcessorType_CORE ProcessorType = "Core"

	// A thread in a processor.
	ProcessorType_THREAD ProcessorType = "Thread"

	// An OEM-defined processing unit.
	ProcessorType_OEM ProcessorType = "OEM"

)
