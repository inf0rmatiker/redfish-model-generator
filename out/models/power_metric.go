/* -----------------------------------------------------------------
* power_metric.go -
*
* DMTF Redfish PowerMetric resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The power metrics for a resource.
type PowerMetric struct {
	// The average power level over the measurement window over the last IntervalInMin minutes.
	AverageConsumedWatts float64 `json:"AverageConsumedWatts,omitempty"`

	// The time interval, or window, over which the power metrics are measured.
	IntervalInMin int `json:"IntervalInMin,omitempty"`

	// The highest power consumption level, in watts, that has occurred over the measurement window within the last IntervalInMin minutes.
	MaxConsumedWatts float64 `json:"MaxConsumedWatts,omitempty"`

	// The lowest power consumption level, in watts, over the measurement window that occurred within the last IntervalInMin minutes.
	MinConsumedWatts float64 `json:"MinConsumedWatts,omitempty"`

}
