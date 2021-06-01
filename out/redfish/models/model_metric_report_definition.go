/* -----------------------------------------------------------------
* metric_report_definition.go -
*
* DMTF Redfish MetricReportDefinition resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MetricReportDefinition - The MetricReportDefinition schema describes set of metrics that are collected into a metric report.
// Modeled after DMTF Redfish schema MetricReportDefinition
type MetricReportDefinition struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// The maximum number of entries that can be appended to a metric report.  When the metric report reaches its limit, its behavior is dictated by the ReportUpdates property.
	AppendLimit int `json:"AppendLimit,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The list of URIs with wildcards and property identifiers to include in the metric report.  If a URI has wildcards, the wildcards are substituted as specified in the Wildcards property.
	MetricProperties []string `json:"MetricProperties,omitempty"`

	// The location where the resultant metric report is placed.
	MetricReport map[string]interface{} `json:"MetricReport,omitempty"`

	// An indication of whether the generation of new metric reports is enabled.
	MetricReportDefinitionEnabled bool `json:"MetricReportDefinitionEnabled,omitempty"`

	// Specifies when the metric report is generated.
	MetricReportDefinitionType MetricReportDefinitionType `json:"MetricReportDefinitionType,omitempty"`

	// The interval at which to send the complete metric report because the Redfish client wants refreshed metric data even when the data has not changed.  This property value is always greater than the recurrence interval of a metric report, and it only applies when the SuppressRepeatedMetricValue property is `true`.
	MetricReportHeartbeatInterval string `json:"MetricReportHeartbeatInterval,omitempty"`

	// The list of metrics to include in the metric report.  The metrics may include metric properties or calculations applied to a metric property.
	Metrics []Metric `json:"Metrics,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The set of actions to perform when a metric report is generated.
	ReportActions []ReportActionsEnum `json:"ReportActions,omitempty"`

	// The behavior for how subsequent metric reports are handled in relationship to an existing metric report created from the metric report definition.  Namely, whether to overwrite, append, or create a report resource.
	ReportUpdates ReportUpdatesEnum `json:"ReportUpdates,omitempty"`

	// The schedule for generating the metric report.
	Schedule map[string]interface{} `json:"Schedule,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status map[string]interface{} `json:"Status,omitempty"`

	// An indication of whether any metrics are suppressed from the generated metric report.  If `true`, any metric that equals the same value in the previously generated metric report is suppressed from the current report.  Also, duplicate metrics are suppressed.  If `false`, no metrics are suppressed from the current report.  The current report may contain no metrics if all metrics equal the values in the previously generated metric report.
	SuppressRepeatedMetricValue bool `json:"SuppressRepeatedMetricValue,omitempty"`

	// The set of wildcards and their substitution values for the entries in the MetricProperties property.
	Wildcards []Wildcard `json:"Wildcards,omitempty"`

}
