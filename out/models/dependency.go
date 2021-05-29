/* -----------------------------------------------------------------
* dependency.go -
*
* DMTF Redfish Dependency resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The dependency expression for one or more attributes in this attribute registry.
type Dependency struct {
	// An array of the map-from conditions for a mapping dependency.
	MapFrom array `json:"MapFrom,omitempty"`

	// The AttributeName of the attribute that is affected by this dependency expression.
	MapToAttribute string `json:"MapToAttribute,omitempty"`

	// The metadata property for the attribute that contains the map-from condition that evaluates this dependency expression.
	MapToProperty MapToProperty `json:"MapToProperty,omitempty"`

	// The value that the map-to property changes to if the dependency expression evaluates to `true`.
	MapToValue string `json:"MapToValue,omitempty"`

}
