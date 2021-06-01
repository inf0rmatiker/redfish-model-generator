/* -----------------------------------------------------------------
* memory_domain.go -
*
* DMTF Redfish MemoryDomain resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MemoryDomain - This is the schema definition of a Memory Domain and its configuration. Memory Domains are used to indicate to the client which Memory (DIMMs) can be grouped together in Memory Chunks to form interleave sets or otherwise grouped together.
// Modeled after DMTF Redfish schema MemoryDomain
type MemoryDomain struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// Indicates if this Memory Domain supports the provisioning of blocks of memory.
	AllowsBlockProvisioning bool `json:"AllowsBlockProvisioning,omitempty"`

	// Indicates if this Memory Domain supports the creation of Memory Chunks.
	AllowsMemoryChunkCreation bool `json:"AllowsMemoryChunkCreation,omitempty"`

	// Indicates if this Memory Domain supports the creation of Memory Chunks with mirroring enabled.
	AllowsMirroring bool `json:"AllowsMirroring,omitempty"`

	// Indicates if this Memory Domain supports the creation of Memory Chunks with sparing enabled.
	AllowsSparing bool `json:"AllowsSparing,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// This is the interleave sets for the memory chunk.
	InterleavableMemorySets []MemorySet `json:"InterleavableMemorySets,omitempty"`

	// A reference to the collection of Memory Chunks associated with this Memory Domain.
	MemoryChunks map[string]interface{} `json:"MemoryChunks,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
