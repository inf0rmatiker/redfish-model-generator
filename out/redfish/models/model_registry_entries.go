/* -----------------------------------------------------------------
* registry_entries.go -
*
* DMTF Redfish RegistryEntries resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// RegistryEntries - List of all attributes and their metadata for this component.
// Modeled after DMTF Redfish schema RegistryEntries
type RegistryEntries struct {
	// The array containing the attributes and their possible values.
	Attributes []map[string]interface{} `json:"Attributes,omitempty"`

	// The array containing a list of dependencies of attributes on this component.
	Dependencies []Dependencies `json:"Dependencies,omitempty"`

	// The array containing the attributes menus and their hierarchy.
	Menus []Menus `json:"Menus,omitempty"`

}
