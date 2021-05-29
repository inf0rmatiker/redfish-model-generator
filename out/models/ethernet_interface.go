/* -----------------------------------------------------------------
* ethernet_interface.go -
*
* DMTF Redfish EthernetInterface resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// The EthernetInterface schema represents a single, logical Ethernet interface or network interface controller (NIC).
type EthernetInterface struct {
	OdataContext string `json:"@odata.context,omitempty"`

	OdataEtag string `json:"@odata.etag,omitempty"`

	OdataId string `json:"@odata.id"`

	OdataType string `json:"@odata.type"`

	// The available actions for this resource.
	Actions map[string]interface{} `json:"Actions,omitempty"`

	// An indication of whether the speed and duplex are automatically negotiated and configured on this interface.
	AutoNeg bool `json:"AutoNeg,omitempty"`

	// DHCPv4 configuration for this interface.
	DHCPv4 DHCPv4Configuration `json:"DHCPv4,omitempty"`

	// DHCPv6 configuration for this interface.
	DHCPv6 DHCPv6Configuration `json:"DHCPv6,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The type of interface.
	EthernetInterfaceType EthernetDeviceType `json:"EthernetInterfaceType,omitempty"`

	// The complete, fully qualified domain name that DNS obtains for this interface.
	FQDN string `json:"FQDN,omitempty"`

	// An indication of whether full-duplex mode is enabled on the Ethernet connection for this interface.
	FullDuplex bool `json:"FullDuplex,omitempty"`

	// The DNS host name, without any domain information.
	HostName string `json:"HostName,omitempty"`

	// The IPv4 addresses currently in use by this interface.
	IPv4Addresses array `json:"IPv4Addresses,omitempty"`

	// The IPv4 static addresses assigned to this interface.  See IPv4Addresses for the addresses in use by this interface.
	IPv4StaticAddresses array `json:"IPv4StaticAddresses,omitempty"`

	// An array that represents the RFC6724-defined address selection policy table.
	IPv6AddressPolicyTable array `json:"IPv6AddressPolicyTable,omitempty"`

	// The IPv6 addresses currently in use by this interface.
	IPv6Addresses array `json:"IPv6Addresses,omitempty"`

	// The IPv6 default gateway address in use on this interface.
	IPv6DefaultGateway string `json:"IPv6DefaultGateway,omitempty"`

	// The IPv6 static addresses assigned to this interface.  See IPv6Addresses for the addresses in use by this interface.
	IPv6StaticAddresses array `json:"IPv6StaticAddresses,omitempty"`

	// The IPv6 static default gateways for this interface.
	IPv6StaticDefaultGateways array `json:"IPv6StaticDefaultGateways,omitempty"`

	// The identifier that uniquely identifies the resource within the collection of similar resources.
	Id string `json:"Id"`

	// An indication of whether this interface is enabled.
	InterfaceEnabled bool `json:"InterfaceEnabled,omitempty"`

	// The link status of this interface, or port.
	LinkStatus LinkStatus `json:"LinkStatus,omitempty"`

	// The links to other resources that are related to this resource.
	Links Links `json:"Links,omitempty"`

	// The currently configured MAC address of the interface, or logical port.
	MACAddress MACAddress `json:"MACAddress,omitempty"`

	// The currently configured maximum transmission unit (MTU), in bytes, on this interface.
	MTUSize int `json:"MTUSize,omitempty"`

	// The maximum number of static IPv6 addresses that can be configured on this interface.
	MaxIPv6StaticAddresses int `json:"MaxIPv6StaticAddresses,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The DNS servers in use on this interface.
	NameServers []string `json:"NameServers,omitempty"`

	// The OEM extension property.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The permanent MAC address assigned to this interface, or port.
	PermanentMACAddress MACAddress `json:"PermanentMACAddress,omitempty"`

	// The current speed, in Mbit/s, of this interface.
	SpeedMbps int `json:"SpeedMbps,omitempty"`

	// Stateless address autoconfiguration (SLAAC) parameters for this interface.
	StatelessAddressAutoConfig StatelessAddressAutoConfiguration `json:"StatelessAddressAutoConfig,omitempty"`

	// The statically-defined set of DNS server IPv4 and IPv6 addresses.
	StaticNameServers []string `json:"StaticNameServers,omitempty"`

	// The status and health of the resource and its subordinate or dependent resources.
	Status Status `json:"Status,omitempty"`

	// The UEFI device path for this interface.
	UefiDevicePath string `json:"UefiDevicePath,omitempty"`

	// If this network interface supports more than one VLAN, this property is absent.  VLAN collections appear in the Links property of this resource.
	VLAN VLAN `json:"VLAN,omitempty"`

	// The link to a collection of VLANs, which applies only if the interface supports more than one VLAN.  If this property applies, the VLANEnabled and VLANId properties do not apply.
	VLANs VLanNetworkInterfaceCollection `json:"VLANs,omitempty"`

}
