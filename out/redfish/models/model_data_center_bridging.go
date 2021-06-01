/* -----------------------------------------------------------------
* data_center_bridging.go -
*
* DMTF Redfish DataCenterBridging resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// DataCenterBridging - Data center bridging (DCB) for capabilities of a controller.
// Modeled after DMTF Redfish schema DataCenterBridging
type DataCenterBridging struct {
	// An indication of whether this controller is capable of data center bridging (DCB).
	Capable bool `json:"Capable,omitempty"`

}
