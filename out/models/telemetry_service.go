/* -----------------------------------------------------------------
* telemetry_service.go -
*
* DMTF Redfish TelemetryService resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The TelemetryService schema describes a telemetry service.  The telemetry service is used to for collecting and reporting metric data within the Redfish Service.
type TelemetryService struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The link to a log service that the telemetry service uses.  This service can be a dedicated log service or a pointer a log service under another resource, such as a manager.
	LogService LogService `json:"LogService,omitempty"`

	// The maximum number of metric reports that this service supports.
	MaxReports int `json:"MaxReports,omitempty"`

	// The link to the collection of metric definitions.
	MetricDefinitions MetricDefinitionCollection `json:"MetricDefinitions,omitempty"`

	// The link to the collection of metric report definitions.
	MetricReportDefinitions MetricReportDefinitionCollection `json:"MetricReportDefinitions,omitempty"`

	// The link to the collection of metric reports.
	MetricReports MetricReportCollection `json:"MetricReports,omitempty"`

	// The minimum time interval between gathering metric data that this service allows.
	MinCollectionInterval string `json:"MinCollectionInterval,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An indication of whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The functions that can be performed over each metric.
	SupportedCollectionFunctions array `json:"SupportedCollectionFunctions,omitempty"`

	// The link to the collection of triggers that apply to metrics.
	Triggers TriggersCollection `json:"Triggers,omitempty"`

}
