/* -----------------------------------------------------------------
* life_time.go -
*
* DMTF Redfish LifeTime resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The memory metrics for the lifetime of the memory.
type LifeTime struct {
	// The number of blocks read for the lifetime of the memory.
	BlocksRead int `json:"BlocksRead,omitempty"`

	// The number of blocks written for the lifetime of the memory.
	BlocksWritten int `json:"BlocksWritten,omitempty"`

	// The number of the correctable errors for the lifetime of the memory.
	CorrectableECCErrorCount int `json:"CorrectableECCErrorCount,omitempty"`

	// The number of the uncorrectable errors for the lifetime of the memory.
	UncorrectableECCErrorCount int `json:"UncorrectableECCErrorCount,omitempty"`

}
