/* -----------------------------------------------------------------
* dhc_pv4_configuration.go -
*
* DMTF Redfish DHCPv4Configuration resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// DHCPv4 configuration for this interface.
type DHCPv4Configuration struct {
	// An indication of whether DHCP v4 is enabled on this Ethernet interface.
	DHCPEnabled bool `json:"DHCPEnabled,omitempty"`

	// DHCPv4 fallback address method for this interface.
	FallbackAddress DHCPFallback `json:"FallbackAddress,omitempty"`

	// An indication of whether this interface uses DHCP v4-supplied DNS servers.
	UseDNSServers bool `json:"UseDNSServers,omitempty"`

	// An indication of whether this interface uses a DHCP v4-supplied domain name.
	UseDomainName bool `json:"UseDomainName,omitempty"`

	// An indication of whether this interface uses a DHCP v4-supplied gateway.
	UseGateway bool `json:"UseGateway,omitempty"`

	// An indication of whether the interface uses DHCP v4-supplied NTP servers.
	UseNTPServers bool `json:"UseNTPServers,omitempty"`

	// An indication of whether the interface uses DHCP v4-supplied static routes.
	UseStaticRoutes bool `json:"UseStaticRoutes,omitempty"`

}
