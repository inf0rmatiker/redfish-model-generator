/* -----------------------------------------------------------------
* net_dev_func_max_bw_alloc.go -
*
* DMTF Redfish NetDevFuncMaxBWAlloc resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// NetDevFuncMaxBWAlloc - A maximum bandwidth allocation percentage for a Network Device Functions associated a port.
// Modeled after DMTF Redfish schema NetDevFuncMaxBWAlloc
type NetDevFuncMaxBWAlloc struct {
	// The maximum bandwidth allocation percentage allocated to the corresponding network device function instance.
	MaxBWAllocPercent float64 `json:"MaxBWAllocPercent,omitempty"`

	// Contains the members of this collection.
	NetworkDeviceFunction map[string]interface{} `json:"NetworkDeviceFunction,omitempty"`

}
