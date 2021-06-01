/* -----------------------------------------------------------------
* enum_metric_type_enum.go -
*
* DMTF Redfish MetricTypeEnum enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type MetricTypeEnum string

const (
	// The trigger is for numeric sensor.
	MetricTypeEnum_NUMERIC MetricTypeEnum = "Numeric"

	// The trigger is for a discrete sensor.
	MetricTypeEnum_DISCRETE MetricTypeEnum = "Discrete"

)
