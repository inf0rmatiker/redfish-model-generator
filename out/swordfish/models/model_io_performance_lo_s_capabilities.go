/* -----------------------------------------------------------------
* io_performance_lo_s_capabilities.go -
*
* SNIA Swordfish IOPerformanceLoSCapabilities resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// IOPerformanceLoSCapabilities - Describe IO performance capabilities.
// Modeled after SNIA Swordfish schema IOPerformanceLoSCapabilities
type IOPerformanceLoSCapabilities struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Limiting IOPS is supported.
	IOLimitingIsSupported bool `json:"IOLimitingIsSupported,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The value identifies this resource.
	Identifier map[string]interface{} `json:"Identifier,omitempty"`

	// Maximum sampling period over which average values are calculated.
	MaxSamplePeriod string `json:"MaxSamplePeriod,omitempty"`

	// Minimum sampling period over which average values are calculated.
	MinSamplePeriod string `json:"MinSamplePeriod,omitempty"`

	// Minimum supported average IO latency.
	MinSupportedIoOperationLatencyMicroseconds int `json:"MinSupportedIoOperationLatencyMicroseconds,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// A collection of supported workloads.
	SupportedIOWorkloads []IOWorkload `json:"SupportedIOWorkloads,omitempty"`

	// Collection of known and supported IOPerformanceLinesOfService.
	SupportedLinesOfService []IOPerformanceLineOfService `json:"SupportedLinesOfService,omitempty"`

	SupportedLinesOfServiceOdataCount int `json:"SupportedLinesOfService@odata.count,omitempty"`

}
