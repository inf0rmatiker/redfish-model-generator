/* -----------------------------------------------------------------
* interleave_set.go -
*
* DMTF Redfish InterleaveSet resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// This an interleave set for a memory chunk.
type InterleaveSet struct {
	// Describes a memory device of the interleave set.
	Memory idRef `json:"Memory,omitempty"`

	// Level of the interleave set for multi-level tiered memory.
	MemoryLevel int `json:"MemoryLevel,omitempty"`

	// Offset within the DIMM that corresponds to the start of this memory region, measured in mebibytes (MiB).
	OffsetMiB int `json:"OffsetMiB,omitempty"`

	// DIMM region identifier.
	RegionId string `json:"RegionId,omitempty"`

	// Size of this memory region measured in mebibytes (MiB).
	SizeMiB int `json:"SizeMiB,omitempty"`

}
