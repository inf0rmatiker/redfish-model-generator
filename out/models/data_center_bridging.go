/* -----------------------------------------------------------------
* data_center_bridging.go -
*
* DMTF Redfish DataCenterBridging resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Data center bridging (DCB) for capabilities of a controller.
type DataCenterBridging struct {
	// An indication of whether this controller is capable of data center bridging (DCB).
	Capable bool `json:"Capable,omitempty"`

}
