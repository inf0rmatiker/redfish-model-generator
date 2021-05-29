/* -----------------------------------------------------------------
* registry_entries.go -
*
* DMTF Redfish RegistryEntries resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The list of all attributes and metadata for this component.
type RegistryEntries struct {
	// An array of attributes and their possible values in the attribute registry.
	Attributes array `json:"Attributes,omitempty"`

	// An array of dependencies of attributes on this component.
	Dependencies array `json:"Dependencies,omitempty"`

	// An array for the attributes menus and their hierarchy in the attribute registry.
	Menus array `json:"Menus,omitempty"`

}
