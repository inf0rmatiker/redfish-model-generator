/* -----------------------------------------------------------------
* controller_links.go -
*
* DMTF Redfish ControllerLinks resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// ControllerLinks - The links to other resources that are related to this resource.
// Modeled after DMTF Redfish schema ControllerLinks
type ControllerLinks struct {
	// An array of links to the network device functions associated with this network controller.
	NetworkDeviceFunctions []NetworkDeviceFunction `json:"NetworkDeviceFunctions,omitempty"`

	NetworkDeviceFunctionsOdataCount int `json:"NetworkDeviceFunctions@odata.count,omitempty"`

	// An array of links to the network ports associated with this network controller.
	NetworkPorts []NetworkPort `json:"NetworkPorts,omitempty"`

	NetworkPortsOdataCount int `json:"NetworkPorts@odata.count,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to the PCIe devices associated with this network controller.
	PCIeDevices []PCIeDevice `json:"PCIeDevices,omitempty"`

	PCIeDevicesOdataCount int `json:"PCIeDevices@odata.count,omitempty"`

}
