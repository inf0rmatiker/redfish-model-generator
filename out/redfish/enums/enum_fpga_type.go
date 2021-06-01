/* -----------------------------------------------------------------
* enum_fpga_type.go -
*
* DMTF Redfish FpgaType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type FpgaType string

const (
	// The FPGA device integrated with other processor in the single chip.
	FpgaType_INTEGRATED FpgaType = "Integrated"

	// The discrete FPGA device.
	FpgaType_DISCRETE FpgaType = "Discrete"

)
