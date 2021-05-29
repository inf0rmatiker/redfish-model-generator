/* -----------------------------------------------------------------
* network_device_function.go -
*
* DMTF Redfish NetworkDeviceFunction resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The NetworkDeviceFunction schema represents a logical interface that a network adapter exposes.
type NetworkDeviceFunction struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// An array of physical ports to which this network device function can be assigned.
	AssignablePhysicalNetworkPorts array `json:"AssignablePhysicalNetworkPorts,omitempty"`

	AssignablePhysicalNetworkPorts@odata.count count `json:"AssignablePhysicalNetworkPorts@odata.count,omitempty"`

	// An array of physical ports to which this network device function can be assigned.
	AssignablePhysicalPorts array `json:"AssignablePhysicalPorts,omitempty"`

	AssignablePhysicalPorts@odata.count count `json:"AssignablePhysicalPorts@odata.count,omitempty"`

	// The boot mode configured for this network device function.
	BootMode BootMode `json:"BootMode,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An indication of whether the network device function is enabled.
	DeviceEnabled bool `json:"DeviceEnabled,omitempty"`

	// The Ethernet capabilities, status, and configuration values for this network device function.
	Ethernet Ethernet `json:"Ethernet,omitempty"`

	// The Fibre Channel capabilities, status, and configuration values for this network device function.
	FibreChannel FibreChannel `json:"FibreChannel,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The InfiniBand capabilities, status, and configuration values for this network device function.
	InfiniBand InfiniBand `json:"InfiniBand,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The number of virtual functions that are available for this network device function.
	MaxVirtualFunctions int `json:"MaxVirtualFunctions,omitempty"`

	// The link to the metrics associated with this network function.
	Metrics NetworkDeviceFunctionMetrics `json:"Metrics,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// An array of capabilities for this network device function.
	NetDevFuncCapabilities array `json:"NetDevFuncCapabilities,omitempty"`

	// The configured capability of this network device function.
	NetDevFuncType NetworkDeviceTechnology `json:"NetDevFuncType,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The physical port to which this network device function is currently assigned.
	PhysicalNetworkPortAssignment Port `json:"PhysicalNetworkPortAssignment,omitempty"`

	// The physical port to which this network device function is currently assigned.
	PhysicalPortAssignment NetworkPort `json:"PhysicalPortAssignment,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// An indication of whether single root input/output virtualization (SR-IOV) virtual functions are enabled for this network device function.
	VirtualFunctionsEnabled bool `json:"VirtualFunctionsEnabled,omitempty"`

	// The iSCSI boot capabilities, status, and configuration values for this network device function.
	iSCSIBoot iSCSIBoot `json:"iSCSIBoot,omitempty"`

}
