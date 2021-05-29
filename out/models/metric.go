/* -----------------------------------------------------------------
* metric.go -
*
* DMTF Redfish Metric resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Specifies a metric to include in the metric report.  The metrics are derived by applying a calculation on each of the listed metric properties.
type Metric struct {
	// The duration over which the function is computed.
	CollectionDuration string `json:"CollectionDuration,omitempty"`

	// Specifies the function to perform on each of the metric properties listed in the MetricProperties property.
	CollectionFunction CalculationAlgorithmEnum `json:"CollectionFunction,omitempty"`

	// The scope of time over which the function is applied.
	CollectionTimeScope CollectionTimeScope `json:"CollectionTimeScope,omitempty"`

	// The label for the metric definition that is derived by applying the CollectionFunction to the metric property.  It matches the Id property of the corresponding metric definition.
	MetricId string `json:"MetricId,omitempty"`

	// The set of URIs for the properties on which this metric is collected.
	MetricProperties []string `json:"MetricProperties,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
