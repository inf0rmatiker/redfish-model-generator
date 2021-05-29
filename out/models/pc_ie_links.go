/* -----------------------------------------------------------------
* pc_ie_links.go -
*
* DMTF Redfish PCIeLinks resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The links to other Resources that are related to this Resource.
type PCIeLinks struct {
	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to the PCIe devices contained in this slot.
	PCIeDevice array `json:"PCIeDevice,omitempty"`

	PCIeDevice@odata.count count `json:"PCIeDevice@odata.count,omitempty"`

}
