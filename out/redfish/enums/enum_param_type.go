/* -----------------------------------------------------------------
* enum_param_type.go -
*
* DMTF Redfish ParamType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type ParamType string

const (
	// The argument is a string.
	ParamType_STRING ParamType = "string"

	// The argument is a number.
	ParamType_NUMBER ParamType = "number"

)
