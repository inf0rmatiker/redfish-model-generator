/* -----------------------------------------------------------------
* io_performance_line_of_service.go -
*
* SNIA Swordfish IOPerformanceLineOfService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// IOPerformanceLineOfService - Describe service option within the IO performance line of service.
// Modeled after SNIA Swordfish schema IOPerformanceLineOfService
type IOPerformanceLineOfService struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// Expected average IO latency.
	AverageIOOperationLatencyMicroseconds int `json:"AverageIOOperationLatencyMicroseconds,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Limit the IOPS.
	IOOperationsPerSecondIsLimited bool `json:"IOOperationsPerSecondIsLimited,omitempty"`

	// A description of the expected workload.
	IOWorkload map[string]interface{} `json:"IOWorkload,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The amount of IOPS a volume of a given committed size can support.
	MaxIOOperationsPerSecondPerTerabyte int `json:"MaxIOOperationsPerSecondPerTerabyte,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Sampling period over which average values are calculated.
	SamplePeriod string `json:"SamplePeriod,omitempty"`

}
