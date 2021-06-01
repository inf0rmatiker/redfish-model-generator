/* -----------------------------------------------------------------
* enum_function_type.go -
*
* DMTF Redfish FunctionType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type FunctionType string

const (
	// A physical PCIe function.
	FunctionType_PHYSICAL FunctionType = "Physical"

	// A virtual PCIe function.
	FunctionType_VIRTUAL FunctionType = "Virtual"

)
