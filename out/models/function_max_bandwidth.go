/* -----------------------------------------------------------------
* function_max_bandwidth.go -
*
* DMTF Redfish FunctionMaxBandwidth resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A maximum bandwidth allocation percentage for a network device functions associated a port.
type FunctionMaxBandwidth struct {
	// The maximum bandwidth allocation percentage allocated to the corresponding network device function instance.
	AllocationPercent int `json:"AllocationPercent,omitempty"`

	// The link to the network device function associated with this bandwidth setting of this network port.
	NetworkDeviceFunction NetworkDeviceFunction `json:"NetworkDeviceFunction,omitempty"`

}
