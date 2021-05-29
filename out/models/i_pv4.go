/* -----------------------------------------------------------------
* i_pv4.go -
*
* DMTF Redfish IPv4 resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// IPv4 and Virtual LAN (VLAN) related addressing for an Ethernet fabric.
type IPv4 struct {
	// The anycast gateway IPv4 address.
	AnycastGatewayIPAddress string `json:"AnycastGatewayIPAddress,omitempty"`

	// The anycast gateway MAC address.
	AnycastGatewayMACAddress string `json:"AnycastGatewayMACAddress,omitempty"`

	// The Dynamic Host Configuration Protocol (DHCP) related addressing for this Ethernet fabric.
	DHCP DHCP `json:"DHCP,omitempty"`

	// The Domain Name Service (DNS) domain name for this Ethernet fabric.
	DNSDomainName string `json:"DNSDomainName,omitempty"`

	// The Domain Name Service (DNS) servers for this Ethernet fabric.
	DNSServer []string `json:"DNSServer,omitempty"`

	// Indicates if host subnets should be distributed into the fabric underlay.
	DistributeIntoUnderlayEnabled bool `json:"DistributeIntoUnderlayEnabled,omitempty"`

	// External BGP (eBGP) related addressing for this Ethernet fabric.
	EBGPAddressRange IPv4AddressRange `json:"EBGPAddressRange,omitempty"`

	// Link related addressing for this Ethernet fabric.
	FabricLinkAddressRange IPv4AddressRange `json:"FabricLinkAddressRange,omitempty"`

	// The gateway IPv4 address.
	GatewayIPAddress string `json:"GatewayIPAddress,omitempty"`

	// IPv4 related end host subnet addressing for this Ethernet fabric.
	HostAddressRange IPv4AddressRange `json:"HostAddressRange,omitempty"`

	// Internal BGP (iBGP) related addressing for this Ethernet fabric.
	IBGPAddressRange IPv4AddressRange `json:"IBGPAddressRange,omitempty"`

	// Loopback related addressing for this Ethernet fabric.
	LoopbackAddressRange IPv4AddressRange `json:"LoopbackAddressRange,omitempty"`

	// Management related addressing for this Ethernet fabric.
	ManagementAddressRange IPv4AddressRange `json:"ManagementAddressRange,omitempty"`

	// The Network Time Protocol (NTP) offset configuration.
	NTPOffsetHoursMinutes int `json:"NTPOffsetHoursMinutes,omitempty"`

	// The Network Time Protocol (NTP) servers for this Ethernet fabric.
	NTPServer []string `json:"NTPServer,omitempty"`

	// The Network Time Protocol (NTP) timezone for this Ethernet fabric.
	NTPTimezone string `json:"NTPTimezone,omitempty"`

	// The native Virtual LAN (VLAN) tag value.
	NativeVLAN int `json:"NativeVLAN,omitempty"`

	// Virtual LAN (VLAN) tag related addressing for this Ethernet fabric or for end host networks.
	VLANIdentifierAddressRange VLANIdentifierAddressRange `json:"VLANIdentifierAddressRange,omitempty"`

}
