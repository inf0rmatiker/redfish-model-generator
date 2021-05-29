/* -----------------------------------------------------------------
* memory_chunks.go -
*
* DMTF Redfish MemoryChunks resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The schema definition of a memory chunk and its configuration.
type MemoryChunks struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// Offset of the memory chunk in the address range in MiB.
	AddressRangeOffsetMiB int `json:"AddressRangeOffsetMiB,omitempty"`

	// Memory type of this memory chunk.
	AddressRangeType AddressRangeType `json:"AddressRangeType,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// A user-configurable string to name the memory chunk.
	DisplayName string `json:"DisplayName,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The interleave sets for the memory chunk.
	InterleaveSets array `json:"InterleaveSets,omitempty"`

	// An indication of whether memory mirroring is enabled for this memory chunk.
	IsMirrorEnabled bool `json:"IsMirrorEnabled,omitempty"`

	// An indication of whether sparing is enabled for this memory chunk.
	IsSpare bool `json:"IsSpare,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// Size of the memory chunk measured in mebibytes (MiB).
	MemoryChunkSizeMiB int `json:"MemoryChunkSizeMiB,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

}
