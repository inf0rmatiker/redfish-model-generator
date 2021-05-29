/* -----------------------------------------------------------------
* cache_metrics_total.go -
*
* DMTF Redfish CacheMetricsTotal resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The total cache metrics for a processor.
type CacheMetricsTotal struct {
	// The cache metrics since the last reset for this processor.
	CurrentPeriod CurrentPeriod `json:"CurrentPeriod,omitempty"`

	// The cache metrics for the lifetime of this processor.
	LifeTime LifeTime `json:"LifeTime,omitempty"`

}
