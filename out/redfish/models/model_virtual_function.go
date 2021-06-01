/* -----------------------------------------------------------------
* virtual_function.go -
*
* DMTF Redfish VirtualFunction resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// VirtualFunction - A virtual function of a controller.
// Modeled after DMTF Redfish schema VirtualFunction
type VirtualFunction struct {
	// The maximum number of virtual functions supported by this controller.
	DeviceMaxCount int `json:"DeviceMaxCount,omitempty"`

	// The minimum number of virtual functions that can be allocated or moved between physical functions for this controller.
	MinAssignmentGroupSize int `json:"MinAssignmentGroupSize,omitempty"`

	// The maximum number of virtual functions supported per network port for this controller.
	NetworkPortMaxCount int `json:"NetworkPortMaxCount,omitempty"`

}
