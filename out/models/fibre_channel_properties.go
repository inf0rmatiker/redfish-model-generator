/* -----------------------------------------------------------------
* fibre_channel_properties.go -
*
* DMTF Redfish FibreChannelProperties resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Fibre Channel-specific properties for a port.
type FibreChannelProperties struct {
	// An array of configured World Wide Names (WWN) that are associated with this network port, including the programmed address of the lowest numbered network device function, the configured but not active address, if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedWorldWideNames []string `json:"AssociatedWorldWideNames,omitempty"`

	// The Fibre Channel Fabric Name provided by the switch.
	FabricName string `json:"FabricName,omitempty"`

	// The number of ports not on the associated device that the associated device has discovered through this port.
	NumberDiscoveredRemotePorts int `json:"NumberDiscoveredRemotePorts,omitempty"`

	// The connection type of this port.
	PortConnectionType PortConnectionType `json:"PortConnectionType,omitempty"`

}
