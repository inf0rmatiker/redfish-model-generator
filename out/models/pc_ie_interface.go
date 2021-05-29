/* -----------------------------------------------------------------
* pc_ie_interface.go -
*
* DMTF Redfish PCIeInterface resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Properties that describe a PCIe interface.
type PCIeInterface struct {
	// The number of PCIe lanes in use by this device.
	LanesInUse int `json:"LanesInUse,omitempty"`

	// The number of PCIe lanes supported by this device.
	MaxLanes int `json:"MaxLanes,omitempty"`

	// The highest version of the PCIe specification supported by this device.
	MaxPCIeType PCIeTypes `json:"MaxPCIeType,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The version of the PCIe specification in use by this device.
	PCIeType PCIeTypes `json:"PCIeType,omitempty"`

}
