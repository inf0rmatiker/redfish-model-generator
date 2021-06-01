/* -----------------------------------------------------------------
* dependency.go -
*
* DMTF Redfish Dependency resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Dependency - The dependency expression for one or more Attributes in this Attribute Registry.
// Modeled after DMTF Redfish schema Dependency
type Dependency struct {
	// Array of the map-from conditions for mapping dependency.
	MapFrom []MapFrom `json:"MapFrom,omitempty"`

	// The Name of the attribute that is affected by this dependency expression.
	MapToAttribute string `json:"MapToAttribute,omitempty"`

	// The meta-data property of the attribute specified in MapFromAttribute that is used to evaluate this dependency expression.
	MapToProperty MapToProperty `json:"MapToProperty,omitempty"`

	// The value that MapToProperty is changed to if the dependency expression evaluates to true.
	MapToValue float64 `json:"MapToValue,omitempty"`

}
