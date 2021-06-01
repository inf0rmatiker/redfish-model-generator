/* -----------------------------------------------------------------
* enum_input_type.go -
*
* DMTF Redfish InputType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type InputType string

const (
	// Alternating Current (AC) input range.
	InputType_AC InputType = "AC"

	// Direct Current (DC) input range.
	InputType_DC InputType = "DC"

)
