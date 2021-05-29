/* -----------------------------------------------------------------
* ethernet_properties.go -
*
* DMTF Redfish EthernetProperties resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// Ethernet-specific properties for a port.
type EthernetProperties struct {
	// An array of configured MAC addresses that are associated with this network port, including the programmed address of the lowest numbered network device function, the configured but not active address, if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedMACAddresses []string `json:"AssociatedMACAddresses,omitempty"`

	// The locally configured 802.3x flow control setting for this port.
	FlowControlConfiguration FlowControl `json:"FlowControlConfiguration,omitempty"`

	// The 802.3x flow control behavior negotiated with the link partner for this port.
	FlowControlStatus FlowControl `json:"FlowControlStatus,omitempty"`

	// Enable/disable LLDP for this port.
	LLDPEnabled bool `json:"LLDPEnabled,omitempty"`

	// LLDP data being received on this link.
	LLDPReceive LLDPReceive `json:"LLDPReceive,omitempty"`

	// LLDP data being transmitted on this link.
	LLDPTransmit LLDPTransmit `json:"LLDPTransmit,omitempty"`

	// The set of Ethernet capabilities that this port supports.
	SupportedEthernetCapabilities array `json:"SupportedEthernetCapabilities,omitempty"`

}
