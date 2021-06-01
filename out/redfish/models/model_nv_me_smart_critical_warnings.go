/* -----------------------------------------------------------------
* nv_me_smart_critical_warnings.go -
*
* DMTF Redfish NVMeSMARTCriticalWarnings resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// NVMeSMARTCriticalWarnings - The NVMe SMART Critical Warnings for a storage controller.
// Modeled after DMTF Redfish schema NVMeSMARTCriticalWarnings
type NVMeSMARTCriticalWarnings struct {
	// Indicates the media has been placed in read only mode.
	MediaInReadOnly bool `json:"MediaInReadOnly,omitempty"`

	// Indicates that the NVM subsystem reliability has been compromised.
	OverallSubsystemDegraded bool `json:"OverallSubsystemDegraded,omitempty"`

	// The Persistent Memory Region has become unreliable.
	PMRUnreliable bool `json:"PMRUnreliable,omitempty"`

	// Indicates that the volatile memory backup device has failed.
	PowerBackupFailed bool `json:"PowerBackupFailed,omitempty"`

	// Indicates that the available spare capacity has fallen below the threshold.
	SpareCapacityWornOut bool `json:"SpareCapacityWornOut,omitempty"`

}
