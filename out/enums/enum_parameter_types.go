/* -----------------------------------------------------------------
 * enum_parameter_types.go -
 *
 * DMTF Redfish ParameterTypes enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type ParameterTypes string

const (
	// A boolean.
	ParameterTypes_BOOLEAN ParameterTypes = "Boolean"

	// A number.
	ParameterTypes_NUMBER ParameterTypes = "Number"

	// An array of numbers.
	ParameterTypes_NUMBER_ARRAY ParameterTypes = "NumberArray"

	// A string.
	ParameterTypes_STRING ParameterTypes = "String"

	// An array of strings.
	ParameterTypes_STRING_ARRAY ParameterTypes = "StringArray"

	// An embedded JSON object.
	ParameterTypes_OBJECT ParameterTypes = "Object"

	// An array of JSON objects.
	ParameterTypes_OBJECT_ARRAY ParameterTypes = "ObjectArray"

)
