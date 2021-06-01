/* -----------------------------------------------------------------
* metric_definition.go -
*
* DMTF Redfish MetricDefinition resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// MetricDefinition - The MetricDefinition schema describes the metadata information for a metric.
// Modeled after DMTF Redfish schema MetricDefinition
type MetricDefinition struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The estimated percent error of measured versus actual values.
	Accuracy float64 `json:"Accuracy,omitempty"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// An indication of whether the metric can be used in a calculation.
	Calculable Calculable `json:"Calculable,omitempty"`

	// The calculation that is performed on a source metric to obtain the metric being defined.
	CalculationAlgorithm CalculationAlgorithmEnum `json:"CalculationAlgorithm,omitempty"`

	// The metric properties that are part of a calculation.
	CalculationParameters []CalculationParamsType `json:"CalculationParameters,omitempty"`

	// The time interval over which the metric calculation is performed.
	CalculationTimeInterval string `json:"CalculationTimeInterval,omitempty"`

	// The calibration offset added to the metric reading.
	Calibration float64 `json:"Calibration,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// This array property specifies possible values of a discrete metric.
	DiscreteValues []string `json:"DiscreteValues,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The implementation of the metric.
	Implementation ImplementationType `json:"Implementation,omitempty"`

	// An indication of whether the metric values are linear versus non-linear.
	IsLinear bool `json:"IsLinear,omitempty"`

	// Maximum value for metric reading.
	MaxReadingRange float64 `json:"MaxReadingRange,omitempty"`

	// The data type of the metric.
	MetricDataType MetricDataType `json:"MetricDataType,omitempty"`

	// The list of URIs with wildcards and property identifiers that this metric definition defines.  If a URI has wildcards, the wildcards are substituted as specified in the Wildcards array property.
	MetricProperties []string `json:"MetricProperties,omitempty"`

	// The type of metric.
	MetricType MetricType `json:"MetricType,omitempty"`

	// Minimum value for metric reading.
	MinReadingRange float64 `json:"MinReadingRange,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM-defined calculation that is performed on a source metric to obtain the metric being defined.
	OEMCalculationAlgorithm string `json:"OEMCalculationAlgorithm,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The physical context of the metric.
	PhysicalContext map[string]interface{} `json:"PhysicalContext,omitempty"`

	// Number of significant digits in the metric reading.
	Precision int `json:"Precision,omitempty"`

	// The time interval between when a metric is updated.
	SensingInterval string `json:"SensingInterval,omitempty"`

	// The accuracy of the timestamp.
	TimestampAccuracy string `json:"TimestampAccuracy,omitempty"`

	// The units of measure for this metric.
	Units string `json:"Units,omitempty"`

	// The wildcards and their substitution values for the entries in the MetricProperties array property.
	Wildcards []Wildcard `json:"Wildcards,omitempty"`

}
