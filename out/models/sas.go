/* -----------------------------------------------------------------
* sas.go -
*
* DMTF Redfish SAS resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The physical metrics for Serial Attached SCSI (SAS).
type SAS struct {
	// The number of invalid dwords that have been received by the phy outside of phy reset sequences.
	InvalidDwordCount int `json:"InvalidDwordCount,omitempty"`

	// The number of times the phy has restarted the link reset sequence because it lost dword synchronization.
	LossOfDwordSynchronizationCount int `json:"LossOfDwordSynchronizationCount,omitempty"`

	// The number of dwords containing running disparity errors that have been received by the phy outside of phy reset sequences.
	RunningDisparityErrorCount int `json:"RunningDisparityErrorCount,omitempty"`

}
