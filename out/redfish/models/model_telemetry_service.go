/* -----------------------------------------------------------------
* telemetry_service.go -
*
* DMTF Redfish TelemetryService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// TelemetryService - This is the schema definition for the Metrics Service.  It represents the properties for the service itself and has links to collections of metric definitions and metric report definitions.
// Modeled after DMTF Redfish schema TelemetryService
type TelemetryService struct {
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

	// This is a reference to a Log Service used by the Telemetry Service.
	LogService map[string]interface{} `json:"LogService,omitempty"`

	// The maximum number of metric reports supported by this service.
	MaxReports int `json:"MaxReports,omitempty"`

	// A link to the collection of Metric Definitions.
	MetricDefinitions map[string]interface{} `json:"MetricDefinitions,omitempty"`

	// A link to the collection of Metric Report Definitions.
	MetricReportDefinitions map[string]interface{} `json:"MetricReportDefinitions,omitempty"`

	// A link to the collection of Metric Reports.
	MetricReports map[string]interface{} `json:"MetricReports,omitempty"`

	// The minimum time interval between collections supported by this service.
	MinCollectionInterval string `json:"MinCollectionInterval,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

	// The functions that can be performed over each metric.
	SupportedCollectionFunctions []CollectionFunction `json:"SupportedCollectionFunctions,omitempty"`

	// A link to the collection of Triggers, which apply to metrics.
	Triggers map[string]interface{} `json:"Triggers,omitempty"`

}
