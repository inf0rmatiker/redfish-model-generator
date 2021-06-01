/* -----------------------------------------------------------------
* cache_metrics.go -
*
* DMTF Redfish CacheMetrics resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// CacheMetrics - The processor core metrics.
// Modeled after DMTF Redfish schema CacheMetrics
type CacheMetrics struct {
	// The number of cache line misses in millions.
	CacheMiss float64 `json:"CacheMiss,omitempty"`

	// The number of cache misses per instruction.
	CacheMissesPerInstruction float64 `json:"CacheMissesPerInstruction,omitempty"`

	// The cache line hit ratio.
	HitRatio float64 `json:"HitRatio,omitempty"`

	// The cache level.
	Level string `json:"Level,omitempty"`

	// The total cache level occupancy in bytes.
	OccupancyBytes int `json:"OccupancyBytes,omitempty"`

	// The total cache occupancy percentage.
	OccupancyPercent float64 `json:"OccupancyPercent,omitempty"`

}
