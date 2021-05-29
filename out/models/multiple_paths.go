/* -----------------------------------------------------------------
* multiple_paths.go -
*
* DMTF Redfish MultiplePaths resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Border Gateway Protocol (BGP) multiple path properties.
type MultiplePaths struct {
	// Maximum paths number.
	MaximumPaths int `json:"MaximumPaths,omitempty"`

	// Border Gateway Protocol (BGP) multiple paths status.
	UseMultiplePathsEnabled bool `json:"UseMultiplePathsEnabled,omitempty"`

}
