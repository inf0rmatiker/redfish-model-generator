/* -----------------------------------------------------------------
* current_period.go -
*
* DMTF Redfish CurrentPeriod resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The memory metrics since the last system reset or ClearCurrentPeriod action.
type CurrentPeriod struct {
	// The number of blocks read since reset.
	BlocksRead int `json:"BlocksRead,omitempty"`

	// The number of blocks written since reset.
	BlocksWritten int `json:"BlocksWritten,omitempty"`

	// The number of the correctable errors since reset.
	CorrectableECCErrorCount int `json:"CorrectableECCErrorCount,omitempty"`

	// The number of the uncorrectable errors since reset.
	UncorrectableECCErrorCount int `json:"UncorrectableECCErrorCount,omitempty"`

}
