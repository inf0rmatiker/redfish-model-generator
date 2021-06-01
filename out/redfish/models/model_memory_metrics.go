/* -----------------------------------------------------------------
* memory_metrics.go -
*
* DMTF Redfish MemoryMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MemoryMetrics - The usage and health statistics for a memory device or system memory summary.
// Modeled after DMTF Redfish schema MemoryMetrics
type MemoryMetrics struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The block size, in bytes.
	BlockSizeBytes int `json:"BlockSizeBytes,omitempty"`

	// The memory metrics since the last reset or ClearCurrentPeriod action.
	CurrentPeriod CurrentPeriod `json:"CurrentPeriod,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The health information of the memory.
	HealthData HealthData `json:"HealthData,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The memory metrics for the lifetime of the memory.
	LifeTime LifeTime `json:"LifeTime,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
