/* -----------------------------------------------------------------
* virtualization_offload.go -
*
* DMTF Redfish VirtualizationOffload resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A Virtualization offload capability of a controller.
type VirtualizationOffload struct {
	// Single-root input/output virtualization (SR-IOV) capabilities.
	SRIOV SRIOV `json:"SRIOV,omitempty"`

	// The virtual function of the controller.
	VirtualFunction VirtualFunction `json:"VirtualFunction,omitempty"`

}
