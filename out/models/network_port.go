/* -----------------------------------------------------------------
* network_port.go -
*
* DMTF Redfish NetworkPort resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The NetworkPort schema describes a network port, which is a discrete physical port that can connect to a network.
type NetworkPort struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// Network port active link technology.
	ActiveLinkTechnology LinkNetworkTechnology `json:"ActiveLinkTechnology,omitempty"`

	// An array of configured MAC or WWN network addresses that are associated with this network port, including the programmed address of the lowest numbered network device function, the configured but not active address, if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedNetworkAddresses []string `json:"AssociatedNetworkAddresses,omitempty"`

	// Network port current link speed.
	CurrentLinkSpeedMbps int `json:"CurrentLinkSpeedMbps,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An indication of whether IEEE 802.3az Energy-Efficient Ethernet (EEE) is enabled for this network port.
	EEEEnabled bool `json:"EEEEnabled,omitempty"`

	// The FC Fabric Name provided by the switch.
	FCFabricName string `json:"FCFabricName,omitempty"`

	// The connection type of this port.
	FCPortConnectionType PortConnectionType `json:"FCPortConnectionType,omitempty"`

	// The locally configured 802.3x flow control setting for this network port.
	FlowControlConfiguration FlowControl `json:"FlowControlConfiguration,omitempty"`

	// The 802.3x flow control behavior negotiated with the link partner for this network port (Ethernet-only).
	FlowControlStatus FlowControl `json:"FlowControlStatus,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The status of the link between this port and its link partner.
	LinkStatus LinkStatus `json:"LinkStatus,omitempty"`

	// The maximum frame size supported by the port.
	MaxFrameSize int `json:"MaxFrameSize,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// An array of maximum bandwidth allocation percentages for the network device functions associated with this port.
	NetDevFuncMaxBWAlloc array `json:"NetDevFuncMaxBWAlloc,omitempty"`

	// An array of minimum bandwidth allocation percentages for the network device functions associated with this port.
	NetDevFuncMinBWAlloc array `json:"NetDevFuncMinBWAlloc,omitempty"`

	// The number of ports not on this adapter that this port has discovered.
	NumberDiscoveredRemotePorts int `json:"NumberDiscoveredRemotePorts,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The physical port number label for this port.
	PhysicalPortNumber string `json:"PhysicalPortNumber,omitempty"`

	// The largest maximum transmission unit (MTU) that can be configured for this network port.
	PortMaximumMTU int `json:"PortMaximumMTU,omitempty"`

	// An indication of whether the port has detected enough signal on enough lanes to establish a link.
	SignalDetected bool `json:"SignalDetected,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The set of Ethernet capabilities that this port supports.
	SupportedEthernetCapabilities array `json:"SupportedEthernetCapabilities,omitempty"`

	// The link capabilities of this port.
	SupportedLinkCapabilities array `json:"SupportedLinkCapabilities,omitempty"`

	// The vendor Identification for this port.
	VendorId string `json:"VendorId,omitempty"`

	// An indication of whether Wake on LAN (WoL) is enabled for this network port.
	WakeOnLANEnabled bool `json:"WakeOnLANEnabled,omitempty"`

}
