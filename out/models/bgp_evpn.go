/* -----------------------------------------------------------------
* bgp_evpn.go -
*
* DMTF Redfish BGPEvpn resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// BGP Ethernet Virtual Private Network (BGP EVPN) related properties for an Ethernet fabric.
type BGPEvpn struct {
	// Address Resolution Protocol (ARP) proxy status.
	ARPProxyEnabled bool `json:"ARPProxyEnabled,omitempty"`

	// Address Resolution Protocol (ARP) suppression status.
	ARPSupressionEnabled bool `json:"ARPSupressionEnabled,omitempty"`

	// The anycast gateway IPv4 address.
	AnycastGatewayIPAddress string `json:"AnycastGatewayIPAddress,omitempty"`

	// The anycast gateway MAC address.
	AnycastGatewayMACAddress string `json:"AnycastGatewayMACAddress,omitempty"`

	// The Ethernet Segment Identifier (ESI) number range for the fabric.
	ESINumberRange ESINumberRange `json:"ESINumberRange,omitempty"`

	// The Ethernet Virtual Private Network (EVPN) Instance number (EVI) number range for the fabric.
	EVINumberRange EVINumberRange `json:"EVINumberRange,omitempty"`

	// The gateway IPv4 address.
	GatewayIPAddress string `json:"GatewayIPAddress,omitempty"`

	// Network Discovery Protocol (NDP) proxy status.
	NDPProxyEnabled bool `json:"NDPProxyEnabled,omitempty"`

	// Network Discovery Protocol (NDP) suppression status.
	NDPSupressionEnabled bool `json:"NDPSupressionEnabled,omitempty"`

	// The Route Distinguisher (RD) number range for the fabric.
	RouteDistinguisherRange RouteDistinguisherRange `json:"RouteDistinguisherRange,omitempty"`

	// The Route Target (RT) number range for the fabric.
	RouteTargetRange RouteTargetRange `json:"RouteTargetRange,omitempty"`

	// Underlay multicast status.
	UnderlayMulticastEnabled bool `json:"UnderlayMulticastEnabled,omitempty"`

	// Suppression of unknown unicast packets.
	UnknownUnicastSuppressionEnabled bool `json:"UnknownUnicastSuppressionEnabled,omitempty"`

	// The VLAN tag range for the fabric.
	VLANIdentifierAddressRange VLANIdentifierAddressRange `json:"VLANIdentifierAddressRange,omitempty"`

}
