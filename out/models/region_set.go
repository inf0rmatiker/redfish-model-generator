/* -----------------------------------------------------------------
* region_set.go -
*
* DMTF Redfish RegionSet resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Memory region information within a memory device.
type RegionSet struct {
	// The classification of memory that the memory region occupies.
	MemoryClassification MemoryClassification `json:"MemoryClassification,omitempty"`

	// Offset within the memory that corresponds to the start of this memory region in mebibytes (MiB).
	OffsetMiB int `json:"OffsetMiB,omitempty"`

	// An indication of whether the passphrase is enabled for this region.
	PassphraseEnabled bool `json:"PassphraseEnabled,omitempty"`

	// An indication of whether the state of the passphrase for this region is enabled.
	PassphraseState bool `json:"PassphraseState,omitempty"`

	// Unique region ID representing a specific region within the memory device.
	RegionId string `json:"RegionId,omitempty"`

	// Size of this memory region in mebibytes (MiB).
	SizeMiB int `json:"SizeMiB,omitempty"`

}
