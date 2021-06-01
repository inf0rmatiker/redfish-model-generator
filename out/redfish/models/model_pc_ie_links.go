/* -----------------------------------------------------------------
* pc_ie_links.go -
*
* DMTF Redfish PCIeLinks resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// PCIeLinks - The links to other Resources that are related to this Resource.
// Modeled after DMTF Redfish schema PCIeLinks
type PCIeLinks struct {
	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to the PCIe devices contained in this slot.
	PCIeDevice []PCIeDevice `json:"PCIeDevice,omitempty"`

	PCIeDeviceOdataCount int `json:"PCIeDevice@odata.count,omitempty"`

}
