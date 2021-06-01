/* -----------------------------------------------------------------
* processor_metrics.go -
*
* DMTF Redfish ProcessorMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ProcessorMetrics - ProcessorMetrics contains usage and health statistics for a Processor (all Cores).
// Modeled after DMTF Redfish schema ProcessorMetrics
type ProcessorMetrics struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The average frequency of the processor.
	AverageFrequencyMHz float64 `json:"AverageFrequencyMHz,omitempty"`

	// The CPU bandwidth as a percentage.
	BandwidthPercent float64 `json:"BandwidthPercent,omitempty"`

	// The processor cache metrics.
	Cache []CacheMetrics `json:"Cache,omitempty"`

	// The power consumed by the processor.
	ConsumedPowerWatt float64 `json:"ConsumedPowerWatt,omitempty"`

	// The processor core metrics.
	CoreMetrics []CoreMetrics `json:"CoreMetrics,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The frequency relative to the nominal processor frequency ratio.
	FrequencyRatio float64 `json:"FrequencyRatio,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The percentage of time spent in kernel mode.
	KernelPercent float64 `json:"KernelPercent,omitempty"`

	// The local memory bandwidth usage in bytes.
	LocalMemoryBandwidthBytes int `json:"LocalMemoryBandwidthBytes,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The remote memory bandwidth usage in bytes.
	RemoteMemoryBandwidthBytes int `json:"RemoteMemoryBandwidthBytes,omitempty"`

	// The temperature of the processor.
	TemperatureCelsius float64 `json:"TemperatureCelsius,omitempty"`

	// The CPU margin to throttle (temperature offset in degree Celsius).
	ThrottlingCelsius float64 `json:"ThrottlingCelsius,omitempty"`

	// The percentage of time spent in user mode.
	UserPercent float64 `json:"UserPercent,omitempty"`

}
