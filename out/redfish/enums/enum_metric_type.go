/* -----------------------------------------------------------------
* enum_metric_type.go -
*
* DMTF Redfish MetricType enum definitions
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

type MetricType string

const (
	// The metric is a numeric metric.  The metric value is any real number.
	MetricType_NUMERIC MetricType = "Numeric"

	// The metric is a discrete metric.  The metric value is discrete.  The possible values are listed in the DiscreteValues property.
	MetricType_DISCRETE MetricType = "Discrete"

	// The metric is a gauge metric.  The metric value is a real number.  When the metric value reaches the gauges extrema, it stays at that value, until the reading falls within the extrema.
	MetricType_GAUGE MetricType = "Gauge"

	// The metric is a counter metric.  The metric reading is a non-negative integer that increases monotonically.  When a counter reaches its maximum, the value resets to 0 and resumes counting.
	MetricType_COUNTER MetricType = "Counter"

	// The metric is a countdown metric.  The metric reading is a non-negative integer that decreases monotonically.  When a counter reaches its minimum, the value resets to preset value and resumes counting down.
	MetricType_COUNTDOWN MetricType = "Countdown"

)
