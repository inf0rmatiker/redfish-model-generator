/* -----------------------------------------------------------------
* storage_pool.go -
*
* SNIA Swordfish StoragePool resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// StoragePool - A container of data storage.
// Modeled after SNIA Swordfish schema StoragePool
type StoragePool struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// A reference to the collection of storage pools allocated from this storage pool.
	AllocatedPools map[string]interface{} `json:"AllocatedPools,omitempty"`

	// A reference to the collection of volumes allocated from this storage pool.
	AllocatedVolumes map[string]interface{} `json:"AllocatedVolumes,omitempty"`

	// Maximum Block size in bytes.
	BlockSizeBytes int `json:"BlockSizeBytes,omitempty"`

	// Capacity utilization.
	Capacity map[string]interface{} `json:"Capacity,omitempty"`

	// An array of space allocations to this store.
	CapacitySources []CapacitySource `json:"CapacitySources,omitempty"`

	CapacitySourcesOdataCount int `json:"CapacitySources@odata.count,omitempty"`

	// The ClassesOfService supported by this storage pool.
	ClassesOfService map[string]interface{} `json:"ClassesOfService,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The value identifies this resource.
	Identifier map[string]interface{} `json:"Identifier,omitempty"`

	// The links object contains the links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// Low space warning threshold specified in percents.
	LowSpaceWarningThresholdPercents []int `json:"LowSpaceWarningThresholdPercents,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The property contains the status of the StoragePool.
	Status map[string]interface{} `json:"Status,omitempty"`

}
