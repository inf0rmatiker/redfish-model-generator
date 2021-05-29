/* -----------------------------------------------------------------
* parameters.go -
*
* DMTF Redfish Parameters resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The information about a parameter included in a Redfish action for this Resource.
type Parameters struct {
	// The allowable values for this parameter as applied to this action target.
	AllowableValues []string `json:"AllowableValues,omitempty"`

	// The JSON property type for this parameter.
	DataType ParameterTypes `json:"DataType,omitempty"`

	// The maximum supported value for this parameter.
	MaximumValue float64 `json:"MaximumValue,omitempty"`

	// The minimum supported value for this parameter.
	MinimumValue float64 `json:"MinimumValue,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The data type of an object-based parameter.
	ObjectDataType string `json:"ObjectDataType,omitempty"`

	// An indication of whether the parameter is required to complete this action.
	Required bool `json:"Required,omitempty"`

}
