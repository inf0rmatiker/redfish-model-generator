/* -----------------------------------------------------------------
* power_allocation.go -
*
* DMTF Redfish PowerAllocation resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PowerAllocation - Power allocation for a subsystem.
// Modeled after DMTF Redfish schema PowerAllocation
type PowerAllocation struct {
	// The total amount of power that has been allocated or budgeted to this subsystem.
	AllocatedWatts float64 `json:"AllocatedWatts,omitempty"`

	// The potential power, in watts, that the subsystem requests, which might be higher than the current level being consumed because the requested power includes a budget that the subsystem wants for future use.
	RequestedWatts float64 `json:"RequestedWatts,omitempty"`

}
