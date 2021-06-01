/* -----------------------------------------------------------------
* capacity_source.go -
*
* SNIA Swordfish CapacitySource resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// CapacitySource - A description of the type and source of storage.
// Modeled after SNIA Swordfish schema CapacitySource
type CapacitySource struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The amount of space that has been provided from the ProvidingDrives, ProvidingVolumes, ProvidingMemory or ProvidingPools.
	ProvidedCapacity Capacity `json:"ProvidedCapacity,omitempty"`

	// The ClassOfService provided from the ProvidingDrives, ProvidingVolumes, ProvidingMemoryChunks, ProvidingMemory or ProvidingPools.
	ProvidedClassOfService map[string]interface{} `json:"ProvidedClassOfService,omitempty"`

	// The drive or drives that provide this space.
	ProvidingDrives map[string]interface{} `json:"ProvidingDrives,omitempty"`

	// The memory that provides this space.
	ProvidingMemory map[string]interface{} `json:"ProvidingMemory,omitempty"`

	// The memory chunks that provide this space.
	ProvidingMemoryChunks map[string]interface{} `json:"ProvidingMemoryChunks,omitempty"`

	// The pool or pools that provide this space.
	ProvidingPools map[string]interface{} `json:"ProvidingPools,omitempty"`

	// The volume or volumes that provide this space.
	ProvidingVolumes map[string]interface{} `json:"ProvidingVolumes,omitempty"`

}
