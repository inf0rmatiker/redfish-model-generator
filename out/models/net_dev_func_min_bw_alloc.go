/* -----------------------------------------------------------------
* net_dev_func_min_bw_alloc.go -
*
* DMTF Redfish NetDevFuncMinBWAlloc resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// A minimum bandwidth allocation percentage for a network device functions associated a port.
type NetDevFuncMinBWAlloc struct {
	// The minimum bandwidth allocation percentage allocated to the corresponding network device function instance.
	MinBWAllocPercent int `json:"MinBWAllocPercent,omitempty"`

	// The link to the network device function associated with this bandwidth setting of this network port.
	NetworkDeviceFunction NetworkDeviceFunction `json:"NetworkDeviceFunction,omitempty"`

}
