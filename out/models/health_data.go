/* -----------------------------------------------------------------
* health_data.go -
*
* DMTF Redfish HealthData resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The health information of the memory.
type HealthData struct {
	// Alarm trip information about the memory.
	AlarmTrips AlarmTrips `json:"AlarmTrips,omitempty"`

	// An indication of whether data loss was detected.
	DataLossDetected bool `json:"DataLossDetected,omitempty"`

	// An indication of whether the last shutdown succeeded.
	LastShutdownSuccess bool `json:"LastShutdownSuccess,omitempty"`

	// An indication of whether performance has degraded.
	PerformanceDegraded bool `json:"PerformanceDegraded,omitempty"`

	// The percentage of reads and writes that are predicted to still be available for the media.
	PredictedMediaLifeLeftPercent float64 `json:"PredictedMediaLifeLeftPercent,omitempty"`

	// The remaining spare blocks, as a percentage.
	RemainingSpareBlockPercentage float64 `json:"RemainingSpareBlockPercentage,omitempty"`

}
