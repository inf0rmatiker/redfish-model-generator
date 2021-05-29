/* -----------------------------------------------------------------
* mapping.go -
*
* DMTF Redfish Mapping resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The mapping between a Resource type and the relevant privileges that accesses the Resource.
type Mapping struct {
	// The Resource name, such as `Manager`.
	Entity string `json:"Entity,omitempty"`

	// List mapping between HTTP methods and privilege required for the Resource.
	OperationMap OperationMap `json:"OperationMap,omitempty"`

	// The privilege overrides of properties within a Resource.
	PropertyOverrides array `json:"PropertyOverrides,omitempty"`

	// The privilege overrides of Resource URIs.
	ResourceURIOverrides array `json:"ResourceURIOverrides,omitempty"`

	// The privilege overrides of the subordinate Resource.
	SubordinateOverrides array `json:"SubordinateOverrides,omitempty"`

}
