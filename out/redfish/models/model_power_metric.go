/* -----------------------------------------------------------------
* power_metric.go -
*
* DMTF Redfish PowerMetric resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

	// The time interval (or window) in which the PowerMetrics are measured over.
	IntervalInMin float64 `json:"IntervalInMin,omitempty"`

	// The lowest power consumption level over the measurement window (the last IntervalInMin minutes).
	MinConsumedWatts float64 `json:"MinConsumedWatts,omitempty"`

	// The highest power consumption level that has occured over the measurement window (the last IntervalInMin minutes).
	MaxConsumedWatts float64 `json:"MaxConsumedWatts,omitempty"`

	// The average power level over the measurement window (the last IntervalInMin minutes).
	AverageConsumedWatts float64 `json:"AverageConsumedWatts,omitempty"`

}
