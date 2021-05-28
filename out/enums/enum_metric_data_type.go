/* -----------------------------------------------------------------
 * enum_metric_data_type.go -
 *
 * DMTF Redfish MetricDataType enum definitions
 *
 * © Copyright 2021 Hewlett Packard Enterprise Development LP
 *
 * ----------------------------------------------------------------- */

package openapi

type MetricDataType string

const (
	// The JSON boolean definition.
	MetricDataType_BOOLEAN MetricDataType = "Boolean"

	// The JSON string definition with the date-time format.
	MetricDataType_DATE_TIME MetricDataType = "DateTime"

	// The JSON decimal definition.
	MetricDataType_DECIMAL MetricDataType = "Decimal"

	// The JSON integer definition.
	MetricDataType_INTEGER MetricDataType = "Integer"

	// The JSON string definition.
	MetricDataType_STRING MetricDataType = "String"

	// The JSON string definition with a set of defined enumerations.
	MetricDataType_ENUMERATION MetricDataType = "Enumeration"

)
