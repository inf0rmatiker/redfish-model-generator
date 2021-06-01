/* -----------------------------------------------------------------
* network_port.go -
*
* DMTF Redfish NetworkPort resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// NetworkPort - A Network Port represents a discrete physical port capable of connecting to a network.
// Modeled after DMTF Redfish schema NetworkPort
type NetworkPort struct {
	OdataContext map[string]interface{} `json:"@odata.context,omitempty"`

	OdataId map[string]interface{} `json:"@odata.id,omitempty"`

	OdataType map[string]interface{} `json:"@odata.type,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	Status map[string]interface{} `json:"Status,omitempty"`

	// The physical port number label for this port.
	PhysicalPortNumber string `json:"PhysicalPortNumber,omitempty"`

	// The status of the link between this port and its link partner.
	LinkStatus LinkStatus `json:"LinkStatus,omitempty"`

	// The self-described link capabilities of this port.
	SupportedLinkCapabilities []SupportedLinkCapabilities `json:"SupportedLinkCapabilities,omitempty"`

	// Network Port Active Link Technology.
	ActiveLinkTechnology LinkNetworkTechnology `json:"ActiveLinkTechnology,omitempty"`

	// The set of Ethernet capabilities that this port supports.
	SupportedEthernetCapabilities []SupportedEthernetCapabilities `json:"SupportedEthernetCapabilities,omitempty"`

	// The array of minimum bandwidth allocation percentages for the Network Device Functions associated with this port.
	NetDevFuncMinBWAlloc []NetDevFuncMinBWAlloc `json:"NetDevFuncMinBWAlloc,omitempty"`

	// The array of maximum bandwidth allocation percentages for the Network Device Functions associated with this port.
	NetDevFuncMaxBWAlloc []NetDevFuncMaxBWAlloc `json:"NetDevFuncMaxBWAlloc,omitempty"`

	// The array of configured network addresses (MAC or WWN) that are associated with this Network Port, including the programmed address of the lowest numbered Network Device Function, the configured but not active address if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedNetworkAddresses []string `json:"AssociatedNetworkAddresses,omitempty"`

	// Whether IEEE 802.3az Energy Efficient Ethernet (EEE) is enabled for this network port.
	EEEEnabled bool `json:"EEEEnabled,omitempty"`

	// Whether Wake on LAN (WoL) is enabled for this network port.
	WakeOnLANEnabled bool `json:"WakeOnLANEnabled,omitempty"`

	// The largest maximum transmission unit (MTU) that can be configured for this network port.
	PortMaximumMTU float64 `json:"PortMaximumMTU,omitempty"`

	// The 802.3x flow control behavior negotiated with the link partner for this network port (Ethernet-only).
	FlowControlStatus FlowControl `json:"FlowControlStatus,omitempty"`

	// The locally configured 802.3x flow control setting for this network port.
	FlowControlConfiguration FlowControl `json:"FlowControlConfiguration,omitempty"`

	// Whether or not the port has detected enough signal on enough lanes to establish link.
	SignalDetected bool `json:"SignalDetected,omitempty"`

}
