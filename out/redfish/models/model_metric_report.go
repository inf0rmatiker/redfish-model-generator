/* -----------------------------------------------------------------
* metric_report.go -
*
* DMTF Redfish MetricReport resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MetricReport - The metric definitions that create a metric report.
// Modeled after DMTF Redfish schema MetricReport
type MetricReport struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this Resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The definitions in the metric report.
	MetricReportDefinition map[string]interface{} `json:"MetricReportDefinition,omitempty"`

	// An array of metric values for the metered items of this Metric.
	MetricValues []MetricValue `json:"MetricValues,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The current sequence identifier for this metric report.
	ReportSequence string `json:"ReportSequence"`

}
