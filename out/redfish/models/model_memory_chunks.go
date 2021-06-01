/* -----------------------------------------------------------------
* memory_chunks.go -
*
* DMTF Redfish MemoryChunks resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MemoryChunks - The schema definition of a memory chunk and its configuration.
// Modeled after DMTF Redfish schema MemoryChunks
type MemoryChunks struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// Memory type of this memory chunk.
	AddressRangeType AddressRangeType `json:"AddressRangeType,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The interleave sets for the memory chunk.
	InterleaveSets []InterleaveSet `json:"InterleaveSets,omitempty"`

	// An indication of whether memory mirroring is enabled for this memory chunk.
	IsMirrorEnabled bool `json:"IsMirrorEnabled,omitempty"`

	// An indication of whether sparing is enabled for this memory chunk.
	IsSpare bool `json:"IsSpare,omitempty"`

	// Size of the memory chunk measured in mebibytes (MiB).
	MemoryChunkSizeMiB int `json:"MemoryChunkSizeMiB,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the Resource and its subordinate or dependent Resources.
	Status map[string]interface{} `json:"Status,omitempty"`

}
