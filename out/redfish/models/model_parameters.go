/* -----------------------------------------------------------------
* parameters.go -
*
* DMTF Redfish Parameters resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Parameters - A parameter associated with the specified Redfish Action.
// Modeled after DMTF Redfish schema Parameters
type Parameters struct {
	// A list of values for this parameter supported by this Action target.
	AllowableValues []string `json:"AllowableValues,omitempty"`

	// The JSON property type used for this parameter.
	DataType ParameterTypes `json:"DataType,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OData Type of an object-based parameter.
	ObjectDataType string `json:"ObjectDataType,omitempty"`

	// Indicates whether the parameter is required to perform this Action.
	Required bool `json:"Required,omitempty"`

}
