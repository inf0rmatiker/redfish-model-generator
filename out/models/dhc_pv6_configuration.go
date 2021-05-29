/* -----------------------------------------------------------------
* dhc_pv6_configuration.go -
*
* DMTF Redfish DHCPv6Configuration resource defined as a Golang model
*
* © Copyright 2021 Hewlett Packard Enterprise Development LP
*
* ----------------------------------------------------------------- */

package openapi

// DHCPv6 configuration for this interface.
type DHCPv6Configuration struct {
	// Determines the DHCPv6 operating mode for this interface.
	OperatingMode DHCPv6OperatingMode `json:"OperatingMode,omitempty"`

	// An indication of whether the interface uses DHCP v6-supplied DNS servers.
	UseDNSServers bool `json:"UseDNSServers,omitempty"`

	// An indication of whether the interface uses a domain name supplied through DHCP v6 stateless mode.
	UseDomainName bool `json:"UseDomainName,omitempty"`

	// An indication of whether the interface uses DHCP v6-supplied NTP servers.
	UseNTPServers bool `json:"UseNTPServers,omitempty"`

	// An indication of whether the interface uses DHCP v6 rapid commit mode for stateful mode address assignments.  Do not enable this option in networks where more than one DHCP v6 server is configured to provide address assignments.
	UseRapidCommit bool `json:"UseRapidCommit,omitempty"`

}
