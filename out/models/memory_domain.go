/* -----------------------------------------------------------------
* memory_domain.go -
*
* DMTF Redfish MemoryDomain resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The MemoryDomain schema describes a memory domain and its configuration.  Memory domains indicate to the client which memory, or DIMMs, can be grouped together in memory chunks to represent addressable memory.
type MemoryDomain struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// An indication of whether this memory domain supports the provisioning of blocks of memory.
	AllowsBlockProvisioning bool `json:"AllowsBlockProvisioning,omitempty"`

	// An indication of whether this memory domain supports the creation of memory chunks.
	AllowsMemoryChunkCreation bool `json:"AllowsMemoryChunkCreation,omitempty"`

	// An indication of whether this memory domain supports the creation of memory chunks with mirroring enabled.
	AllowsMirroring bool `json:"AllowsMirroring,omitempty"`

	// An indication of whether this memory domain supports the creation of memory chunks with sparing enabled.
	AllowsSparing bool `json:"AllowsSparing,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The interleave sets for the memory chunk.
	InterleavableMemorySets array `json:"InterleavableMemorySets,omitempty"`

	// The links to other Resources that are related to this Resource.
	Links Links `json:"Links,omitempty"`

	// The link to the collection of memory chunks associated with this memory domain.
	MemoryChunks MemoryChunksCollection `json:"MemoryChunks,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
