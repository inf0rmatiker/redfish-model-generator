/* -----------------------------------------------------------------
* net_dev_func_min_bw_alloc.go -
*
* DMTF Redfish NetDevFuncMinBWAlloc resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// NetDevFuncMinBWAlloc - A minimum bandwidth allocation percentage for a Network Device Functions associated a port.
// Modeled after DMTF Redfish schema NetDevFuncMinBWAlloc
type NetDevFuncMinBWAlloc struct {
	// The minimum bandwidth allocation percentage allocated to the corresponding network device function instance.
	MinBWAllocPercent float64 `json:"MinBWAllocPercent,omitempty"`

	// Contains the members of this collection.
	NetworkDeviceFunction map[string]interface{} `json:"NetworkDeviceFunction,omitempty"`

}
