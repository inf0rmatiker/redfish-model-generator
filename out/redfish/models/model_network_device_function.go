/* -----------------------------------------------------------------
* network_device_function.go -
*
* DMTF Redfish NetworkDeviceFunction resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// NetworkDeviceFunction - A Network Device Function represents a logical interface exposed by the network adapter.
// Modeled after DMTF Redfish schema NetworkDeviceFunction
type NetworkDeviceFunction struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataEtag map[string]interface{} `json:"@odata.etag,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id"`

	OdataType map[string]interface{} `json:"@odata.type"`

	// The array of physical port references that this network device function may be assigned to.
	AssignablePhysicalPorts []NetworkPort `json:"AssignablePhysicalPorts,omitempty"`

	AssignablePhysicalPortsOdataCount int `json:"AssignablePhysicalPorts@odata.count,omitempty"`

	// The boot mode configured for this network device function.
	BootMode BootMode `json:"BootMode,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// Whether the network device function is enabled.
	DeviceEnabled bool `json:"DeviceEnabled,omitempty"`

	// Ethernet.
	Ethernet Ethernet `json:"Ethernet,omitempty"`

	// Fibre Channel.
	FibreChannel FibreChannel `json:"FibreChannel,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// Links.
	Links Links `json:"Links,omitempty"`

	// The number of virtual functions (VFs) that are available for this Network Device Function.
	MaxVirtualFunctions int `json:"MaxVirtualFunctions,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// Capabilities of this network device function.
	NetDevFuncCapabilities []NetworkDeviceTechnology `json:"NetDevFuncCapabilities,omitempty"`

	// The configured capability of this network device function.
	NetDevFuncType NetworkDeviceTechnology `json:"NetDevFuncType,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The physical port that this network device function is currently assigned to.
	PhysicalPortAssignment map[string]interface{} `json:"PhysicalPortAssignment,omitempty"`

	// This property describes the status and health of the resource and its children.
	Status map[string]interface{} `json:"Status,omitempty"`

	// Whether Single Root I/O Virtualization (SR-IOV) Virual Functions (VFs) are enabled for this Network Device Function.
	VirtualFunctionsEnabled bool `json:"VirtualFunctionsEnabled,omitempty"`

	// iSCSI Boot.
	iSCSIBoot iSCSIBoot `json:"iSCSIBoot,omitempty"`

}
