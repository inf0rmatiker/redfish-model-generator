/* -----------------------------------------------------------------
* metric_value.go -
*
* DMTF Redfish MetricValue resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MetricValue - A metric Value.
// Modeled after DMTF Redfish schema MetricValue
type MetricValue struct {
	// The link to the metric.
	MetricDefinition map[string]interface{} `json:"MetricDefinition,omitempty"`

	// The metric definitions identifier for this metric.
	MetricId string `json:"MetricId,omitempty"`

	// The URI for the property from which this metric is derived.
	MetricProperty string `json:"MetricProperty,omitempty"`

	// The metric value, as a string.
	MetricValue string `json:"MetricValue,omitempty"`

	// The time when the metric is obtained.  A management application may establish a time series of metric data by retrieving the instances of metric value and sorting them according to their Timestamp.
	Timestamp string `json:"Timestamp,omitempty"`

}
