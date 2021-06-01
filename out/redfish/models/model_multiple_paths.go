/* -----------------------------------------------------------------
* multiple_paths.go -
*
* DMTF Redfish MultiplePaths resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MultiplePaths - Border Gateway Protocol (BGP) multiple path properties.
// Modeled after DMTF Redfish schema MultiplePaths
type MultiplePaths struct {
	// Maximum paths number.
	MaximumPaths int `json:"MaximumPaths,omitempty"`

	// Border Gateway Protocol (BGP) multiple paths status.
	UseMultiplePathsEnabled bool `json:"UseMultiplePathsEnabled,omitempty"`

}
