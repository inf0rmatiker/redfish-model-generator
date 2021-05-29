/* -----------------------------------------------------------------
* memory_metrics.go -
*
* DMTF Redfish MemoryMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The usage and health statistics for a memory device or system memory summary.
type MemoryMetrics struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The memory bandwidth utilization as a percentage.
	BandwidthPercent float64 `json:"BandwidthPercent,omitempty"`

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

	// Operating speed of memory in MHz or MT/s as appropriate.
	OperatingSpeedMHz int `json:"OperatingSpeedMHz,omitempty"`

}
