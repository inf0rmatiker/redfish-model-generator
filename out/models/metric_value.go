/* -----------------------------------------------------------------
* metric_value.go -
*
* DMTF Redfish MetricValue resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Properties that capture a metric value and other associated information.
type MetricValue struct {
	// The link to the metric definition for this metric.
	MetricDefinition MetricDefinition `json:"MetricDefinition,omitempty"`

	// The metric definitions identifier for this metric.
	MetricId string `json:"MetricId,omitempty"`

	// The URI for the property from which this metric is derived.
	MetricProperty string `json:"MetricProperty,omitempty"`

	// The metric value, as a string.
	MetricValue string `json:"MetricValue,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The date and time when the metric is obtained.  A management application can establish a time series of metric data by retrieving the instances of metric value and sorting them according to their timestamp.
	Timestamp string `json:"Timestamp,omitempty"`

}
