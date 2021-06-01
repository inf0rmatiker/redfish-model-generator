/* -----------------------------------------------------------------
* ethernet_interface.go -
*
* DMTF Redfish EthernetInterface resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// EthernetInterface - This type defines an Ethernet interface.
// Modeled after DMTF Redfish schema EthernetInterface
type EthernetInterface struct {
	// The number of lanes supported by this interface.
	MaxLanes int `json:"MaxLanes,omitempty"`

	// The maximum speed supported by this interface.
	MaxSpeedMbps int `json:"MaxSpeedMbps,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

}
