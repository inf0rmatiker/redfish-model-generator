/* -----------------------------------------------------------------
* controller_capabilities.go -
*
* DMTF Redfish ControllerCapabilities resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The capabilities of a controller.
type ControllerCapabilities struct {
	// Data center bridging (DCB) for this controller.
	DataCenterBridging DataCenterBridging `json:"DataCenterBridging,omitempty"`

	// NIC Partitioning (NPAR) capabilities for this controller.
	NPAR NicPartitioning `json:"NPAR,omitempty"`

	// N_Port ID Virtualization (NPIV) capabilities for this controller.
	NPIV NPIV `json:"NPIV,omitempty"`

	// The maximum number of physical functions available on this controller.
	NetworkDeviceFunctionCount int `json:"NetworkDeviceFunctionCount,omitempty"`

	// The number of physical ports on this controller.
	NetworkPortCount int `json:"NetworkPortCount,omitempty"`

	// Virtualization offload for this controller.
	VirtualizationOffload VirtualizationOffload `json:"VirtualizationOffload,omitempty"`

}
