/* -----------------------------------------------------------------
* metric_report.go -
*
* DMTF Redfish MetricReport resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The MetricReport schema represents a set of collected metrics.
type MetricReport struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// A context can be supplied at subscription time.  This property is the context value supplied by the subscriber.
	Context string `json:"Context,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The link to the definition of this metric report.
	MetricReportDefinition MetricReportDefinition `json:"MetricReportDefinition,omitempty"`

	// An array of metric values for the metered items of this metric report.
	MetricValues array `json:"MetricValues,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The current sequence identifier for this metric report.
	ReportSequence string `json:"ReportSequence,omitempty"`

	// The time associated with the metric report in its entirety.  The time of the metric report can be relevant when the time of individual metrics are minimally different.
	Timestamp string `json:"Timestamp,omitempty"`

}
