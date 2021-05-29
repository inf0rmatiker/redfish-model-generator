/* -----------------------------------------------------------------
* storage_controller_links.go -
*
* DMTF Redfish StorageControllerLinks resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The links to other resources that are related to this resource.
type StorageControllerLinks struct {
	// An array of links to the endpoints that connect to this controller.
	Endpoints array `json:"Endpoints,omitempty"`

	Endpoints@odata.count count `json:"Endpoints@odata.count,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to the PCIe functions that the storage controller produces.
	PCIeFunctions array `json:"PCIeFunctions,omitempty"`

	PCIeFunctions@odata.count count `json:"PCIeFunctions@odata.count,omitempty"`

	// An array of links to the storage services that connect to this controller.
	StorageServices array `json:"StorageServices,omitempty"`

	StorageServices@odata.count count `json:"StorageServices@odata.count,omitempty"`

}
