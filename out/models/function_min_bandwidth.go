/* -----------------------------------------------------------------
* function_min_bandwidth.go -
*
* DMTF Redfish FunctionMinBandwidth resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A minimum bandwidth allocation percentage for a network device functions associated a port.
type FunctionMinBandwidth struct {
	// The minimum bandwidth allocation percentage allocated to the corresponding network device function instance.
	AllocationPercent int `json:"AllocationPercent,omitempty"`

	// The link to the network device function associated with this bandwidth setting of this network port.
	NetworkDeviceFunction NetworkDeviceFunction `json:"NetworkDeviceFunction,omitempty"`

}
