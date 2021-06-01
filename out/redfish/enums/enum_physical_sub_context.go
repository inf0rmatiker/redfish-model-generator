/* -----------------------------------------------------------------
* enum_physical_sub_context.go -
*
* DMTF Redfish PhysicalSubContext enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type PhysicalSubContext string

const (
	// The input.
	PhysicalSubContext_INPUT PhysicalSubContext = "Input"

	// The output.
	PhysicalSubContext_OUTPUT PhysicalSubContext = "Output"

)
