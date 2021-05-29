/* -----------------------------------------------------------------
* deep_operations.go -
*
* DMTF Redfish DeepOperations resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The information about deep operations that the service supports.
type DeepOperations struct {
	// An indication of whether the service supports the deep PATCH operation.
	DeepPATCH bool `json:"DeepPATCH,omitempty"`

	// An indication of whether the service supports the deep POST operation.
	DeepPOST bool `json:"DeepPOST,omitempty"`

	// The maximum levels of resources allowed in deep operations.
	MaxLevels int `json:"MaxLevels,omitempty"`

}
