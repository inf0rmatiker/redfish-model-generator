/* -----------------------------------------------------------------
* mapping.go -
*
* DMTF Redfish Mapping resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Mapping - This type describes a mapping between an entity and the relevant privileges used to access it.
// Modeled after DMTF Redfish schema Mapping
type Mapping struct {
	// Indicates entity name. e.g., Manager.
	Entity string `json:"Entity,omitempty"`

	// List mapping between HTTP method and privilege required for entity.
	OperationMap OperationMap `json:"OperationMap,omitempty"`

	// Indicates privilege overrides of property or element within a entity.
	PropertyOverrides []Target_PrivilegeMap `json:"PropertyOverrides,omitempty"`

	// Indicates privilege overrides of Resource URI.
	ResourceURIOverrides []Target_PrivilegeMap `json:"ResourceURIOverrides,omitempty"`

	// Indicates privilege overrides of subordinate resource.
	SubordinateOverrides []Target_PrivilegeMap `json:"SubordinateOverrides,omitempty"`

}
