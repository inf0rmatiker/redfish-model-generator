/* -----------------------------------------------------------------
* controller_links.go -
*
* DMTF Redfish ControllerLinks resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The links to other resources that are related to this resource.
type ControllerLinks struct {
	// An array of links to the network device functions associated with this network controller.
	NetworkDeviceFunctions array `json:"NetworkDeviceFunctions,omitempty"`

	NetworkDeviceFunctions@odata.count count `json:"NetworkDeviceFunctions@odata.count,omitempty"`

	// An array of links to the network ports associated with this network controller.
	NetworkPorts array `json:"NetworkPorts,omitempty"`

	NetworkPorts@odata.count count `json:"NetworkPorts@odata.count,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to the PCIe devices associated with this network controller.
	PCIeDevices array `json:"PCIeDevices,omitempty"`

	PCIeDevices@odata.count count `json:"PCIeDevices@odata.count,omitempty"`

	// An array of links to the ports associated with this network controller.
	Ports array `json:"Ports,omitempty"`

	Ports@odata.count count `json:"Ports@odata.count,omitempty"`

}
