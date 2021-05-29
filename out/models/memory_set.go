/* -----------------------------------------------------------------
* memory_set.go -
*
* DMTF Redfish MemorySet resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The interleave sets for a memory chunk.
type MemorySet struct {
	// The set of memory for a particular interleave set.
	MemorySet array `json:"MemorySet,omitempty"`

	MemorySet@odata.count count `json:"MemorySet@odata.count,omitempty"`

}
