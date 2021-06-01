/* -----------------------------------------------------------------
* memory_set.go -
*
* DMTF Redfish MemorySet resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MemorySet - This is the interleave sets for a memory chunk.
// Modeled after DMTF Redfish schema MemorySet
type MemorySet struct {
	// This is the collection of memory for a particular interleave set.
	MemorySet []Memory `json:"MemorySet,omitempty"`

	MemorySetOdataCount int `json:"MemorySet@odata.count,omitempty"`

}
