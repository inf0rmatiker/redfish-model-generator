/* -----------------------------------------------------------------
* virtualization_offload.go -
*
* DMTF Redfish VirtualizationOffload resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// VirtualizationOffload - A Virtualization offload capability of a controller.
// Modeled after DMTF Redfish schema VirtualizationOffload
type VirtualizationOffload struct {
	// Single-root input/output virtualization (SR-IOV) capabilities.
	SRIOV SRIOV `json:"SRIOV,omitempty"`

	// The virtual function of the controller.
	VirtualFunction VirtualFunction `json:"VirtualFunction,omitempty"`

}
